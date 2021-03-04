package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("sqlite3", "./db/toggl-questions.db")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
