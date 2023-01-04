package controller

import (
	"net/http"

	"apigolang/model"

	"github.com/gin-gonic/gin"
)

type dataLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type dataLoginResponse struct {
	Token string `json:"token"`
}

// @Summary route de connexion
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param login body dataLogin true "Login"
// @Success 200 {object} dataLoginResponse
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// authentifiez l'utilisateur et récupérez le JWT
	token, err := model.ModelInstance.AuthenticateUser(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// renvoyez le JWT dans la réponse
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary route d'inscription
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param register body dataLogin true "Login"
// @Success 200 {object} dataBoolean
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// enregistrez l'utilisateur dans la base de données
	if err := model.ModelInstance.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
