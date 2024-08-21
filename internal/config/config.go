package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	p, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	p = filepath.Dir(p)

	err = godotenv.Load(path.Join(p, ".env"))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}

	userDb := os.Getenv("DB_USER")
	pwdDb := os.Getenv("DB_PASSWORD")
	hostDb := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	strConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userDb, pwdDb, hostDb, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(strConn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
}
