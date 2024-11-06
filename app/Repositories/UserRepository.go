package Repositories

import (
	"errors"
	"gorm.io/gorm"
	"user/app/Models"
	"user/database"
)

func GetUserById(id int) (*Models.User, error) {
	db := database.DBConn
	var user Models.User

	if err := db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(e string) (*Models.User, error) {
	db := database.DBConn
	var user Models.User

	if err := db.Where(&Models.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
