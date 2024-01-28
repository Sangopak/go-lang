package server

import (
    "github.com/gin-gonic/gin"
)

func CreateDefaultRouter() *gin.Engine {
	return gin.Default()
}