package routes

import (
	"appGestion/cmd/api/controllers/user"
	"appGestion/cmd/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, ctrl *user.Controller) {
	users := app.Group("/users")
	users.Post("/create", middleware.JWTAauth(true), ctrl.Create) 
}
