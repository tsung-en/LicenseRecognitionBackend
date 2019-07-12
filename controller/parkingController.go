package controller

import (
	models "LicenseRecognitionBackend/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
