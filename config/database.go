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
	queryUsers := `
CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL
);
`

	queryProducts := `
CREATE TABLE IF NOT EXISTS products (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	price VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);
`

	_, err = DB.Exec(queryUsers)
	if err != nil {
		log.Fatalf("❌ Failed to create users table: %v", err)
	}

	_, err = DB.Exec(queryProducts)
	if err != nil {
		log.Fatalf("❌ Failed to create products table: %v", err)
	}

	log.Println("✅ Connected to MySQL and initialized database")
}
