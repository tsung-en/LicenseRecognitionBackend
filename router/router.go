package router

import (
	. "LicenseRecognitionBackend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", Hello)

	return router
}
