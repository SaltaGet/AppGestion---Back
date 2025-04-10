package auth

import (
	"appGestion/pkg/models/auth"
	resp "appGestion/pkg/models"

	"github.com/gofiber/fiber/v2"
)

//  Login Login user
//	@Summary		Login user
//	@Description	Login user required identifier and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		auth.AuthLogin	true	"Credentials"
//	@Success		200			{object}	resp.Response
//	@Failure		400			{object}	resp.Response
//	@Failure		401			{object}	resp.Response
//	@Failure		422			{object}	resp.Response
//	@Failure		404			{object}	resp.Response
//	@Failure		500			{object}	resp.Response
//	@Router			/auth/login [post]
func (ctrl Controller) Login(c *fiber.Ctx) error {
	var authLogin auth.AuthLogin

	if err := c.BodyParser(&authLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al leer el JSON de la petición",
		})
	}

	if err := authLogin.Validate(); err != nil {
		return c.Status(422).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	token, err := ctrl.AuthService.Login(&authLogin)

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
		Body:    token,
		Message: "Token obtenido con éxito",
	})
}
