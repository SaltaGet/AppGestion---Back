package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", getUsers)       // GET /users
	users.Get("/:id", getUserByID) // GET /users/:id
	users.Post("/", createUser)    // POST /users
}

func getUsers(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getUserByID(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createUser(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }