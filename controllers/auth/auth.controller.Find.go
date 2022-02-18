package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller) Find(context *gin.Context) {
	c.service.Find()
	fmt.Println("controller")
	context.JSON(http.StatusOK, "hi")
}
