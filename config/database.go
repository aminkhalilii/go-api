package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		AppConfig.DBUser, AppConfig.DBPass, AppConfig.DBHost, AppConfig.DBPort,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", AppConfig.DBName))
	if err != nil {
		log.Fatal("❌ Failed to create DB:", err)
	}

	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		AppConfig.DBUser, AppConfig.DBPass, AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBName,
	)
	DB, err = sql.Open("mysql", dsnWithDB)
	if err != nil {
		log.Fatal(err)
	}

	// Migration ساده
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("❌ Failed to create users table:", err)
	}

	log.Println("✅ Connected to MySQL and initialized database")
}
