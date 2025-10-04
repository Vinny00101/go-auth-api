package repository

import (
	"go-api/config"
	"go-api/model"
)

func Where_user_verify(email string) bool {
	return config.DB.Where("email = ?", email).First(&model.User{}).Error == nil
}

func Get_user_by_email(email string) (user *model.User) {
	config.DB.Where("email = ?", email).First(&user)
	return user
}

func Create_User(user *model.User) bool {
	if err := config.DB.Create(&user).Error; err != nil {
		return false
	}
	return true
}

func Get_User_By_ID(userID uint) (user *model.User) {
	config.DB.First(&user, userID)
	return user
}
