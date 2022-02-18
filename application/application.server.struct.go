package application

import (
	"check/controllers/auth"

	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
	auth   auth.Controller
}
