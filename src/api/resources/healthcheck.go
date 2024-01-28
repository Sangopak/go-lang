package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHeatlhCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
}
