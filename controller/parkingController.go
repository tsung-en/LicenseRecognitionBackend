package controller

import (
	models "LicenseRecognitionBackend/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckInType struct {
	Car  string    `json:"car"  binding:"required"`
	Time time.Time `json:"time" binding:"required"`
}

func ParkingCheckIn(c *gin.Context) {
	// c.JSON(200, c)
	var checkInType CheckInType
	err := c.ShouldBindJSON(&checkInType)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	var parking models.Parking
	parking.Car = checkInType.Car
	parking.CheckIn = checkInType.Time
	err = parking.Insert()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"checkintype": &checkInType,
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
