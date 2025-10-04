package main

import (
	"fmt"
	"go-api/config"
	"go-api/middlewares"
	"go-api/model"
	"go-api/routes"
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

	config.Connect()
	config.Migrate(&model.User{})
	server.Use(middlewares.SanitizerMiddleware())

	api := server.Group("/api")
	{
		routes.Setup_routes_auth(api)
		routes.Setup_routes_user(api)
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
