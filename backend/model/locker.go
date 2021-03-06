package model

import (
	"github.com/go-playground/log"
)

type Locker struct {
	Id           int     `gorm:"primary_key;not null;index" json:"id"`
	Position     int     `gorm:"not null" json:"position"`
	Availability bool    `gorm:"not null;index" json:"availability"`
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

func GetLockersByUid(id int) ([]*Locker, error) {
	var lockers []*Locker
	err := db.
		Preload("CabinetInfo").
		Preload("UserInfo").
		Where("Uid = (?) AND Availability = 0", id).
		Find(&lockers).
		Error
	return lockers, err
}

func GetLockerById(id int) (Locker, error) {
	var result Locker
	err := db.
		Preload("CabinetInfo").
		Preload("UserInfo").
		First(&result, id).
		Error
	if err != nil {
		return Locker{}, err
	}
	return result, nil
}

func ReleaseLockerById(id int, uid int) error {
	var result Locker
	err := db.
		Where("uid = (?)", uid).
		First(&result, id).
		Error
	if err != nil {
		return err
	}
	// ref:https://github.com/jinzhu/gorm/issues/202
	//result.Availability = true
	//result.Uid = 0
	err = db.Model(&result).
		Updates(map[string]interface{}{
			"Availability": true,
			"Uid":          0,
		}).Error
	return err
}

func OccupyLockerById(id int, uid int) error {
	var result Locker
	err := db.First(&result, id).Error
	if err != nil {
		return err
	}
	// ref:https://github.com/jinzhu/gorm/issues/202
	//result.Availability = false
	//result.Uid = uid
	err = db.Model(&result).
		Updates(map[string]interface{}{
			"Availability": false,
			"Uid":          uid,
		}).Error
	return err
}

func GetFreeLockers(cid int) ([]int, error) {
	var lockers []*Locker
	err := db.Model(&Locker{}).
		Select("id").
		Where("cid = (?)", cid).
		Not("availability", false).
		Find(&lockers).
		Error
	if err != nil {
		return nil, err
	}

	var result []int
	for i := range lockers {
		result = append(result, lockers[i].Id)
	}
	return result, nil
}
