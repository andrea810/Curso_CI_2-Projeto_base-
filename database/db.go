package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	// Lê as variáveis de ambiente. Se estiverem vazias, usa um valor padrão.
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "root"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "root"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5435" // Porta exposta no seu docker-compose para a máquina host
	}

	// Monta a string de conexão dinamicamente usando as variáveis (agora elas estão sendo "usadas")
	stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	DB.AutoMigrate(&models.Aluno{})
}
