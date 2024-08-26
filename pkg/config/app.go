package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
		d, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/hello?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Successful")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
