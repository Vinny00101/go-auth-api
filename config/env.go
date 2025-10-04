package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	GOENV       string
	SERVER_PORT string
	SECRET_KEY  string
}

var Env_Config *Configs

func Load() (Config *Configs, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Aviso: não foi possível carregar .env, usando variáveis do sistema", err)
	}

	GOENV := os.Getenv("GOENV")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	SECRET_KEY := os.Getenv("SECRET_KEY")

	if GOENV == "" || SERVER_PORT == "" || SECRET_KEY == "" {
		log.Fatal("Variáveis de ambiente não definidas corretamente")
	}

	Config = &Configs{
		GOENV:       GOENV,
		SERVER_PORT: SERVER_PORT,
		SECRET_KEY:  SECRET_KEY,
	}
	return Config, nil
}
