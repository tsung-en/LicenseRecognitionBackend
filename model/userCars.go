package models

import (
	//"LicenseRecognitionBackend/db"
	"github.com/jinzhu/gorm"
)

type UserCar struct {
	gorm.Model
	Car    string
	UserID uint
}
