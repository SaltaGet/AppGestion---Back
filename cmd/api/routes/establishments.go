package routes

import (
	"appGestion/cmd/api/controllers/establishment"
	"appGestion/cmd/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func EntablishmentsRoutes(app *fiber.App, ctrl *establishment.Controller) {
	establishments := app.Group("/establishments")
	establishments.Post("/create", middleware.JWTAauth(true), ctrl.Create)  
	establishments.Get("/get_all_admin", middleware.JWTAauth(true), ctrl.GetAllAdmin) 
}
