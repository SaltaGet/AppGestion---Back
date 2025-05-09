package database

import (
	"database/sql"
	"log"
)

// func CreatePrincipalTables(db *sql.DB) {
// 	query := `
//     CREATE TABLE IF NOT EXISTS users (
//         id TEXT PRIMARY KEY,
//         first_name TEXT NOT NULL CHECK (LENGTH(first_name) <= 100),
//         last_name TEXT NOT NULL CHECK (LENGTH(last_name) <= 100),
//         email TEXT NOT NULL CHECK (LENGTH(email) <= 100),
//         identifier TEXT NOT NULL UNIQUE CHECK (LENGTH(identifier) <= 100),
//         phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
//         address TEXT CHECK (LENGTH(address) <= 255),
//         city TEXT CHECK (LENGTH(city) <= 100),
//         country TEXT CHECK (LENGTH(country) <= 255),
//         zip_code TEXT CHECK (LENGTH(zip_code) <= 20),
//         password TEXT NOT NULL,
//         created DATETIME NOT NULL DEFAULT (datetime('now')),
//         updated DATETIME NOT NULL DEFAULT (datetime('now'))
//     );

// 		CREATE TABLE IF NOT EXISTS entities (
//         id TEXT PRIMARY KEY,
//         email TEXT NOT NULL UNIQUE CHECK (LENGTH(email) <= 50),
//         cuit TEXT NOT NULL UNIQUE CHECK (LENGTH(cuit) <= 20),
//         name TEXT NOT NULL CHECK (LENGTH(name) <= 100),
//         phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
//         start_activities DATE NOT NULL,
//         address TEXT NOT NULL CHECK (LENGTH(address) <= 255),
//         city TEXT NOT NULL CHECK (LENGTH(city) <= 100),
//         country TEXT NOT NULL CHECK (LENGTH(country) <= 255),
//         zip_code TEXT NOT NULL CHECK (LENGTH(zip_code) <= 20),
//         is_active BOOLEAN NOT NULL DEFAULT TRUE,
//         created DATETIME NOT NULL DEFAULT (datetime('now')),
//         updated DATETIME NOT NULL DEFAULT (datetime('now'))
//     );

//     CREATE TABLE IF NOT EXISTS establishments (
//         id TEXT PRIMARY KEY,
//         name TEXT NOT NULL CHECK (LENGTH(name) <= 100),
//         phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
//         address TEXT NOT NULL CHECK (LENGTH(address) <= 255),
//         city TEXT NOT NULL CHECK (LENGTH(city) <= 100),
//         country TEXT NOT NULL CHECK (LENGTH(country) <= 255),
//         zip_code TEXT NOT NULL CHECK (LENGTH(zip_code) <= 20),
//         date_create DATE NOT NULL,
//         created DATETIME NOT NULL DEFAULT (datetime('now')),
//         updated DATETIME NOT NULL DEFAULT (datetime('now')),
//         connection TEXT NOT NULL,
//         entity_id TEXT NOT NULL,
//         FOREIGN KEY (entity_id) REFERENCES entities(id)
//     );
    
//     CREATE TABLE IF NOT EXISTS roles (
//         id TEXT PRIMARY KEY,
//         name TEXT NOT NULL UNIQUE CHECK (LENGTH(name) <= 50),
//         hierarchy INTEGER UNIQUE CHECK (hierarchy < 1000),
//         created DATETIME NOT NULL DEFAULT (datetime('now')),
//         updated DATETIME NOT NULL DEFAULT (datetime('now'))
//     );

//     CREATE TABLE IF NOT EXISTS establishment_user (
//         id TEXT PRIMARY KEY,
//         user_id TEXT NOT NULL,
//         establishment_id TEXT NOT NULL,
//         is_active BOOLEAN NOT NULL DEFAULT TRUE,
//         created DATETIME NOT NULL DEFAULT (datetime('now')),
//         updated DATETIME NOT NULL DEFAULT (datetime('now')),
//         FOREIGN KEY (user_id) REFERENCES users(id)
//         FOREIGN KEY (establishment_id) REFERENCES establishments(id)
//     );
// cannot use entityRepo (variable of type *"api-stock/pkg/repository/entity".Repository) as *ports.EntityRepository value in struct literal: *"api-stock/pkg/repository/entity".Repository does not implement *ports.EntityRepository (type *ports.EntityRepository is pointer to interface, not interface)compilerInvalidIfaceAssign
// var entityRepo *entRep.Repository
//     CREATE TABLE IF NOT EXISTS entity_user (
//         id TEXT PRIMARY KEY,
//         user_id TEXT NOT NULL,
//         entity_id TEXT NOT NULL,
//         role_id TEXT NOT NULL,
//         is_active BOOLEAN NOT NULL DEFAULT TRUE,
//         created DATETIME NOT NULL DEFAULT (datetime('now')),
//         updated DATETIME NOT NULL DEFAULT (datetime('now')),
//         FOREIGN KEY (user_id) REFERENCES users(id)
//         FOREIGN KEY (role_id) REFERENCES roles(id)
//         FOREIGN KEY (entity_id) REFERENCES entities(id)
//     );

//     CREATE INDEX IF NOT EXISTS idx_establishment_user ON establishment_user (user_id, establishment_id);
//     CREATE INDEX IF NOT EXISTS idx_entity_user ON entity_user (user_id, entity_id);
// 	`

