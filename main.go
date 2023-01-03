package main

import (
	"log"
	"os"

	"apigolang/controller"
	"apigolang/model"

	_ "apigolang/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes enregistre les routes de l'API avec leur handler correspondant
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", controller.GetAllProducts)
			products.POST("", controller.CreateProduct)
			products.PUT("", controller.UpdateProduct)
			products.DELETE("/:id", controller.DeleteProduct)
			products.GET("/:id", controller.GetProductById)
		}
		payments := api.Group("/payments")
		{
			payments.GET("", controller.GetAllPayments)
			payments.POST("", controller.CreatePayment)
			payments.PUT("", controller.UpdatePayment)
			payments.DELETE("/:id", controller.DeletePayment)
			payments.GET("/:id", controller.GetPaymentById)
			payments.GET("/stream", controller.StreamPayments)
		}
		auth := api.Group("/auth")
		{
			auth.POST("/login", controller.Login)
			auth.POST("/register", controller.Register)
		}
		swagger := api.Group("/swagger")
		{
			swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
