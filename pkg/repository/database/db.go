package database

import (
	"fmt"
	"database/sql"
	"log"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3" 
	"sync"
	"time"
	"os"
)

var (
	mainDB     *sql.DB
	tenantDBs  = make(map[string]*sql.DB)
	mu         sync.RWMutex
)

func ConectDB(uri string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión: %w", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(3 * time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)
	
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	CreatePrincipalTables(db)
	CreateRoles(db)
	CreateAdmin(db)

	mainDB = db
	
	return db, nil
}

func GetTenantDB(tenantID string) (*sql.DB, error) {
	if tenantID == "default" {
		return mainDB, nil
	}

	mu.RLock()
	db, ok := tenantDBs[tenantID]
	mu.RUnlock()

	if ok {
		return db, nil
	}

	mu.Lock()
	defer mu.Unlock()

	// Doble verificación
	if db, ok := tenantDBs[tenantID]; ok {
		return db, nil
	}

	uri := getTenantURI(tenantID)
	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(3 * time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	tenantDBs[tenantID] = db
	return db, nil
}

func getTenantURI(tenantID string) string {
	// Implementa tu lógica para obtener la URI del tenant
	// Por ahora usamos la misma que la principal
	return os.Getenv("URI_DB")
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}

func CloseAllTenantDBs() {
	mu.Lock()
	defer mu.Unlock()

	for tenant, db := range tenantDBs {
		if err := db.Close(); err != nil {
			log.Printf("Error cerrando conexión tenant %s: %v", tenant, err)
		}
		delete(tenantDBs, tenant)
	}
}

func ExecuteTransaction(db *sql.DB, query string, args ...interface{}) error {
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

func ExecuteGroupTransactions(db *sql.DB, queries []string, args [][]interface{}) error {
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

func GetRow(db *sql.DB, query string, args ...interface{}) *sql.Row  {
	row := db.QueryRow(query, args...)
	return row
}

func GetRows(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	return rows, nil
}






















// Primera conexion
// func ConectDB(uri string) (*sql.DB, error) {
// 	db, err := sql.Open("sqlite3", uri)
// 	if err != nil {
// 		return nil, fmt.Errorf("error al abrir la conexión: %w", err)
// 	}
	
// 	if err = db.Ping(); err != nil {
// 		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
// 	}

// 	CreatePrincipalTables(db)
// 	CreateRoles(db)
// 	CreateAdmin(db)
	
// 	return db, nil
// }


// func CloseDB(db *sql.DB) error {
// 	if db != nil {
// 		return db.Close()
// 	}
// 	return nil
// }

// func ExecuteTransaction(db *sql.DB, query string, args ...interface{}) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 			log.Printf("Error starting transaction: %v", err)
// 			return err
// 	}

// 	_, err = tx.Exec(query, args...) // Acepta argumentos para consultas parametrizadas
// 	if err != nil {
// 			log.Printf("Error executing query: %v", err)
// 			tx.Rollback()
// 			return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 			log.Printf("Error committing transaction: %v", err)
// 			return err
// 	}

// 	return nil
// }

// func ExecuteGroupTransactions(db *sql.DB, queries []string, args [][]interface{}) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 					log.Printf("Error starting transaction: %v", err)
// 					return err
// 	}

// 	for i, query := range queries {
// 					_, err := tx.Exec(query, args[i]...)
// 					if err != nil {
// 									log.Printf("Error executing query: %v", err)
// 									tx.Rollback()
// 									return err
// 					}
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 					log.Printf("Error committing transaction: %v", err)
// 					return err
// 	}

// 	return nil
// }

// func GetRow(db *sql.DB, query string, args ...interface{}) *sql.Row  {
// 	row := db.QueryRow(query, args...)
// 	return row
// }

// func GetRows(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
// 	rows, err := db.Query(query, args...)

// 	if err != nil {
// 		log.Printf("Error executing query: %v", err)
// 		return nil, err
// 	}

// 	return rows, nil
// }


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