package auth_service

// Service contém a lógica de autenticação, como registro e login de usuários.

import (
	"errors"

	structs_Auth "go-api/dto"
	user_model "go-api/model"
	database "go-api/repository"
	"go-api/service/utils"

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

func (Service *AuthService) Create_User(request structs_Auth.Auth_User_Register) (*DataReturn, error) {
	// Lógica para criar um usuário no banco de dados
	if err := validate.Struct(request); err != nil {
		return nil, err
	}

	if request.Name == "" || request.Email == "" || request.Password == "" {
		return nil, errors.New("todos os campos são obrigatórios")
	}

	// Sanitiza os inputs
	request.Name = Service.SanitizeInput(request.Name)
	request.Email = Service.SanitizeInput(request.Email)
	request.Password = Service.SanitizeInput(request.Password)

	if database.DB.Where("email = ?", request.Email).First(&user_model.User{}).Error == nil {
		return nil, errors.New("email já cadastrado")
	}

	if len(request.Password) < 8 {
		return nil, errors.New("senha deve ter pelo menos 8 caracteres")
	}

	Hash, _ := Service.HashPassword(request.Password)

	user := user_model.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: Hash,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, errors.New("erro ao criar usuário")
	}

	token, err := utils.JwtGeneration(int(user.ID))
	if err != nil {
		return nil, errors.New("erro ao gerar token")
	}

	Data_user := DataReturn{
		USER: &user,
		JWT:  token,
	}

	return &Data_user, nil
}

func (Service *AuthService) Authenticate_User(request structs_Auth.Auth_User_Login) (*DataReturn, error) {
	// Lógica para autenticar um usuário no banco de dados
	if err := validate.Struct(request); err != nil {
		return nil, err
	}

	if request.Email == "" || request.Password == "" {
		return nil, errors.New("todos os campos são obrigatórios")
	}

	// Sanitiza os inputs
	request.Email = Service.SanitizeInput(request.Email)
	request.Password = Service.SanitizeInput(request.Password)

	var user user_model.User
	if database.DB.Where("email = ?", request.Email).First(&user).Error != nil {
		return nil, errors.New("email não cadastrado")
	}

	if !Service.VerifyPassword(user.PasswordHash, request.Password) {
		return nil, errors.New("senha incorreta")
	}

	token, err := utils.JwtGeneration(int(user.ID))
	if err != nil {
		return nil, errors.New("erro ao gerar token")
	}

	Data_user := DataReturn{
		USER: &user,
		JWT:  token,
	}

	return &Data_user, nil
}
