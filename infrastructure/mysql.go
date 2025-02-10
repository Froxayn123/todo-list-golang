package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	DbName string
}


func NewDbConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}


	config := Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
	}

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.DbName,)

	db, err := sql.Open("mysql", mysqlInfo)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}


	return db
	
}