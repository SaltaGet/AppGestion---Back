package database

import (
	"fmt"
	"database/sql"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3" 
)

var db *sql.DB

func InitDB() error {
	var err error

	dsn := os.Getenv("URI_DB")
	
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión: %w", err)
	}
	
	if err = db.Ping(); err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	CreatePrincipalTables()

	
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

func ExecuteTransaction(query string, args ...interface{}) error {
	tx, err := db.Begin()
	if err != nil {
			log.Printf("Error starting transaction: %v", err)
			return err
	}

	_, err = tx.Exec(query, args...) // Acepta argumentos para consultas parametrizadas
	if err != nil {
			log.Printf("Error executing query: %v", err)
			tx.Rollback()
			return err
	}

	err = tx.Commit()
	if err != nil {
			log.Printf("Error committing transaction: %v", err)
			return err
	}

	return nil
}

func ExecuteGroupTransactions(queries []string, args [][]interface{}) error {
	tx, err := db.Begin()
	if err != nil {
					log.Printf("Error starting transaction: %v", err)
					return err
	}

	for i, query := range queries {
					_, err := tx.Exec(query, args[i]...)
					if err != nil {
									log.Printf("Error executing query: %v", err)
									tx.Rollback()
									return err
					}
	}

	err = tx.Commit()
	if err != nil {
					log.Printf("Error committing transaction: %v", err)
					return err
	}

	return nil
}

func GetRow(query string, args ...interface{}) (*sql.Row, error) {
	row := db.QueryRow(query, args...)

	if err := row.Err(); err != nil{
		log.Printf("Error executing QueryRow: %v", err)
		return nil, err
	}

	return row, nil
}

func GetRows(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	return rows, nil
}

// // MYSQL
// func InitDB() error {
// 	var err error

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



// tables de las DB

// CREATE TABLE IF NOT EXISTS Buy (
//     id TEXT PRIMARY KEY,
//     date DATETIME NOT NULL,
//     data TEXT NOT NULL CHECK (length(data) <= 1000000)
// );

// CREATE TABLE IF NOT EXISTS Category (
//     id TEXT PRIMARY KEY,
//     name TEXT NOT NULL
// );

// CREATE TABLE IF NOT EXISTS DailyMovement (
//     id TEXT PRIMARY KEY,
//     product_id TEXT NOT NULL,
//     created_at DATETIME NOT NULL,
//     updated_at DATETIME NOT NULL,
//     cost REAL NOT NULL,
//     price REAL NOT NULL,
//     movement INTEGER NOT NULL,
//     FOREIGN KEY (product_id) REFERENCES Product(id)
// );

// CREATE TABLE IF NOT EXISTS Product (
//     id TEXT PRIMARY KEY,
//     sku TEXT NOT NULL,
//     name TEXT NOT NULL,
//     unit_measure INTEGER NOT NULL,
//     category_id TEXT NOT NULL,
//     cost REAL NOT NULL,
//     price REAL NOT NULL,
//     date_from DATETIME NOT NULL,
//     date_to DATETIME NOT NULL,
//     created_at DATETIME NOT NULL,
//     updated_at DATETIME NOT NULL,
//     FOREIGN KEY (category_id) REFERENCES Category(id)
// );

// CREATE TABLE IF NOT EXISTS ProductPriceHistory (
//     id TEXT PRIMARY KEY,
//     data TEXT NOT NULL CHECK (length(data) <= 1000000)
// );

// CREATE TABLE IF NOT EXISTS Sale (
//     id TEXT PRIMARY KEY,
//     date DATETIME NOT NULL,
//     data TEXT NOT NULL CHECK (length(data) <= 1000000)
// );

// CREATE TABLE IF NOT EXISTS Discontinued (
//     id TEXT PRIMARY KEY,
//     date DATETIME NOT NULL,
//     data TEXT NOT NULL CHECK (length(data) <= 1000000)
// );

// CREATE TABLE IF NOT EXISTS Stock (
//     id TEXT PRIMARY KEY,
//     product_id TEXT NOT NULL,
//     stock REAL NOT NULL,
//     FOREIGN KEY (product_id) REFERENCES Product(id)
// );

// CREATE TABLE IF NOT EXISTS User (
//     id TEXT PRIMARY KEY,
//     username TEXT NOT NULL,
//     first_name TEXT NOT NULL,
//     last_name TEXT NOT NULL,
//     cellphone TEXT NOT NULL,
//     email TEXT NOT NULL,
//     password TEXT NOT NULL,
//     created_at DATETIME NOT NULL,
//     updated_at DATETIME NOT NULL,
//     is_active BOOLEAN NOT NULL
// );