package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	UserRoutes(app)
	ProductRoutes(app)
	StocksRoutes(app)
	DiscontinuedRoutes(app)
	SalesRoutes(app)
	DailyMovementsRoutes(app)
}

// type User struct {
// 	ID int `json:"id"`
// 	Userame string `json:"username"`
// 	Email string `json:"email"`
// }

// func SetupRoutes(app *fiber.App) {
// 	app.Get("/", handlerInicio)

// 	// app.Get("/", func(c *fiber.Ctx) error {
// 	// 	return c.SendString("Hello, World!")
// 	// })

// 	app.Get("/about", handlerAbout)

// 	app.Get("/contact", func(c *fiber.Ctx) error {
// 		return c.SendString("estas en contacto my friend")
// 	})
	
// 	app.Get("/saludo/:nombre", handlerSaludo)
	
// 	app.Post("/user", func(c *fiber.Ctx) error {
// 		var user User
// 		if err := c.BodyParser(&user); err != nil {
// 			return err
// 		}
// 		return c.JSON(user)
// 	})
// }

// func handlerInicio(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World!")
// }

// func handlerAbout(c *fiber.Ctx) error {
// 	return c.SendString("estas en abaut")
// }

// func handlerSaludo(c *fiber.Ctx) error {
// 	nombre := c.Params("nombre")
// 	return c.SendString("Hello, " + nombre)
// }