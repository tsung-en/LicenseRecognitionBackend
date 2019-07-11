package main

import (
	"LicenseRecognitionBackend/db"
	"LicenseRecognitionBackend/router"
	. "LicenseRecognitionBackend/model"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("env.ini")
	if err != nil {
		fmt.Println("Fail to read env file")
		os.Exit(1)
	}

	port := cfg.Section("server").Key("port").MustInt(9999)

	db.Eloquent.AutoMigrate(&User{}, &UserCar{}, &Parking{})
	defer db.Eloquent.Close()

	router := router.InitRouter()
	router.Run(fmt.Sprintf(":%d", port))

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// })
	// r.Run(fmt.Sprintf(":%d", port))
}
