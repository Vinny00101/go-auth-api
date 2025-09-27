package auth_controllers

import (
	structs_Auth "go-api/dto"
	auth_service "go-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (controller *AuthController) Register_auth(context *gin.Context) {
	service := &auth_service.AuthService{}
	var request structs_Auth.Auth_User_Register

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.Create_User(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// pode incluir token JWT futuramente
	context.JSON(http.StatusOK, gin.H{
		"message": "Usuário registrado com sucesso",
		"user": gin.H{
			"name":  res.Name,
			"email": res.Email,
		},
	})

}

func (controller *AuthController) Login_auth(context *gin.Context) {
	service := &auth_service.AuthService{}
	var request structs_Auth.Auth_User_Login

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.Authenticate_User(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// pode incluir token JWT futuramente
	context.JSON(http.StatusOK, gin.H{
		"message": "Login efetuado",
		"token":   "fake-jwt-token",
		"user": gin.H{
			"name":  res.Name,
			"email": res.Email,
		},
	})
}
