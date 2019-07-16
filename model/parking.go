package models

import (
	"LicenseRecognitionBackend/db"
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

func (parking *Parking) Update(update *Parking) (err error) {
	if err = db.Eloquent.Model(&update).Update(&parking).Error; err != nil {
		return
	}

	return
}

func (parking *Parking) SelectByCarVal() (re Parking, err error) {
	if err = db.Eloquent.Where("car = ?", parking.Car).Last(&re).Error; err != nil {
		return
	}
	return
}
