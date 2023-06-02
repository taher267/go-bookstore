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

	d, err := gorm.Open("mysql", "golang:golang@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(`DB Has been connected`)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
