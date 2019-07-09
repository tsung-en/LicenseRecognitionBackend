package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

var Eloquent *gorm.DB
var err error

func init() {
	cfg, _ := ini.Load("env.ini")

	connector := cfg.Section("db").Key("connector").String()
	config := cfg.Section("db").Key("config").String()

	Eloquent, err = gorm.Open(connector, config)

	if err != nil {
		fmt.Printf("DB connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("Eloquent error %v", Eloquent.Error)
	}
}
