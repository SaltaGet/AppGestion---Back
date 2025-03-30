package database

import (
	ent "api-stock/models/entity"
	"api-stock/utils"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

func createRoles() error {
	queries := []string{}
	args := [][]interface{}{}

	initRoles := map[string]int{
		"ADMIN":           999,
		"GENERAL_MANAGER": 800,
		"SUPERVISOR":      600,
		"MANAGEMENT":      400,
		"SPECIAL":         200,
		"EMPLOYEE":        100,
	}

	for role, hierarchy := range initRoles {
		newId := uuid.NewString()
		query := `INSERT INTO roles (id, name, hierarchy) VALUES (?, ?, ?);`
		queries = append(queries, query)
		args = append(args, []interface{}{newId, role, hierarchy})
	}

	err := ExecuteGroupTransactions(queries, args)
	if err != nil {
		log.Fatalf("Error inserting roles: %v", err)
	}

	return nil

}

func createAdmin() error {
	query := `SELECT * FROM entities WHERE cuit = ?)`

	row, err := GetRow(query, os.Getenv("CUIT_ADMIN"))

	if err != nil {
		log.Fatalf("Error al obtener admin: %v", err)
	}

	var entity ent.Entity
	err = row.Scan(&entity)

	if err != nil {
		log.Fatalf("Error al obtener entity: %v", err)
	}

	if entity.CUIT != "" {
		return nil
	}

	query = `SELECT id FROM roles WHERE name = ?`
	row, err = GetRow(query, "ADMIN")

	if err != nil {
		log.Fatalf("Error al obtener role: %v", err)
	}

	var roleId string
	err = row.Scan(&roleId)

	if roleId == "" {
		log.Fatalf("Error al obtener rol de admin")
	}

	startActivitiesAdmin := os.Getenv("START_ACTIVITIES_ADMIN")
	layout := "2006/01/02"
	parsedTime, err := time.Parse(layout, startActivitiesAdmin)
	if err != nil {
		log.Fatalf("Error al generar la fecha de inicio de actividades")
	}

	userId := uuid.NewString()
	entityId := uuid.NewString()
	entityUserId := uuid.NewString()
	queries := []string{}
	args := [][]interface{}{}

	emailAdmin := os.Getenv("EMAIL_ADMIN")
	cuitAdmin := os.Getenv("CUIT_ADMIN")
	nameAdmin := os.Getenv("NAME_ADMIN")
	phoneAdmin := os.Getenv("PHONE_ADMIN")
	addressAdmin := os.Getenv("ADDRESS_ADMIN")
	cityAdmin := os.Getenv("CITY_ADMIN")
	countryAdmin := os.Getenv("COUNTRY_ADMIN")
	zipCodeAdmin := os.Getenv("ZIPCODE_ADMIN")
	firstNameAdmin := os.Getenv("FIRSTNAME_ADMIN")
	lastNameAdmin := os.Getenv("LASTNAME_ADMIN")
	identifierAdmin := os.Getenv("IDENTIFIER_ADMIN")
	passwordAdmin := os.Getenv("PASSWORD_ADMIN")

	hashPassword, err := utils.HashPassword(passwordAdmin)

	if err != nil {
		log.Fatalf("Error al crear admin, hash_pass: %v", err)
	}

	query = `INSERT INTO users (id, first_name, last_name, email, identifier, phone, address, city, country, zip_code, password) values (?,?,?,?,?,?,?,?,?,?,?);`
	args = append(args, []interface{}{&userId, &firstNameAdmin, &lastNameAdmin, &emailAdmin, &identifierAdmin, &phoneAdmin, &addressAdmin, &cityAdmin, &countryAdmin, &zipCodeAdmin, &hashPassword})
	queries = append(queries, query)

	query = `INSERT INTO entities (id, email, cuit, name, phone, start_activities, address, city, country, zip_code) values (?,?,?,?,?,?,?,?,?,?);`
	args = append(args, []interface{}{&entityId, &emailAdmin, &cuitAdmin, &nameAdmin, &phoneAdmin, &parsedTime, &addressAdmin, &cityAdmin, &countryAdmin, &zipCodeAdmin})
	queries = append(queries, query)

	query = `INSERT INTO entity_user (id, user_id, entity_id, role_id) values (?,?,?,?)`
	queries = append(queries, query)
	args = append(args, []interface{}{&entityUserId, &userId, &entityId, &roleId})

	err = ExecuteGroupTransactions(queries, args)

	if err != nil {
		log.Fatalf("Error al crear admin: %v", err)
	}

	log.Println("Admin creado con exito")
	return nil
}
