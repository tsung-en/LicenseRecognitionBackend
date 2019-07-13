package controller

import (
	models "LicenseRecognitionBackend/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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
		"checkintype": &parking,
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
