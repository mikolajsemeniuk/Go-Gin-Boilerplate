package auth

import "github.com/gin-gonic/gin"

type Controller interface {
	Find(context *gin.Context)
}
