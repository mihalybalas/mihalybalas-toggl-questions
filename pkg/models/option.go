package models

import (
	"../config"
	"github.com/jinzhu/gorm"
)

var optionDb *gorm.DB

type Option struct {
	Model
	Body        string `json:"body"`
	Correct     bool `json:"correct"`
	QuestionId 	int `json:"-"`
	Question 	Question `gorm:"foreignKey:QuestionId" json:"-"`
}

func init() {
	config.Connect()
	optionDb = config.GetDB()
	optionDb.AutoMigrate(&Option{})
}

func (o *Option) CreateOption() *Option {
	optionDb.NewRecord(o)
	optionDb.Create(&o)
	return o
}
