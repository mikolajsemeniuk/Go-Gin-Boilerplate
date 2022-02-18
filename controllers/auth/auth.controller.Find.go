package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *controller) Find(context *gin.Context) {
	controller.service.Find()
	fmt.Println("controller")
	context.JSON(http.StatusOK, "hi")
}
