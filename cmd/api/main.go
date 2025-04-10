package main

import (
	"appGestion/pkg/dependencies"
	mdw "appGestion/cmd/api/middleware"
	"appGestion/cmd/api/routes"
	"appGestion/pkg/repository/database"
	"log"
	"os"
	"time"

	_ "appGestion/cmd/api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

//	@title						APP GESTION
//	@version					1.0
//	@description				This is a api to app gestion
//	@termsOfService				http://swagger.io/terms/
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the JWT token. Example: "Bearer eyJhbGciOiJIUz..."

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

	
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Tenant",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
	}))
	// app.Use(mdw.RecoverMiddleware)
	// app.Use(mdw.RequestIDMiddleware)
	// app.Use(mdw.LoggerMiddleware)
	// app.Use(mdw.CompressMiddleware)
	// app.Use(mdw.SwaggerMiddleware(swagger.Config{
	// 	URL:     "/swagger/doc.json",
	// 	DocPath: "/swagger/doc",
	// }))
	// app.Use(mdw.SwaggerUIHandler(swagger.Config{
	// 	URL:     "/swagger/doc.json",
	// 	DocPath: "/swagger/doc",
	// }))
	// app.Use(mdw.SwaggerUIHandler(swagger.Config{
		// 	URL:     "/swagger/doc.json",
		// 	DocPath: "/swagger/doc",
		// }))
		mdw.StartDBStatsLogger(db, dbMonitorInterval)
		
		app.Use(mdw.LoggingMiddleware)
		app.Use(mdw.InjectApp(appDependencies))
		// app.Use(mdw.AuthTenantMiddleware(appDependencies)) // <-- Nuevo middleware
		
		routes.SetupRoutes(app, appDependencies)
		
		app.Get("/swagger/*", swagger.HandlerDefault)

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
