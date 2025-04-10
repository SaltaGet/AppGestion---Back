package establishment

import (
	"github.com/gofiber/fiber/v2"
	"appGestion/pkg/models/establishment"
	resp "appGestion/pkg/models"
)

//  Establishment Establishment create
//	@Summary		Establishment create
//	@Description	Establishment create
//	@Tags			Establishment
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			establishment	body		establishment.EstablishmentCreate	true	"Establishment data"
//	@Success		200				{object}	resp.Response
//	@Failure		400				{object}	resp.Response
//	@Failure		401				{object}	resp.Response
//	@Failure		404				{object}	resp.Response
//	@Failure		422				{object}	resp.Response
//	@Failure		500				{object}	resp.Response
//	@Router			/establishments/create [post]
func (ctrl *Controller) Create(c *fiber.Ctx) error {
	var establishment establishment.EstablishmentCreate
	
	if err := c.BodyParser(&establishment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al leer el JSON de la petición",
		})
	}

	if err := establishment.Validate(); err != nil {
		return c.Status(422).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := ctrl.EstablishmentService.Create(&establishment)

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
		Body:    id,
		Message: "Token obtenido con éxito",
	})
}