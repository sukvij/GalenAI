package user

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateUser(user *User, db *gorm.DB) error {
	return db.Create(&user).Error
}

func GetUserByUserName(credential *Credentials, db *gorm.DB) (*User, error) {
	var res User
	err := db.Where("user_name = ? AND password = ?",
		credential.UserName, credential.Password).
		First(&res).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(res)
	return &res, nil
}
