package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type typeDataProducts struct {
	ID          string
	ProductName string
	ProductType string
}

var dataProducts = []typeDataProducts{
	{
		ID:          "1",
		ProductName: "Lenovo Thinkpad",
		ProductType: "Laptop",
	},
}

// GetProducts
func GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "success",
		"message":  "Get Data Products Success",
		"data":     dataProducts,
	})
}

// End GetProducts

// AddProduct
func AddProduct(c *gin.Context) {

	var dataInput typeDataProducts
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		log.Fatal(err)
	}

	// CheckUserWithIDIsExist
	var dataProductById typeDataProducts
	for _, typeDataProducts := range dataProducts {
		if typeDataProducts.ID == dataInput.ID {
			dataProductById = typeDataProducts
			break
		}
	}
	if dataInput.ID == dataProductById.ID {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "Product with ID: " + dataInput.ID + " Already Exist",
		})
		return
	}
	// End CheckUserWithIDIsExist

	dataProducts = append(dataProducts, dataInput)

	c.JSON(200, gin.H{
		"response":  "success",
		"message":   "Get Data Products Success",
		"dataInput": dataInput,
		"data":      dataProducts,
	})
}

// End AddProduct

// GetProductById
func GetProductById(c *gin.Context) {
	productId := c.Param("id")

	// CheckUserWithIDIsExist
	var dataProductById typeDataProducts
	for _, typeDataProducts := range dataProducts {
		if typeDataProducts.ID == productId {
			dataProductById = typeDataProducts
			break
		}
	}
	if dataProductById.ID != productId {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "Product with ID: " + productId + " Not Found",
		})
		return
	}
	// End CheckUserWithIDIsExist

	c.JSON(200, gin.H{
		"response": "success",
		"message":  "Get Data Products Success",
		"data":     dataProductById,
	})
}

// End GetProductById

// UpdateProduct
func UpdateProduct(c *gin.Context) {

	productId := c.Param("id")

	var dataInput typeDataProducts
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		log.Fatal(err)
	}

	// CheckUserWithIDIsExist
	var dataProductById typeDataProducts
	for _, typeDataProducts := range dataProducts {
		if typeDataProducts.ID == productId {
			dataProductById = typeDataProducts
			break
		}
	}
	if dataProductById.ID != productId {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "Product with ID: " + productId + " Not Found",
		})
		return
	}
	// End CheckUserWithIDIsExist

	// UpdateUser
	for i, typeDataProducts := range dataProducts {
		if typeDataProducts.ID == productId {
			dataProducts[i].ProductName = dataInput.ProductName
			dataProducts[i].ProductType = dataInput.ProductType
			break
		}
	}
	// End UpdateUser

	c.JSON(200, gin.H{
		"response":  "success",
		"message":   "Get Data Products Success",
		"dataInput": dataInput,
		"data":      dataProducts,
	})
}

// End UpdateProduct

// DeleteProduct
func DeleteProduct(c *gin.Context) {

	productId := c.Param("id")

	// CheckUserWithIDIsExist
	var dataProductById typeDataProducts
	for _, typeDataProducts := range dataProducts {
		if typeDataProducts.ID == productId {
			dataProductById = typeDataProducts
			break
		}
	}
	if dataProductById.ID != productId {
		c.JSON(400, gin.H{
			"response": "fail",
			"message":  "Product with ID: " + productId + " Not Found",
		})
		return
	}
	// End CheckUserWithIDIsExist

	// Mencari objek dengan ID yang sesuai dan menghapusnya
	for i, typeDataProducts := range dataProducts {
		if typeDataProducts.ID == productId {
			dataProducts = append(dataProducts[:i], dataProducts[i+1:]...)
			break
		}
	}

	c.JSON(200, gin.H{
		"response": "success",
		"message":  "Get Data Products Success",
		"data":     dataProducts,
	})
}

// End DeleteProduct
