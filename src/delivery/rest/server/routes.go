package server

import (
	"github.com/kkmmttdd/question-api/src/delivery/rest"
)

func Mapper() {
	router.GET("/questions", rest.QuestionController.Index)
}


