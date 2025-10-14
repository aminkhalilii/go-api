// # تنظیمات دیتابیس و متغیرهای محیطی (DB_USER, DB_PASS, DB_NAME, PORT)

package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitMysql() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", dbUser, dbPass, dbHost, dbPort)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName))
	if err != nil {
		log.Fatal("❌ Failed to create database:", err)
	}
	//test
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Connected to MySQL successfully \n")

}
