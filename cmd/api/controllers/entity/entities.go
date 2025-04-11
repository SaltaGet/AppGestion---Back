package entity

import (
	ent "appGestion/pkg/models/entity"
	resp "appGestion/pkg/models"
	"github.com/gofiber/fiber/v2"
)

//  Entity Entity create
//	@Summary		Entity create
//	@Description	Entity create
//	@Tags			Entity
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
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
		Body:    id,
		Message: "Entidad creada con éxito",
	})
}

//  Entity Entity get
//	@Summary		Entity all get
//	@Description	Entity all create
//	@Tags			Entity
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	resp.Response{body=[]ent.Entity}
//	@Failure		401	{object}	resp.Response{body=nil}
//	@Failure		500	{object}	resp.Response{body=nil}
//	@Router			/entities/get_all [get]
func (ctrl Controller) GetAll(c *fiber.Ctx) error {
	entities, err := ctrl.EntityService.GetAll()

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
		Body:    entities,
		Message: "Entidades obtenidas con éxito",
	})
}
