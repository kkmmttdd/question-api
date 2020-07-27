package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	QuestionController QuestionControllerInterface = &questionController{}
)

type QuestionControllerInterface interface {
	Index(context * gin.Context)
}

type questionController struct{}

func (c *questionController) Index(context * gin.Context) {
	context.String(http.StatusOK, "pong")
}
