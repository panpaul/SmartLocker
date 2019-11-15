package model

import "github.com/go-playground/log"

type User struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT;index;not_null" json:"id"`
	Username string `gorm:"not null;index;UNIQUE" json:"username"`
	Password []byte `gorm:"not null" json:"password"`
	Role     int    `gorm:"not null" json:"role"`
	Connect  string `gorm:"not null" json:"connect"`
}

const (
	ADMIN = 1
	USER  = 0
)

func AddUser(username string, password []byte, role int, connect string) error {
	u := User{
		Username: username,
		Password: password,
		Role:     role,
		Connect:  connect,
	}
	err := db.Create(&u).Error
	return err
}

func AddUsers(u []*User) error {
	tx := db.Begin()
	if tx.Error != nil {
		log.WithError(tx.Error).Warn("couldn't start the transaction")
		return tx.Error
	}
	for _, v := range u {
		if err := tx.Create(&v).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func GetUserInfoById(id int) (*User, error) {
	var result User
	err := db.First(&result, id).Error
	return &result, err
}

func GetUserInfoByName(username string) (*User, error) {
	var result User
	err := db.Where(&User{Username: username}).Find(&result).Error
	return &result, err
}