//   err := ExecuteTransaction(db,query, nil)

//   if err != nil {
//     log.Fatalf("Error al crear las PrincipalTables: %v", err)
//   }

//   log.Println("PrincipalTables Validadas")
// }

func CreatePrincipalTables(db *sql.DB) {
  query := `
  CREATE TABLE IF NOT EXISTS users (
      id TEXT PRIMARY KEY,
      first_name TEXT NOT NULL CHECK (LENGTH(first_name) <= 100),
      last_name TEXT NOT NULL CHECK (LENGTH(last_name) <= 100),
      email TEXT NOT NULL CHECK (LENGTH(email) <= 100),
      identifier TEXT NOT NULL UNIQUE CHECK (LENGTH(identifier) <= 100),
      phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
      address TEXT CHECK (LENGTH(address) <= 255),
      city TEXT CHECK (LENGTH(city) <= 100),
      country TEXT CHECK (LENGTH(country) <= 255),
      zip_code TEXT CHECK (LENGTH(zip_code) <= 20),
      password TEXT NOT NULL,
      is_admin BOOLEAN NOT NULL DEFAULT FALSE,
      created DATETIME NOT NULL DEFAULT (datetime('now')),
      updated DATETIME NOT NULL DEFAULT (datetime('now'))
  );

  CREATE TABLE IF NOT EXISTS entities (
      id TEXT PRIMARY KEY,
      email TEXT NOT NULL UNIQUE CHECK (LENGTH(email) <= 50),
      cuit TEXT NOT NULL UNIQUE CHECK (LENGTH(cuit) <= 20),
      name TEXT NOT NULL CHECK (LENGTH(name) <= 100),
      phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
      start_activities DATE NOT NULL,
      address TEXT NOT NULL CHECK (LENGTH(address) <= 255),
      city TEXT NOT NULL CHECK (LENGTH(city) <= 100),
      country TEXT NOT NULL CHECK (LENGTH(country) <= 255),
      zip_code TEXT NOT NULL CHECK (LENGTH(zip_code) <= 20),
      is_active BOOLEAN NOT NULL DEFAULT TRUE,
      created DATETIME NOT NULL DEFAULT (datetime('now')),
      updated DATETIME NOT NULL DEFAULT (datetime('now'))
  );

  CREATE TABLE IF NOT EXISTS establishments (
      id TEXT PRIMARY KEY,
      name TEXT NOT NULL CHECK (LENGTH(name) <= 100),
      email TEXT NOT NULL UNIQUE CHECK (LENGTH(email) <= 50),
      phone TEXT NOT NULL CHECK (LENGTH(phone) <= 20),
      address TEXT NOT NULL CHECK (LENGTH(address) <= 255),
      city TEXT NOT NULL CHECK (LENGTH(city) <= 100),
      country TEXT NOT NULL CHECK (LENGTH(country) <= 255),
      zip_code TEXT NOT NULL CHECK (LENGTH(zip_code) <= 20),
      date_create DATE NOT NULL,
      created DATETIME NOT NULL DEFAULT (datetime('now')),
      updated DATETIME NOT NULL DEFAULT (datetime('now')),
      connection TEXT NOT NULL,
      entity_id TEXT NOT NULL,
      FOREIGN KEY (entity_id) REFERENCES entities(id)
  );
  
  CREATE TABLE IF NOT EXISTS roles (
      id TEXT PRIMARY KEY,
      name TEXT NOT NULL UNIQUE CHECK (LENGTH(name) <= 50),
      hierarchy INTEGER UNIQUE CHECK (hierarchy < 10),
      created DATETIME NOT NULL DEFAULT (datetime('now')),
      updated DATETIME NOT NULL DEFAULT (datetime('now'))
  );

  CREATE TABLE IF NOT EXISTS establishment_user (
      id TEXT PRIMARY KEY,
      user_id TEXT NOT NULL,
      establishment_id TEXT NOT NULL,
      is_active BOOLEAN NOT NULL DEFAULT TRUE,
      created DATETIME NOT NULL DEFAULT (datetime('now')),
      updated DATETIME NOT NULL DEFAULT (datetime('now')),
      FOREIGN KEY (user_id) REFERENCES users(id),
      FOREIGN KEY (establishment_id) REFERENCES establishments(id)
  );

  CREATE TABLE IF NOT EXISTS entity_user (
      id TEXT PRIMARY KEY,
      user_id TEXT NOT NULL,
      entity_id TEXT NOT NULL,
      role_id TEXT NOT NULL,
      is_active BOOLEAN NOT NULL DEFAULT TRUE,
      created DATETIME NOT NULL DEFAULT (datetime('now')),
      updated DATETIME NOT NULL DEFAULT (datetime('now')),
      FOREIGN KEY (user_id) REFERENCES users(id),
      FOREIGN KEY (role_id) REFERENCES roles(id),
      FOREIGN KEY (entity_id) REFERENCES entities(id)
  );

  CREATE INDEX IF NOT EXISTS idx_establishment_user ON establishment_user (user_id, establishment_id);
  CREATE INDEX IF NOT EXISTS idx_entity_user ON entity_user (user_id, entity_id);
  `

  err := ExecuteTransaction(db, query, nil)

  if err != nil {
      log.Fatalf("Error al crear las PrincipalTables: %v", err)
  }

  log.Println("PrincipalTables Validadas")
}

func CreateEstablishmentTables() error{
  return nil
}