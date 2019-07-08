package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("env.ini")
	if err != nil {
		fmt.Println("Fail to read env file")
		os.Exit(1)
	}

	port := cfg.Section("server").Key("port").MustInt(9999)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	r.Run(fmt.Sprintf(":%d", port))
}
