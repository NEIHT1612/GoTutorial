package routes

import (
	"net/http"
	"strconv"

	"example.com/main/models"
	"github.com/gin-gonic/gin"
)

func getProducts(context *gin.Context){
	//Get all products
	products, err := models.GetAllProducts()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch list products"})
		return
	}
	context.JSON(http.StatusOK, products)
}

func createProducts(context *gin.Context){
	var product models.Product
	err := context.ShouldBindJSON(&product)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse product data"})
		return
	}
	product.ID = 1
	err = product.CreateProduct()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create product"})
		return
	}
	context.JSON(http.StatusOK, "Create successfully")
}

func getProductById(context *gin.Context){
	productId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't get product id"})
		return
	}
	product, err := models.GetProductById(productId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch product"})
		return
	}
	context.JSON(http.StatusOK, product)
}

func updateProduct(context *gin.Context){
	productId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't get product id"})
		return
	}
	_, err = models.GetProductById(productId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch product"})
		return
	}
	var updateProduct models.Product
	err = context.ShouldBindJSON(&updateProduct)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse JSON product"})
		return
	}
	updateProduct.ID = productId
	err = updateProduct.UpdateProduct()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't update product"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Update successfully"})
}

func deleteProduct(context *gin.Context){
	productId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't get product id"})
		return
	}
	
	product, err := models.GetProductById(productId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch product"})
		return
	}

	err = product.DeleteProduct()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't delete product"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})
}