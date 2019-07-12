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
	CheckOut time.Time
	Payed    bool `gorm:"default:0"`
	Income   int
}

func (parking Parking) SelectByCarVal() (re Parking, err error) {
	if err = db.Eloquent.Where("car = ?", parking.Car).Find(&re).Error; err != nil {
		return
	}
	return
}
