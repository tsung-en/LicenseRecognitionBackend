package controller

import (
	models "LicenseRecognitionBackend/model"
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

type CarType struct {
	Car string `json:"car"  binding:"required"`
}

func ParkingCheckIn(c *gin.Context) {

	// c.JSON(200, c)
	var CarType CarType
	err := c.ShouldBindJSON(&CarType)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	var parking models.Parking
	parking.Car = CarType.Car

	p, e := parking.SelectByCarVal()

	if e == nil && p.CheckOut == nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%s dosen't checkout", parking.Car),
		})
		return
	}

	parking.CheckIn = time.Now()
	err = parking.Insert()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"result": &parking,
	})
	return
}

func ParkingCheckOut(c *gin.Context) {
	cfg, err := ini.Load("env.ini")

	var CarType CarType
	err = c.ShouldBindJSON(&CarType)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	var parking models.Parking
	parking.Car = CarType.Car
	p, err := parking.SelectByCarVal()

	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	if checkOutValidation := p.CheckOut; checkOutValidation != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%s not checkIn", CarType.Car),
		})
		return
	}

	now := time.Now()
	parking.CheckOut = &now

	hours := now.Sub(p.CheckIn).Hours()
	fHours := math.Floor(hours)
	dHours := hours - fHours

	cost := cfg.Section("parking").Key("cost").MustInt(99)
	amount := int(fHours) * cost
	if dHours > 0.5 {
		amount += cost
	}
	parking.Income = amount

	if err = parking.Update(&p); err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"amount": parking.Income,
		"car":    parking.Car,
	})
	return
}

func ParkingInfo(c *gin.Context) {
	var parkingInfo models.Parking

	parkingInfo.Car = c.Param("car")

	result, err := parkingInfo.SelectByCarVal()

	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(200, result)
	return

}
