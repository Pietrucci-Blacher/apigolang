package controller

import (
	"net/http"
	"strconv"

	"apigolang/model"

	"github.com/gin-gonic/gin"
)

type dataGetAllProductsReturn struct {
	Data []model.Product `json:"data"`
}

type dataGetProductByIdReturn struct {
	Data model.Product `json:"data"`
}

type dataCreateProductPost struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type dataBoolean struct {
	Data bool `json:"data"`
}

// @Summary récupère tous les produits de la base de données et les renvoie en format JSON
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200 {object} dataGetAllProductsReturn
// @Router /api/products [get]
func GetAllProducts(c *gin.Context) {
	products, err := model.ModelInstance.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// @Summary récupère le produit de la base de données avec son id et les renvoie en format JSON
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} dataGetProductByIdReturn
// @Router /api/products/{id} [get]
func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := model.ModelInstance.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// @Summary crée un nouveau produit dans la base de données et le renvoie en format JSON
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body dataCreateProductPost true "Product object"
// @Success 200 {object} dataGetProductByIdReturn
// @Router /api/products/ [post]
func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := model.ModelInstance.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

// @Summary met à jour un produit dans la base de données et le renvoie en format JSON
// @Tags Product
// @Accept  json
// @Produce  json
// @Param product body dataCreateProductPost true "Product object"
// @Success 200 {object} dataGetProductByIdReturn
// @Router /api/products/ [put]
func UpdateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := model.ModelInstance.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// @Summary supprime un produit de la base de données
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} dataBoolean
// @Router /api/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := model.ModelInstance.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
