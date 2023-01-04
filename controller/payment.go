package controller

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"apigolang/model"

	"github.com/gin-gonic/gin"
)

type dataGetAllPaymentReturn struct {
	Data []model.Payment `json:"data"`
}

type dataGetPaymentByIdReturn struct {
	Data model.Payment `json:"data"`
}

type dataCreatePaymentPost struct {
	ProductID int     `json:"product_id"`
	PricePaid float64 `json:"price_paid"`
}

// @Summary récupère tous les paiements de la base de données et les renvoie en format JSON
// @Tags Payments
// @Accept  json
// @Produce  json
// @Success 200 {object} dataGetAllPaymentReturn
// @Router /api/payments [get]
func GetAllPayments(c *gin.Context) {
	payments, err := model.ModelInstance.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// GetPaymentById
// @Summary récupère un paiement de la base de données par son ID et le renvoie en format JSON
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} dataGetPaymentByIdReturn
// @Router /api/payments/{id} [get]
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

// @Summary crée un nouveau paiement dans la base de données et le renvoie en format JSON
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param payment body dataCreatePaymentPost true "Product object"
// @Success 200 {object} dataGetPaymentByIdReturn
// @Router /api/payments [post]
func CreatePayment(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := model.ModelInstance.CreatePayment(&payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payment.ID = id
	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// @Summary met à jour un paiement dans la base de données et le renvoie en format JSON
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param payment body dataCreatePaymentPost true "Product object"
// @Success 200 {object} dataGetPaymentByIdReturn
// @Router /api/payments [put]
func UpdatePayment(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := model.ModelInstance.UpdatePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// @Summary met à jour un paiement dans la base de données et le renvoie en format JSON
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} dataBoolean
// @Router /api/payments/{id} [delete]
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

// @Summary récupère tous les paiements de la base de données et les renvoie
// @Tags Payments
// @Accept  json
// @Produce  json
// @Router /api/payments/stream [get]
func StreamPayments(c *gin.Context) {
	// Set the headers for SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	// Create a new ticker that sends a message to the client every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Create a channel that receives a message when the client closes the connection
	closed := make(chan bool, 1)

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", "je suis un test")
		select {
		case <-ticker.C:
			return true
		case <-closed:
			return false
		}
	})
}
