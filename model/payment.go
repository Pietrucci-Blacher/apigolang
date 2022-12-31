package model

import "time"

// Payment est la structure de données pour un paiement
type Payment struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	PricePaid float64   `json:"price_paid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreatePayment crée un nouveau paiement dans la base de données et renvoie l'ID du paiement
func (conn *Connection) CreatePayment(payment *Payment) (int, error) {
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
func (conn *Connection) UpdatePayment(payment *Payment) error {
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
func (conn *Connection) DeletePayment(id int) error {
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
func (conn *Connection) GetPaymentById(id int) (Payment, error) {
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
func (conn *Connection) GetAllPayments() ([]Payment, error) {
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
