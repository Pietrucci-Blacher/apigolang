package model

import "time"

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
