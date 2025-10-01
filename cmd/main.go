package main

import (
	"fmt"
	"go-api/config"
	"go-api/middlewares"
	user_model "go-api/model"
	database "go-api/repository"
	auth_routes "go-api/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	configs, err := config.Load()
	if err != nil {
		log.Fatal("Erro ao carregar configurações")
	}
	config.Env_Config = configs

	if config.Env_Config.GOENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()

	database.Connect()
	database.Migrate(&user_model.User{})

	server.Use(middlewares.SanitizerMiddleware())

	api := server.Group("/api")
	{
		auth_routes.Setup_routes_auth(api)
	}

	addr := fmt.Sprintf(":%s", config.Env_Config.SERVER_PORT)
	srv := &http.Server{
		Addr:    addr,
		Handler: server,
	}

	log.Printf("Servidor rodando na porta %s (env: %s)", config.Env_Config.SERVER_PORT, config.Env_Config.GOENV)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Erro ao inicia o servidor")
	}
}
