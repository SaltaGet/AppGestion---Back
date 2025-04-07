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
	"time"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	// Conexión principal (para auth)
	db, err := database.ConectDB(os.Getenv("URI_DB"))
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer func() {
		database.CloseDB(db)
		database.CloseAllTenantDBs()
	}()

	dbMonitorInterval := 5 * time.Minute
	appDependencies := dependencies.NewApplication(db)
	
	app := fiber.New()
	app.Use(mdw.LogginMiddleware)
	app.Use(mdw.AuthTenantMiddleware(appDependencies)) // <-- Nuevo middleware
	app.Use(mdw.DBStatsLogger(db, dbMonitorInterval))
	
	routes.SetupRoutes(app, appDependencies)
	
	log.Fatal(app.Listen(":3000"))
}



// func main() {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		log.Fatalf("Error cargando el archivo .env: %v", err)
// 	}

// 	db, err := database.ConectDB(os.Getenv("URI_DB"))

// 	if err != nil {
// 		log.Fatalf("Error al conectar con la base de datos: %v", err)
// 	}
// 	defer database.CloseDB(db)

// 	appDependencies := dependencies.NewApplication(db)

// 	app := fiber.New()

// 	app.Use(mdw.LogginMiddleware)

// 	routes.SetupRoutes(app, appDependencies)

// 	log.Fatal(app.Listen(":3000"))
// }

