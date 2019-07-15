package models

import (
	"LicenseRecognitionBackend/db"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Parking struct {
	gorm.Model
	Car      string
	CheckIn  time.Time
	CheckOut *time.Time `gorm:"default: null"`
	Payed    bool       `gorm:"default:0"`
	Income   int
}

func (parking *Parking) Insert() (err error) {
	err = db.Eloquent.Create(&parking).Error
	return
}

func (parking *Parking) Update() (err error) {
	var update Parking
	if err = db.Eloquent.Where("car = ?", parking.Car).Order("check_in").First(&update).Error; err != nil {
		fmt.Printf("%v", err)
		return
	}

	if checkOutValidation := update.CheckOut; checkOutValidation != nil {
		err = errors.New(fmt.Sprintf("not found %s check in", update.Car))
		return
	}

	if err = db.Eloquent.Model(&update).Update(&parking).Error; err != nil {
		return
	}

	return
}

func (parking *Parking) SelectByCarVal() (re Parking, err error) {
	if err = db.Eloquent.Where("car = ?", parking.Car).Find(&re).Error; err != nil {
		return
	}
	return
}
