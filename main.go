package main

import (
	"log"
	"os"

	"apigolang/controller"
	"apigolang/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// RegisterRoutes enregistre les routes de l'API avec leur handler correspondant
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", controller.GetAllProducts)
			products.POST("", controller.CreateProduct)
			products.PUT("/:id", controller.UpdateProduct)
			products.DELETE("/:id", controller.DeleteProduct)
			products.GET("/:id", controller.GetProductById)
		}
		payments := api.Group("/payments")
		{
			payments.GET("", controller.GetAllPayments)
			payments.POST("", controller.CreatePayment)
			payments.PUT("/:id", controller.UpdatePayment)
			payments.DELETE("/:id", controller.DeletePayment)
			payments.GET("/:id", controller.GetPaymentById)
			payments.GET("/stream", controller.StreamPayments)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		HOST string = os.Getenv("HOST")
		PORT string = os.Getenv("PORT")
	)

	// initialise la connexion à la base de données
	model.ModelInstance = model.Connect()
	defer model.ModelInstance.DB.Close()

	// initialise le serveur gin
	r := gin.Default()

	// enregistre les routes de l'API
	RegisterRoutes(r)

	// démarre le serveur
	if err := r.Run(HOST + ":" + PORT); err != nil {
		log.Fatal(err)
	}
}
