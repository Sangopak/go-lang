package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckMessage struct{
	Message string `json:"message"`
}

func HeatlhCheck(context *gin.Context) {
	ok := HealthCheckMessage{
		Message: "Ok",
	}
	context.IndentedJSON(http.StatusOK, ok)
}