package main

import (
	"net/http"

	"testapi-go/controllers"

	"github.com/gin-gonic/gin"
)

// DataDummyController
type typeArraySatu struct {
	ID   string
	Name string
	Age  int
}

var arraySatu = [4]typeArraySatu{
	{
		ID:   "1",
		Name: "John Doe",
		Age:  25,
	},
	{
		ID:   "2",
		Name: "Jane Smith",
		Age:  30,
	},
	{
		ID:   "3",
		Name: "Michael Johnson",
		Age:  40},
	{
		ID:   "4",
		Name: "Emily Davis",
		Age:  35,
	},
}

func GetDataDummy(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"message":  "Get Data Success",
		"data":     arraySatu,
	})
}

func GetDataDummyById(c *gin.Context) {
	// id := 2
	id := c.Param("id")
	var dataDummyById typeArraySatu
	for _, typeArraySatu := range arraySatu {
		if typeArraySatu.ID == id {
			dataDummyById = typeArraySatu
			break
		}
	}

	if dataDummyById.ID != id {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "Data with ID: " + id + " Not Found",
		})
		return
	}

	c.JSON(200, gin.H{
		"response": "success",
		"message":  "Data with ID: " + id + " Success",
		"data":     dataDummyById,
	})
}

// End DataDummyController

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
	router.GET("/datadummy", GetDataDummy)
	router.GET("/datadummy-byid/:id", GetDataDummyById)
	// End DataDummy Router

	// User Router
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUserById)
	router.POST("/adduser", controllers.AddUser)
	router.PATCH("/updateuser/:id", controllers.UpdateUser)
	// End User Router

	router.Run(":3002")
}
