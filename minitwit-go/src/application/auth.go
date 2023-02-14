package application

import (
	"errors"
	"go-minitwit/src/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Register struct {
	Username  string `form:"username" binding:"required,min=1"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required,min=1"`
	Password2 string `form:"password2" binding:"required,min=1"`
}

type Login struct {
	Username string `form:"username" binding:"required,min=1"`
	Password string `form:"password" binding:"required,min=1"`
}

func HandleRegister(context *gin.Context, db *gorm.DB) error {
	var register Register
	if err := context.ShouldBind(&register); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return errors.New(ve[0].Field() + " is required")
		}
	}

	if register.Password != register.Password2 {
		return errors.New("Passwords don't match")
	}

	var existingUser = User{Username: register.Username}
	var result = db.Find(&existingUser)
	if result != nil {
		return errors.New("The username is already taken")
	}

	db.Save(User{
		Username: register.Username,
		Email:    register.Email,
		PW_hash:  util.HashPassword(register.Password),
	})

	return nil
}

func HandleLogin(context *gin.Context, db *gorm.DB) error {
	var login Login
	if err := context.ShouldBind(&login); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return errors.New(ve[0].Field() + " is required")
		}
	}

	var user = User{Username: login.Username}
	var result = db.Find(&user)
	if result == nil {
		return errors.New("Invalid username")
	}

	if util.CheckPasswordHash(login.Password, user.PW_hash) {
		return errors.New("Invalid password")
	}

	return nil
}
