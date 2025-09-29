package config

import (
	"log"
	"os"
)

type Configs struct {
	GOENV       string
	SERVER_PORT string
	SECRET_KEY  string
}

var Env_Config *Configs

func Load() {
	GOENV := os.Getenv("GOENV")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	SECRET_KEY := os.Getenv("SECRET_KEY")

	if GOENV == "" || SERVER_PORT == "" || SECRET_KEY == "" {
		log.Fatal("Variáveis de ambiente não definidas corretamente")
	}

	Env_Config = &Configs{
		GOENV:       GOENV,
		SERVER_PORT: SERVER_PORT,
		SECRET_KEY:  SECRET_KEY,
	}
}
