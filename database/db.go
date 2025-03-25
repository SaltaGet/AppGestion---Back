package database

import (
	"fmt"
	"database/sql"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3" 
)

var db *sql.DB

func InitDB() error {
	var err error
	
	err = godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error cargando el archivo .env")
	}

	dsn := os.Getenv("URI_DB")
	fmt.Println(dsn)
	
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión: %w", err)
	}
	
	if err = db.Ping(); err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	CreateTables(db)
	
	return nil
}


func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// GetDB retorna la conexión a la base de datos
func GetDB() *sql.DB {
	return db
}

func CreateTables(db *sql.DB) error {
	var err error
	
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS products (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			price DECIMAL(10,2) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			total DECIMAL(10,2) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`)

	if err != nil{
		log.Fatal("Error al crear las tablas")
	}

	return nil
}

// // MYSQL
// func InitDB() error {
// 	var err error
	
// 	err = godotenv.Load()
// 	if err != nil {
// 		log.Fatal("❌ Error cargando el archivo .env")
// 	}

// 	dsn := os.Getenv("URI_DB")
// 	fmt.Println(dsn)
	
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		return fmt.Errorf("error al abrir la conexión: %w", err)
// 	}
	
// 	if err = db.Ping(); err != nil {
// 		return fmt.Errorf("error al conectar con la base de datos: %w", err)
// 	}
	
// 	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS fiber")
// 	if err != nil {
// 		return fmt.Errorf("error al crear la base de datos: %w", err)
// 	}
	
// 	_, err = db.Exec("USE fiber")
// 	if err != nil {
// 		return fmt.Errorf("error al seleccionar la base de datos: %w", err)
// 	}

// 	return nil
// }

// func CloseDB() error {
// 	if db != nil {
// 		return db.Close()
// 	}
// 	return nil
// }

// func GetDB() *sql.DB {
// 	return db
// }
