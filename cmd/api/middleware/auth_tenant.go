package middleware

import (
	"api-stock/cmd/api/dependencies"
	"api-stock/pkg/models"
	"api-stock/pkg/repository/database"
	"api-stock/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthTenantMiddleware(deps *dependencies.Application) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// if isPublic, ok := c.Locals("is_public").(bool); ok && isPublic {
		// 	tenantDB, err := database.GetTenantDB("default")
		// 	if err != nil {
		// 		return c.Status(500).JSON(models.Response{
		// 			Status:  false,
		// 			Message: "Error de conexión al tenant",
		// 		})
		// 	}
		// 	deps.SetDBRepository(tenantDB)
		// 	return c.Next()
		// }
		// 1. Autenticación (usa DB principal)
		token := c.Get("Authorization")
		if token == "" {
			tenantDB, err := database.GetTenantDB("default")
			if err != nil {
				return c.Status(500).JSON(models.Response{
					Status:  false,
					Message: "Error de conexión al tenant",
				})
			}

			deps.SetDBRepository(tenantDB)
			return c.Next()
		}

		// Valida token y obtiene tenant del usuario
		claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(401).JSON(models.Response{
				Status:  false,
				Message: "Token inválido",
			})
		}

		TenantID := claims.(jwt.MapClaims)["tenant_id"].(string)
		// UserID := claims.(jwt.MapClaims)["user_id"].(string)

		// 2. Obtiene conexión del tenant
		tenantDB, err := database.GetTenantDB(TenantID)
		if err != nil {
			return c.Status(500).JSON(models.Response{
				Status:  false,
				Message: "Error de conexión al tenant",
			})
		}

		deps.SetDBRepository(tenantDB)

		return c.Next()
	}
}