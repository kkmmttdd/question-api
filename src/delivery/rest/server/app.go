package server

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	Mapper()
	router.Run("localhost:8080")
}
