package user

import (
	"appGestion/pkg/models/user"
	resp "appGestion/pkg/models"

	"github.com/gofiber/fiber/v2"
)

//  User User create
//	@Summary		User create
//	@Description	User create
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			user	body		user.UserCreate	true	"User data"
//	@Success		200		{object}	resp.Response
//	@Failure		400		{object}	resp.Response
//	@Failure		401		{object}	resp.Response
//	@Failure		404		{object}	resp.Response
//	@Failure		422		{object}	resp.Response
//	@Failure		500		{object}	resp.Response
//	@Router			/users/create [post]
func (ctrl *Controller) Create(c *fiber.Ctx) error {
	var user user.UserCreate

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al leer el JSON de la petición",
		})
	}

	if err := user.Validate(); err != nil {
		return c.Status(422).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := ctrl.UserService.Create(&user)

	if err != nil {
		if errResp, ok := err.(*resp.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(resp.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	return c.Status(200).JSON(resp.Response{
		Status:  true,
		Body:    map[string]string{"user_id":id},
		Message: "Usuario creado con éxito",
	})
}