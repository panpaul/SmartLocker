package database

import "github.com/go-playground/log"

type Locker struct {
	Id           int     `gorm:"primary_key;not null;index" json:"id"`
	Position     int     `gorm:"not null;UNIQUE" json:"position"`
	Availability bool    `gorm:"not null" json:"availability"`
	Uid          int     `gorm:"index" json:"uid"`
	UserInfo     User    `gorm:"foreignkey:id;association_foreignkey:uid" json:"user"`
	Cid          int     `gorm:"not null;index" json:"cid"`
	CabinetInfo  Cabinet `gorm:"foreignkey:id;association_foreignkey:cid" json:"cabinet"`
}

func AddLocker(position int, cid int) error {
	l := Locker{
		Position:     position,
		Availability: true,
		Cid:          cid,
	}
	err := db.Create(&l).Error
	return err
}

func AddLockers(u []*Locker) error {
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
