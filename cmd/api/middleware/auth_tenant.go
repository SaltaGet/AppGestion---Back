package middleware

import (
	"appGestion/pkg/dependencies"
	"appGestion/pkg/key"
	"appGestion/pkg/models"
	"appGestion/pkg/models/user"
	"appGestion/pkg/repository/database"
	"appGestion/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAauthTenant(c *fiber.Ctx) error {
	token := c.Get("X-Tenant")
	if token == "" {
		return c.Status(401).JSON(models.Response{
			Status:  false,
			Message: "Token no encontrado",
		})
	}

	ctx := c.UserContext()
	deps := ctx.Value(key.AppKey).(*dependencies.Application)
	// deps.SetDBRepository(tenantDB)

	claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(401).JSON(models.Response{
				Status:  false,
				Message: "Token inválido",
			})
		}
	
	establishmentId := claims.(jwt.MapClaims)["establishment_id"].(string)
	user := c.Locals("user").(*user.User)

	uri, err := deps.AuthController.AuthService.GetConnectionTenant(establishmentId, user.Id)

	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	tenantDB, err := database.GetTenantDB(uri)

	if err != nil {
		return c.Status(500).JSON(models.Response{
			Status:  false,
			Message: "Error de conexión al tenant",
		})
	}

	deps.SetDBRepository(tenantDB)
	return c.Next()
}