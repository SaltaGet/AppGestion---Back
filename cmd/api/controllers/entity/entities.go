package entity

import (
	ent "appGestion/pkg/models/entity"
	resp "appGestion/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//  Entity Entity create
//	@Summary		Entity create
//	@Description	Entity create
//	@Tags			Entity
//	@Accept			json
//	@Produce		json
//	@Param			entity	body		ent.EntityCreate	true	"Entity data"
//	@Success		200		{object}	resp.Response
//	@Failure		400		{object}	resp.Response
//	@Failure		401		{object}	resp.Response
//	@Failure		404		{object}	resp.Response
//	@Failure		422		{object}	resp.Response
//	@Failure		500		{object}	resp.Response
//	@Router			/entities/create [post]
func (ctrl Controller) CreateEntity(c *fiber.Ctx) error {
	var entity ent.EntityCreate

	err := c.BodyParser(&entity)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al crear entidad",
		})
	}

	if err := entity.Validate(); err != nil {
		return c.Status(422).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := ctrl.EntityService.Create(&entity)

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

	return c.Status(fiber.StatusOK).JSON(resp.Response{
		Status:  true,
		Body:    map[string]string{"entity_id": id},
		Message: "Entidad creada con éxito",
	})
}

// func ClientLogin(c *fiber.Ctx) error {
// 	var clienLogin cl.ClientLogin

// 	err := c.BodyParser(&clienLogin)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(m.Response{
// 			Status:  false,
// 			Body:    nil,
// 			Message: "Error al intentar loguearse",
// 		})
// 	}

// 	if err := clienLogin.Validate(); err != nil {
// 		return c.Status(422).JSON(m.Response{
// 			Status:  false,
// 			Body:    nil,
// 			Message: err.Error(),
// 		})
// 	}

// 	status, message, err := s.LoginClient(&clienLogin)

// 	if status > 299 {
// 		return c.Status(status).JSON(m.Response{
// 			Status:  false,
// 			Body:    err,
// 			Message: message,
// 		})
// 	}

// 	return c.Status(status).JSON(m.Response{
// 		Status:  true,
// 		Body:    map[string]string{"token": message},
// 		Message: "Token generado con exito",
// 	})
// }