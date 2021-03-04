package models

import (
	"../config"
	"github.com/jinzhu/gorm"
)

var questionDb *gorm.DB

type Question struct {
	Model
	Body        string `json:"body"`
	Options		[]Option `gorm:"foreignKey:QuestionId"`
}

func init() {
	config.Connect()
	questionDb = config.GetDB()
	questionDb.AutoMigrate(&Question{})
}

func (q *Question) CreateQuestion() *Question {
	questionDb.Create(&q)
	return q
}

func  GetAllQuestions() []Question {
	var Questions []Question
	questionDb.Preload("Options").Order("ID").Find(&Questions)
	return Questions
}

func GetQuestionById(Id int64) (*Question , *gorm.DB){
	var getQuestion Question
	db:=questionDb.Preload("Options").Where("ID = ?", Id).Find(&getQuestion)
	return &getQuestion, db
}
