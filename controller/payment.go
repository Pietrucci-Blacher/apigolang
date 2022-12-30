package controller

import (
	"io"
	"net/http"
	"strconv"

	"apigolang/model"

	"github.com/gin-gonic/gin"
)

// CreatePayment crée un nouveau paiement dans la base de données et le renvoie en format JSON
func CreatePayment(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := model.ModelInstance.CreatePayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payment.ID = id
	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// UpdatePayment met à jour un paiement dans la base de données et le renvoie en format JSON
func UpdatePayment(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := model.ModelInstance.UpdatePayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DeletePayment supprime un paiement de la base de données
func DeletePayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}
	if err := model.ModelInstance.DeletePayment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetPaymentById récupère un paiement de la base de données par son ID et le renvoie en format JSON
func GetPaymentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}
	payment, err := model.ModelInstance.GetPaymentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

func StreamPayments(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", "Hello world")
		return true
	})
}
