package routes

import "github.com/gofiber/fiber/v2"

func DailyMovementsRoutes(app *fiber.App) {
	users := app.Group("/daily_movements")
	users.Get("/", getDailyMovements)       // GET /users
	users.Get("/:id", getDailyMovementsById) // GET /users/:id
	users.Post("/", createDailyMovements)    // POST /user}
}

func getDailyMovements(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getDailyMovementsById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createDailyMovements(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }