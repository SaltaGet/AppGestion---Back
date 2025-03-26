package database

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) error {
	var err error
	
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
        id TEXT PRIMARY KEY,
        email TEXT NOT NULL CHECK (LENGTH(email) <= 50),
        cuit TEXT NOT NULL CHECK (LENGTH(cuit) <= 20),
        name TEXT NOT NULL CHECK (LENGTH(name) <= 100),
        password TEXT NOT NULL,
        cellphone TEXT NOT NULL CHECK (LENGTH(cellphone) <= 20),
        role INTEGER NOT NULL DEFAULT 1,
        is_active BOOLEAN NOT NULL DEFAULT FALSE,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    );

    CREATE TABLE IF NOT EXISTS establishments (
        id TEXT PRIMARY KEY,
        company_name TEXT NOT NULL CHECK (LENGTH(company_name) <= 100),
        address TEXT NOT NULL CHECK (LENGTH(address) <= 255),
        city TEXT NOT NULL CHECK (LENGTH(city) <= 100),
        country TEXT NOT NULL CHECK (LENGTH(country) <= 255),
        zip_code TEXT NOT NULL CHECK (LENGTH(zip_code) <= 20),
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL,
        client_id TEXT NOT NULL,
        connection TEXT NOT NULL,
        FOREIGN KEY (client_id) REFERENCES Client(id)
    );
	`)

	if err != nil{
		log.Fatalf("Error al crear las tablas: %v", err)
	}

	return nil
}