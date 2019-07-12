package main

import (
	"LicenseRecognitionBackend/db"
	models "LicenseRecognitionBackend/model"
	"fmt"
	"time"
)

func main() {
	car1 := models.UserCar{
		Car: "1111-AB",
	}
	car2 := models.UserCar{
		Car: "2222-CC",
	}

	user1 := models.User{
		Username: "img21326",
		Password: "g462",
		Email:    "img21326@gmail.com",
		Token:    "asdqwe",
		Cars:     []models.UserCar{car1, car2},
	}

	car3 := models.UserCar{
		Car: "3333-ED",
	}

	user2 := models.User{
		Username: "img",
		Password: "ggg",
		Email:    "img@gmail.com",
		Token:    "treewq",
		Cars:     []models.UserCar{car3},
	}

	checkin, err := time.Parse(
		time.RFC3339,
		"2019-07-01T08:00:00+08:00",
	)
	if err != nil {
		fmt.Printf("%v", err)
	}
	chekout, err := time.Parse(
		time.RFC3339,
		"2019-07-01T18:00:00+08:00",
	)
	parking1 := models.Parking{
		Car:      "1111-AB",
		CheckIn:  checkin,
		CheckOut: chekout,
	}

	checkin, err = time.Parse(
		time.RFC3339,
		"2019-07-03T09:00:00+08:00",
	)
	if err != nil {
		fmt.Printf("%v", err)
	}
	chekout, err = time.Parse(
		time.RFC3339,
		"2019-07-04T13:00:00+08:00",
	)
	parking2 := models.Parking{
		Car:      "3333-ED",
		CheckIn:  checkin,
		CheckOut: chekout,
	}

	db.Eloquent.Create(&user1)
	db.Eloquent.Create(&user2)

	db.Eloquent.Create(&parking1)
	db.Eloquent.Create(&parking2)
	fmt.Print("Finished Seeding Database")
}
