package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

var RegisterQuestionRoutes = func(router *mux.Router) {
	router.HandleFunc("/question/", controllers.GetQuestions).Methods("GET")
	router.HandleFunc("/question/", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/question/{questionId}", controllers.UpdateQuestion).Methods("PUT")
}