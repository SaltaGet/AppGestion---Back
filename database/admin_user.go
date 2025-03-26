package database

import(
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"api-stock/utils"
	"time"
	"log"
	"os"
)

func CreateAdmin() error {
	var err error
	
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	db := GetDB()

	new_id := uuid.NewString()

	query := `
		INSERT INTO clients (id, email, cuit, name, password, cellphone, role, is_active, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?,?,?)
	`

	pass_hash, err := utils.HashPassword(os.Getenv("PASSWORD_ADMIN"))

	if err != nil {
		log.Fatalf("Error al crear admin: %v", err)
	}

	_, err = db.Exec(
		query, 
		new_id, 
		os.Getenv("EMAIL_ADMIN"), 
		os.Getenv("CUIT_ADMIN"), 
		os.Getenv("NAME_ADMIN"), 
		pass_hash,
		os.Getenv("CELLPHONE_ADMIN"),
		os.Getenv("ROLE_ADMIN"),
		true,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Fatalf("Error al crear admin: %v", err)
	}

	return nil
}