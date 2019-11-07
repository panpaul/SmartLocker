package model

import "github.com/go-playground/log"

type Cabinet struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT;unique_index;not_null" json:"id"`
	Name     string `gorm:"not null;UNIQUE" json:"name"`
	Location string `gorm:"not null" json:"location"`
}

func AddCabinet(name string, location string) error {
	c := Cabinet{
		Name:     name,
		Location: location,
	}
	err := db.Create(&c).Error
	return err
}

func AddCabinets(c []*Cabinet) error {
	tx := db.Begin()
	if tx.Error != nil {
		log.WithError(tx.Error).Warn("couldn't start the transaction")
		return tx.Error
	}
	for _, v := range c {
		if err := tx.Create(&v).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func GetCabinetsByLocation(where string) ([]int, error) {
	var r []int
	err := db.Model(&Cabinet{}).
		Where(&Cabinet{Location: where}).
		Pluck("DISTINCT(id)", &r).
		Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetCabinetLocations() ([]string, error) {
	var r []string
	err := db.Model(&Cabinet{}).Pluck("DISTINCT(location)", &r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}
