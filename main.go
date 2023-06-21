package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// TestRun Rotuer
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "success",
			"message":  "Running project on port",
		})
	})
	// End TestRun Rotuer

	// DataDummy Router
	router.GET("/data-dummy", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "success",
			"message":  "get Data Success",
		})
	})
	// End DataDummy Router

	router.Run()
}
