package controllers

import (
	"github.com/gin-gonic/gin"
)

// DataUserDummy
type typeDataUsers struct {
	ID   string
	Name string
	Age  int
}

var dataUsers = []typeDataUsers{
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
		Age:  40,
	},
}

// End DataUserDummy

// GetUsers
func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "success",
		"message":  "Get Data Users Success",
		"data":     dataUsers,
	})
}

// End GetUsers

// GetUserById
func GetUserById(c *gin.Context) {
	userId := c.Param("id")

	// CheckUserWithIDIsExist
	var dataUserById typeDataUsers
	for _, typeDataUsers := range dataUsers {
		if typeDataUsers.ID == userId {
			dataUserById = typeDataUsers
			break
		}
	}
	if dataUserById.ID != userId {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "User with ID: " + userId + " Not Found",
		})
		return
	}
	// End CheckUserWithIDIsExist

	c.JSON(200, gin.H{
		"response": "success",
		"message":  "Get Data User Success",
		"data":     dataUserById,
	})
}

// End GetUserById

// AddUsers
func AddUser(c *gin.Context) {

	dataInput := typeDataUsers{
		ID:   "4",
		Name: "Emily Davis",
		Age:  35,
	}

	// CheckUserWithIDIsExist
	var dataUserById typeDataUsers
	for _, typeDataUsers := range dataUsers {
		if typeDataUsers.ID == dataInput.ID {
			dataUserById = typeDataUsers
			break
		}
	}
	if dataInput.ID == dataUserById.ID {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "User with ID: " + dataInput.ID + " Already Exist",
		})
		return
	}
	// End CheckUserWithIDIsExist

	dataUsers = append(dataUsers, dataInput)

	c.JSON(200, gin.H{
		"response":  "success",
		"message":   "Add User Success",
		"dataInput": dataInput,
		"dataUsers": dataUsers,
	})
}

// End AddUsers

// UpdateUser
func UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	// CheckUserWithIDIsExist
	var dataUserById typeDataUsers
	for _, typeDataUsers := range dataUsers {
		if typeDataUsers.ID == userId {
			dataUserById = typeDataUsers
			break
		}
	}
	if dataUserById.ID != userId {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "Users with ID: " + userId + " Not Found",
		})
		return
	}
	// End CheckUserWithIDIsExist

	// UpdateUser
	for i, typeArraySatu := range dataUsers {
		if typeArraySatu.ID == userId {
			dataUsers[i].Name = "Jon Heri"
			break
		}
	}
	// End UpdateUser

	c.JSON(200, gin.H{
		"response": "success",
		"message":  "Add User Success",
		"id":       userId,
		"data":     dataUsers,
	})
}

// End UpdateUser
