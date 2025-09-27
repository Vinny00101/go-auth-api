package main

import (
	user_model "go-api/model"
	database "go-api/repository"
	auth_routes "go-api/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	database.Connect()
	database.Migrate(&user_model.User{})

	api := server.Group("/api")
	{
		auth_routes.Setup_routes_auth(api)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}
	srv.ListenAndServe()
}
