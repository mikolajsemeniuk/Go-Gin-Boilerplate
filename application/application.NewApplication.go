package application

import (
	"check/controllers/auth"

	"github.com/gin-gonic/gin"
)

func NewServer() Server {
	return &server{
		router: gin.Default(),
		auth:   auth.NewController(),
	}
}
