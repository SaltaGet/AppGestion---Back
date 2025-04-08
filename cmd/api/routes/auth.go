package routes

import (
	"github.com/gofiber/fiber/v2"
	authCtrl "appGestion/cmd/api/controllers/auth"
	// mdw "api-stock/cmd/api/middleware"
)

func AuthRoutes(app *fiber.App, controller *authCtrl.Controller) {
	entities := app.Group("/auth")
	entities.Post("/login", controller.Login) 
}

