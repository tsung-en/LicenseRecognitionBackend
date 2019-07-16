package router

import (
	. "LicenseRecognitionBackend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", Hello)

	parkingGroup := router.Group("/parking")

	parkingGroup.POST("/checkin", ParkingCheckIn)
	parkingGroup.PUT("/checkout", ParkingCheckOut)
	parkingGroup.GET("/:car/info", ParkingInfo)

	return router
}
