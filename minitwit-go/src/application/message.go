package application

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID  uint
	Text    string
	flagged bool
}

type MessageDTO struct {
	CreatedAt string
	Username  string
	Text      string
	AvatarURL string
}

func GetAllMessages(db *gorm.DB) []MessageDTO {
	var messages []Message
	db.Find(&messages)

	return toMessageDTO(db, messages)
}

func GetMessagesByUser(db *gorm.DB, userID uint) []MessageDTO {
	var messages []Message
	db.Find(&messages, userID)

	return toMessageDTO(db, messages)
}

func toMessageDTO(db *gorm.DB, messages []Message) []MessageDTO {
	var messageDTOs []MessageDTO
	var user User
	db.Find(&user)
	var avatarURL = getAvatarURL(user.Email)
	for _, message := range messages {
		messageDTO := MessageDTO{
			CreatedAt: message.CreatedAt.Format("2006-01-02"),
			Username:  user.Username,
			Text:      message.Text,
			AvatarURL: avatarURL,
		}
		messageDTOs = append(messageDTOs, messageDTO)
	}

	return messageDTOs
}

func getAvatarURL(email string) string {
	email_md5 := fmt.Sprintf("%s", md5.Sum([]byte(email)))
	hex_md5_email := hex.EncodeToString([]byte(email_md5))
	url := fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon&s=48", hex_md5_email)

	return url
}
