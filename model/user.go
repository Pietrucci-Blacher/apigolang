package model

import (
	"fmt"
	"os"
	"time"

	"apigolang/utils"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (conn *Connection) CreateUser(user *User) error {
	stmt, err := conn.DB.Prepare("INSERT INTO user (username, password, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword := utils.Sha512(user.Password)

	_, err = stmt.Exec(user.Username, hashedPassword, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (conn *Connection) AuthenticateUser(user *User) (string, error) {
	var storedUser User
	query := `SELECT * FROM user WHERE username = ?`
	err := conn.DB.QueryRow(query, user.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password, &storedUser.CreatedAt, &storedUser.UpdatedAt)
	if err != nil {
		return "", err
	}

	hashedPassword := utils.Sha512(user.Password)

	if hashedPassword != storedUser.Password {
		return "", fmt.Errorf("Invalid password")
	}

	// créez un JWT signé avec les informations de l'utilisateur
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
	})

	JWT_SECRET := os.Getenv("JWT_SECRET")

	signedToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
