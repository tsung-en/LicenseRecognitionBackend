package models

import (
	//"LicenseRecognitionBackend/db"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Token    string
	Cars     []UserCar `gorm:"foreignkey:UserID"`
}
