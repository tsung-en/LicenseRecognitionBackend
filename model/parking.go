package models

import (
	//"LicenseRecognitionBackend/db"
	"time"

	"github.com/jinzhu/gorm"
)

type Parking struct {
	gorm.Model
	Car      string
	Email    string
	CheckIn  time.Time
	CheckOut time.Time
}
