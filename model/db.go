package model

import (
	"SmartLocker/config"
	"fmt"
	"github.com/go-playground/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Setup() {
	var err error
	//TODO adapt sqlite
	db, err = gorm.Open(config.Conf.Database.Type,
		fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True",
			config.Conf.Database.User,
			config.Conf.Database.Password,
			config.Conf.Database.Address,
			config.Conf.Database.Port,
			config.Conf.Database.Database))

	if err != nil {
		log.WithError(err).Fatal("Couldn't connect to the database")
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Conf.Database.Prefix + defaultTableName
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(200)

	db.AutoMigrate(&User{}, &Cabinet{}, &Locker{})
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		log.WithError(err).Warn("Couldn't close the database")
	}
	log.Info("Successfully closed the database")
}
