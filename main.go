package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var model Connection

// Connection est une structure de données pour la connexion à la base de données
type Connection struct {
	DB *sql.DB
}

// Connect établit une connexion à la base de données et renvoie une Connection
func Connect() Connection {
	db, err := sql.Open("mysql", "root:azerty@tcp(localhost:3306)/go_api?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return Connection{db}
}

// Model Product

// Product est la structure de données pour un produit
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateProduct crée un nouveau produit dans la base de données et renvoie l'ID du produit
func (conn Connection) CreateProduct(product Product) (int, error) {
	// prépare la requête pour insérer le produit dans la base de données
	stmt, err := conn.DB.Prepare("INSERT INTO product (name, price, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// exécute la requête avec les données du produit
	res, err := stmt.Exec(product.Name, product.Price, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	// récupère l'ID du produit créé
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// UpdateProduct met à jour un produit dans la base de données
func (conn Connection) UpdateProduct(product Product) error {
	// prépare la requête pour mettre à jour le produit dans la base de données
	stmt, err := conn.DB.Prepare("UPDATE product SET name = ?, price = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// exécute la requête avec les données du produit
	_, err = stmt.Exec(product.Name, product.Price, time.Now(), product.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct supprime un produit de la base de données
func (conn Connection) DeleteProduct(id int) error {
	// prépare la requête pour supprimer le produit de la base de données
	stmt, err := conn.DB.Prepare("DELETE FROM product WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// exécute la requête avec l'ID du produit
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// GetProductById récupère un produit de la base de données par son ID
func (conn Connection) GetProductById(id int) (Product, error) {
	// prépare la requête pour récupérer le produit de la base de données
	stmt, err := conn.DB.Prepare("SELECT id, name, price, created_at, updated_at FROM product WHERE id = ?")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()

	// exécute la requête avec l'ID du produit
	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

// GetAllProducts récupère tous les produits de la base de données
func (conn Connection) GetAllProducts() ([]Product, error) {
	// prépare la requête pour récupérer tous les produits de la base de données
	stmt, err := conn.DB.Prepare("SELECT id, name, price, created_at, updated_at FROM product")
	if err != nil {
		return []Product{}, err
	}
	defer stmt.Close()

	// exécute la requête
	rows, err := stmt.Query()
	if err != nil {
		return []Product{}, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return []Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Model Payment

// Payment est la structure de données pour un paiement
type Payment struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	PricePaid float64   `json:"price_paid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreatePayment crée un nouveau paiement dans la base de données et renvoie l'ID du paiement
func (conn Connection) CreatePayment(payment Payment) (int, error) {
	// prépare la requête pour insérer le paiement dans la base de données
	stmt, err := conn.DB.Prepare("INSERT INTO payment (product_id, price_paid, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// exécute la requête avec les données du paiement
	res, err := stmt.Exec(payment.ProductID, payment.PricePaid, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	// récupère l'ID du paiement créé
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// UpdatePayment met à jour un paiement dans la base de données
func (conn Connection) UpdatePayment(payment Payment) error {
	// prépare la requête pour mettre à jour le paiement dans la base de données
	stmt, err := conn.DB.Prepare("UPDATE payment SET product_id = ?, price_paid = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// exécute la requête avec les données du paiement
	_, err = stmt.Exec(payment.ProductID, payment.PricePaid, time.Now(), payment.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeletePayment supprime un paiement de la base de données
func (conn Connection) DeletePayment(id int) error {
	// prépare la requête pour supprimer le paiement de la base de données
	stmt, err := conn.DB.Prepare("DELETE FROM payment WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// exécute la requête avec l'ID du paiement
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// GetPaymentById récupère un paiement de la base de données par son ID
func (conn Connection) GetPaymentById(id int) (Payment, error) {
	// prépare la requête pour récupérer le paiement de la base de données
	stmt, err := conn.DB.Prepare("SELECT id, product_id, price_paid, created_at, updated_at FROM payment WHERE id = ?")
	if err != nil {
		return Payment{}, err
	}
	defer stmt.Close()

	// exécute la requête avec l'ID du paiement
	var payment Payment
	err = stmt.QueryRow(id).Scan(&payment.ID, &payment.ProductID, &payment.PricePaid, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return Payment{}, err
	}

	return payment, nil
}

// GetAllPayments récupère tous les paiements de la base de données
func (conn Connection) GetAllPayments() ([]Payment, error) {
	// prépare la requête pour récupérer tous les paiements de la base de données
	stmt, err := conn.DB.Prepare("SELECT id, product_id, price_paid, created_at, updated_at FROM payment")
	if err != nil {
		return []Payment{}, err
	}
	defer stmt.Close()

	// exécute la requête
	rows, err := stmt.Query()
	if err != nil {
		return []Payment{}, err
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var payment Payment
		err = rows.Scan(&payment.ID, &payment.ProductID, &payment.PricePaid, &payment.CreatedAt, &payment.UpdatedAt)
		if err != nil {
			return []Payment{}, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

// Controller Product

// GetAllProducts récupère tous les produits de la base de données et les renvoie en format JSON
func GetAllProducts(c *gin.Context) {
	products, err := model.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := model.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// CreateProduct crée un nouveau produit dans la base de données et le renvoie en format JSON
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := model.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

// UpdateProduct met à jour un produit dans la base de données et le renvoie en format JSON
func UpdateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := model.UpdateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct supprime un produit de la base de données
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := model.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetAllPayments récupère tous les paiements de la base de données et les renvoie en format JSON
func GetAllPayments(c *gin.Context) {
	payments, err := model.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// Controller Payment

// CreatePayment crée un nouveau paiement dans la base de données et le renvoie en format JSON
func CreatePayment(c *gin.Context) {
	var payment Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := model.CreatePayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payment.ID = id
	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// UpdatePayment met à jour un paiement dans la base de données et le renvoie en format JSON
func UpdatePayment(c *gin.Context) {
	var payment Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := model.UpdatePayment(payment); err != nil {
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
	if err := model.DeletePayment(id); err != nil {
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
	payment, err := model.GetPaymentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// RegisterRoutes enregistre les routes de l'API avec leur handler correspondant
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", GetAllProducts)
			products.POST("", CreateProduct)
			products.PUT("/:id", UpdateProduct)
			products.DELETE("/:id", DeleteProduct)
			products.GET("/:id", GetProductById)
		}
		payments := api.Group("/payments")
		{
			payments.GET("", GetAllPayments)
			payments.POST("", CreatePayment)
			payments.PUT("/:id", UpdatePayment)
			payments.DELETE("/:id", DeletePayment)
			payments.GET("/:id", GetPaymentById)
			// payments.GET("/stream", StreamPayments)
		}
	}
}

func main() {
	// initialise la connexion à la base de données
	model = Connect()
	defer model.DB.Close()

	// initialise le serveur gin
	r := gin.Default()

	// enregistre les routes de l'API
	RegisterRoutes(r)

	// démarre le serveur
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
