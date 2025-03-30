package database

import (
	"log"
)

func CreatePrincipalTables() {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        first_name TEXT NOT NULL CHECK (LENGTH(first_name) <= 100),
        last_name TEXT NOT NULL CHECK (LENGTH(last_name) <= 100),
        email TEXT NOT NULL CHECK (LENGTH(email) <= 100),
        identifier TEXT NOT UNIQUE NULL CHECK (LENGTH(identifier) <= 100),
        phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
        address TEXT CHECK (LENGTH(address) <= 255),
        city TEXT CHECK (LENGTH(city) <= 100),
        country TEXT CHECK (LENGTH(country) <= 255),
        zip_code TEXT CHECK (LENGTH(zip_code) <= 20),
        password TEXT NOT NULL,
        created_at DATETIME NOT NULL DEFAULT (datetime('now')),
        updated_at DATETIME NOT NULL DEFAULT (datetime('now'))
    );

		CREATE TABLE IF NOT EXISTS entities (
        id TEXT PRIMARY KEY,
        email TEXT NOT NULL UNIQUE CHECK (LENGTH(email) <= 50),
        cuit TEXT NOT NULL UNIQUE CHECK (LENGTH(cuit) <= 20),
        name TEXT NOT NULL CHECK (LENGTH(name) <= 100),
        phone TEXT NOT NULL CHECK (LENGTH(cellphone) <= 20),
        start_activities DATE NOT NULL,
        address TEXT NOT NULL CHECK (LENGTH(address) <= 255),
        city TEXT NOT NULL CHECK (LENGTH(city) <= 100),
        country TEXT NOT NULL CHECK (LENGTH(country) <= 255),
        zip_code TEXT NOT NULL CHECK (LENGTH(zip_code) <= 20),
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at DATETIME NOT NULL DEFAULT (datetime('now')),
        updated_at DATETIME NOT NULL DEFAULT (datetime('now'))
    );

    CREATE TABLE IF NOT EXISTS establishments (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL CHECK (LENGTH(name) <= 100),
        phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
        address TEXT NOT NULL CHECK (LENGTH(address) <= 255),
        city TEXT NOT NULL CHECK (LENGTH(city) <= 100),
        country TEXT NOT NULL CHECK (LENGTH(country) <= 255),
        zip_code TEXT NOT NULL CHECK (LENGTH(zip_code) <= 20),
        date_create DATE NOT NULL,
        created_at DATETIME NOT NULL DEFAULT (datetime('now')),
        updated_at DATETIME NOT NULL DEFAULT (datetime('now')),
        connection TEXT NOT NULL,
        entity_id TEXT NOT NULL,
        FOREIGN KEY (entity_id) REFERENCES entities(id)
    );
    
    CREATE TABLE IF NOT EXISTS roles (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL UNIQUE CHECK (LENGTH(name) <= 50),
        hierarchy INTEGER UNIQUE CHECK (hierarchy < 1000)
        created_at DATETIME NOT NULL DEFAULT (datetime('now')),
        updated_at DATETIME NOT NULL DEFAULT (datetime('now'))
    );

    CREATE TABLE IF NOT EXISTS establishment_user (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        establishment_id TEXT NOT NULL,
        role_id TEXT NOT NULL,
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at DATETIME NOT NULL DEFAULT (datetime('now')),
        updated_at DATETIME NOT NULL DEFAULT (datetime('now')),
        FOREIGN KEY (user_id) REFERENCES users(id)
        FOREIGN KEY (role_id) REFERENCES roles(id)
        FOREIGN KEY (establishment_id) REFERENCES establishments(id)
    );

    CREATE TABLE IF NOT EXISTS entity_user (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        entity_id TEXT NOT NULL,
        role_id TEXT NOT NULL,
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at DATETIME NOT NULL DEFAULT (datetime('now')),
        updated_at DATETIME NOT NULL DEFAULT (datetime('now')),
        FOREIGN KEY (user_id) REFERENCES users(id)
        FOREIGN KEY (role_id) REFERENCES roles(id)
        FOREIGN KEY (entity_id) REFERENCES entities(id)
    );

    CREATE INDEX idx_name_hierarchy ON establishment_user (user_id, establishment_id);
    CREATE INDEX idx_name_hierarchy ON entity_user (user_id, entity_id);
	`

  err := ExecuteTransaction(query, nil)

  if err != nil {
    log.Fatalf("Error al crear las PrincipalTables: %v", err)
  }

  log.Println("PrincipalTables Validadas")
}

func CreateEstablishmentTables() error{
  return nil
}