package main

import (
	mdw "api-stock/cmd/api/middleware"
	"api-stock/cmd/api/routes"
	"api-stock/cmd/api/dependencies"
	"api-stock/pkg/repository/database"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	db, err := database.ConectDB(os.Getenv("URI_DB"))

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer database.CloseDB(db)

	appDependencies := dependencies.NewApplication(db)

	app := fiber.New()

	app.Use(mdw.LogginMiddleware)
	// app.Use(mdw.JWTProtected)

	routes.SetupRoutes(app, appDependencies)

	log.Fatal(app.Listen(":3000"))
}

