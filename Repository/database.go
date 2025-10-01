package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("../SQlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("❌ Erro ao conectar com o banco:", err)
	}

	DB = db
	log.Println("✅ Banco de dados conectado com sucesso")
}

func Migrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatal("❌ Erro ao migrar as models:", err)
	}
	log.Println("✅ Tabelas migradas com sucesso")
}
