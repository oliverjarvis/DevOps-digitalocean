package persistence

import (
	"go-minitwit/src/models"
	"go-minitwit/src/util"

	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
	var users []models.User
	result := db.Find(&users)
	if result.RowsAffected == 0 {
		addUsersAndMessages(db)
	}
}

func addUsersAndMessages(db *gorm.DB) {
	user1 := models.User{Username: "Tester",
		Email:   "tester@gmail.com",
		PW_hash: util.HashPassword("Test"),
		Messages: []*models.Message{
			{Text: "In Japan"},
		},
	}
	user2 := models.User{
		Username: "Cool",
		Email:    "cool@gmail.com",
		PW_hash:  util.HashPassword("Secret"),
		Messages: []*models.Message{
			{Text: "Hello World"},
		},
		Followers: []*models.User{&user1},
	}

	db.Create([]*models.User{&user1, &user2})
	db.Model(&user1).Association("Followers").Append(&user2)
	db.Model(&user2).Association("Followers").Append(&user1)
	db.Save([]*models.User{&user1, &user2})
}
