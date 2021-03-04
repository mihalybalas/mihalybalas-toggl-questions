package controllers

import (
	"encoding/json"
	"fmt"
	"../utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"../models"
)

var NewQuestion models.Question

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	CreateQuestion := &models.Question{}
	utils.ParseBody(r, CreateQuestion)
	b:= CreateQuestion.CreateQuestion()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions:= models.GetAllQuestions()
	res, _ := json.Marshal(questions)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	var updateQuestion = &models.Question{}
	utils.ParseBody(r, updateQuestion)
	vars := mux.Vars(r)
	questionId := vars["questionId"]
	id, err:= strconv.ParseInt(questionId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	questionDetails, db:= models.GetQuestionById(id)
	if updateQuestion.Body != "" {
		questionDetails.Body = updateQuestion.Body
	}

	db.Save(&questionDetails)
	res, _ := json.Marshal(questionDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
