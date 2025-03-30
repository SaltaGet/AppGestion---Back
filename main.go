package main

import (
	"log"
	"api-stock/routes"
	"api-stock/database"
	mdw "api-stock/middleware"
	"github.com/gofiber/fiber/v2"
		"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	if err := database.InitDB(); err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer database.CloseDB()

	app := fiber.New()

	app.Use(mdw.LogginMiddleware)
	// app.Use(mdw.JWTProtected)

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

