package database

import (
	"database/sql"
	"fmt"
	"os"
)

func PrepareDB(uri string) error {
	// Abrir conexión (esto no crea conexión real aún)
	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		return fmt.Errorf("error inicializando db: %w", err)
	}
	defer db.Close()

	// Ejecutar algo trivial para forzar la creación del archivo
	if _, err = db.Exec("PRAGMA user_version"); err != nil {
		_ = os.Remove(uri)
		return fmt.Errorf("error inicial: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		_ = os.Remove(uri)
		return fmt.Errorf("error al iniciar transacción: %w", err)
	}

	// Crear las tablas por defecto
	createTables := `
		CREATE TABLE IF NOT EXISTS buy (
				id TEXT PRIMARY KEY,
				date DATETIME NOT NULL,
				data TEXT NOT NULL CHECK (length(data) <= 1000000)
		);

		CREATE TABLE IF NOT EXISTS category (
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS daily_movement (
				id TEXT PRIMARY KEY,
				product_id TEXT NOT NULL,
				created_at DATETIME NOT NULL,
				updated_at DATETIME NOT NULL,
				cost REAL NOT NULL,
				price REAL NOT NULL,
				movement INTEGER NOT NULL,
				FOREIGN KEY (product_id) REFERENCES Product(id)
		);

		CREATE TABLE IF NOT EXISTS product (
				id TEXT PRIMARY KEY,
				sku TEXT NOT NULL,
				name TEXT NOT NULL,
				unit_measure INTEGER NOT NULL,
				category_id TEXT NOT NULL,
				cost REAL NOT NULL,
				price REAL NOT NULL,
				date_from DATETIME NOT NULL,
				date_to DATETIME NOT NULL,
				created DATETIME NOT NULL,
				updated DATETIME NOT NULL,
				FOREIGN KEY (category_id) REFERENCES Category(id)
		);

		CREATE TABLE IF NOT EXISTS product_price_history (
				id TEXT PRIMARY KEY,
				data TEXT NOT NULL CHECK (length(data) <= 1000000)
		);

		CREATE TABLE IF NOT EXISTS sale (
				id TEXT PRIMARY KEY,
				date DATETIME NOT NULL,
				data TEXT NOT NULL CHECK (length(data) <= 1000000)
		);

		CREATE TABLE IF NOT EXISTS discontinued (
				id TEXT PRIMARY KEY,
				date DATETIME NOT NULL,
				data TEXT NOT NULL CHECK (length(data) <= 1000000)
		);

		CREATE TABLE IF NOT EXISTS stock (
				id TEXT PRIMARY KEY,
				product_id TEXT NOT NULL,
				stock REAL NOT NULL,
				FOREIGN KEY (product_id) REFERENCES Product(id)
		);

		CREATE TABLE IF NOT EXISTS roles (
				id TEXT PRIMARY KEY,
				date DATETIME NOT NULL,
				data TEXT NOT NULL CHECK (length(data) <= 1000000)
		);
	`

	_, err = tx.Exec(createTables)
	if err != nil {
		tx.Rollback()
		_ = os.Remove(uri)
		return fmt.Errorf("error al crear tablas: %w", err)
	}

	// Insertar datos por defecto
	// _, err = tx.Exec(`INSERT OR IGNORE INTO roles (name) VALUES ('admin')`)
	// if err != nil {
	// 	tx.Rollback()
	// 	_ = os.Remove(uri)
	// 	return fmt.Errorf("error al insertar datos por defecto: %w", err)
	// }

	// Confirmar cambios
	if err = tx.Commit(); err != nil {
		_ = os.Remove(uri)
		return fmt.Errorf("error al confirmar la transacción: %w", err)
	}

	return nil
}
