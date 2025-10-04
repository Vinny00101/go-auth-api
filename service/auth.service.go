package auth_service

// Service contém a lógica de autenticação, como registro e login de usuários.

import (
	"net/http"

	"go-api/config/err"
	structs_Auth "go-api/dto"
	user_model "go-api/model"
	"go-api/repository"
	"go-api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type AuthService struct{}
type DataReturn struct {
	USER *user_model.User
	JWT  string
}

func (Service *AuthService) SanitizeInput(input string) string {
	policy := bluemonday.UGCPolicy()
	clean := policy.Sanitize(input)
	return clean
}

func (Service *AuthService) HashPassword(password string) (string, error) {
	Hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(Hash), err
}

func (Service *AuthService) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (Service *AuthService) Create_User(request structs_Auth.Auth_User_Register) (*DataReturn, *err.HttpError) {
	// Lógica para criar um usuário no banco de dados
	if errs := validate.Struct(request); errs != nil {
		return nil, err.NewErrorHttp(http.StatusBadRequest, "Campos invalidos")
	}

	if request.Name == "" || request.Email == "" || request.Password == "" {
		return nil, err.NewErrorHttp(http.StatusBadRequest, "todos os campos são obrigatórios")
	}

	// Sanitiza os inputs
	request.Name = Service.SanitizeInput(request.Name)
	request.Email = Service.SanitizeInput(request.Email)
	request.Password = Service.SanitizeInput(request.Password)

	if repository.Where_user_verify(request.Email) {
		return nil, err.NewErrorHttp(http.StatusFound, "email já cadastrado")
	}

	if len(request.Password) < 8 {
		return nil, err.NewErrorHttp(http.StatusLengthRequired, "senha deve ter pelo menos 8 caracteres")
	}

	Hash, _ := Service.HashPassword(request.Password)

	user := user_model.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: Hash,
	}

	if !repository.Create_User(&user) {
		return nil, err.NewErrorHttp(http.StatusInternalServerError, "erro ao criar usuário")
	}

	token, errs := utils.JwtGeneration(int(user.ID), user.Email)
	if errs != nil {
		return nil, err.NewErrorHttp(http.StatusInternalServerError, "erro ao gerar token")
	}

	Data_user := DataReturn{
		USER: &user,
		JWT:  token,
	}

	return &Data_user, nil
}

func (Service *AuthService) Authenticate_User(request structs_Auth.Auth_User_Login) (*DataReturn, *err.HttpError) {
	// Lógica para autenticar um usuário no banco de dados
	if errs := validate.Struct(request); errs != nil {
		return nil, err.NewErrorHttp(http.StatusBadRequest, "Campos invalidos")
	}
	if request.Email == "" || request.Password == "" {
		return nil, err.NewErrorHttp(http.StatusBadRequest, "todos os campos são obrigatórios")
	}

	// Sanitiza os inputs
	request.Email = Service.SanitizeInput(request.Email)
	request.Password = Service.SanitizeInput(request.Password)

	user := repository.Get_user_by_email(request.Email)
	if user == nil {
		return nil, err.NewErrorHttp(http.StatusBadRequest, "usuário não encontrado")
	}

	if !Service.VerifyPassword(user.PasswordHash, request.Password) {
		return nil, err.NewErrorHttp(http.StatusUnauthorized, "senha incorreta")
	}

	token, errs := utils.JwtGeneration(int(user.ID), user.Email)
	if errs != nil {
		return nil, err.NewErrorHttp(http.StatusInternalServerError, "erro ao gerar token")
	}

	Data_user := DataReturn{
		USER: user,
		JWT:  token,
	}

	return &Data_user, nil
}

func (Service *AuthService) Get_User_By_ID(userID uint) (*structs_Auth.Auth_User_Response, *err.HttpError) {
	user := repository.Get_User_By_ID(userID)
	if user == nil {
		return nil, err.NewErrorHttp(http.StatusNotFound, "usuário não encontrado")
	}
	user_response := &structs_Auth.Auth_User_Response{
		Name:  user.Name,
		Email: user.Email,
	}
	return user_response, nil
}
