package auth_controllers

import (
	structs_Auth "go-api/dto"
	auth_service "go-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var service = &auth_service.AuthService{}

func (controller *AuthController) Register_auth(context *gin.Context) {
	var request structs_Auth.Auth_User_Register

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.Create_User(request)
	if err != nil {
		context.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}

	// pode incluir token JWT futuramente
	context.JSON(http.StatusOK, gin.H{
		"message": "Usuário registrado com sucesso",
		"token":   res.JWT,
		"user": gin.H{
			"name":  res.USER.Name,
			"email": res.USER.Email,
		},
	})

}

func (controller *AuthController) Login_auth(context *gin.Context) {
	var request structs_Auth.Auth_User_Login

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.Authenticate_User(request)
	if err != nil {
		context.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}

	// pode incluir token JWT futuramente
	context.JSON(http.StatusOK, gin.H{
		"message": "Login efetuado",
		"token":   res.JWT,
		"user": gin.H{
			"name":  res.USER.Name,
			"email": res.USER.Email,
		},
	})
}

func (controller *AuthController) Me_auth(context *gin.Context) {
	userID, exits := context.Get("user_id")

	if !exits {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	user, err := service.Get_User_By_ID(userID.(uint))

	if err != nil {
		context.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
