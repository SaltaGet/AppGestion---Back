package middleware

import (
	"appGestion/pkg/dependencies"
	"appGestion/pkg/key"
	"appGestion/pkg/models"
	"appGestion/pkg/repository/database"
	"appGestion/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)


func JWTAauth(isAdmin ...bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(401).JSON(models.Response{
				Status:  false,
				Message: "Token no encontrado",
			})
		}
	
		tenantDB, err := database.GetTenantDB("default")
		if err != nil {
			return c.Status(500).JSON(models.Response{
				Status:  false,
				Message: "Error de conexión al tenant",
			})
		}
	
		ctx := c.UserContext()
		deps := ctx.Value(key.AppKey).(*dependencies.Application)
		deps.SetDBRepository(tenantDB)
	
		claims, err := utils.VerifyToken(token)
			if err != nil {
				return c.Status(401).JSON(models.Response{
					Status:  false,
					Message: "Token inválido",
				})
			}
		
		userId := claims.(jwt.MapClaims)["id"].(string)
	
		user, err := deps.AuthController.AuthService.GetCurrentUser(userId)
	
		if err != nil {
			return c.Status(500).JSON(models.Response{
				Status:  false,
				Message: "Error de conexión al tenant",
			})
		}
	
		if user == nil {
			return c.Status(500).JSON(models.Response{
				Status:  false,
				Message: "Error de conexión al tenant",
			})
		}
	
		requireAdmin := false
		if len(isAdmin) > 0 {
			requireAdmin = isAdmin[0]
		}
	
		if requireAdmin && !user.IsAdmin {
			return c.Status(403).JSON(models.Response{
				Status:  false,
				Message: "Acceso solo para administradores",
			})
		}
	
		c.Locals("user", user)
		return c.Next()
	}
}

// func RequireAuth(isAdmin ...bool) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return JWTAauth(c, isAdmin...)
// 	}
// }

// func JWTAauth(c *fiber.Ctx, isAdmin ...bool) error {
// 	token := c.Get("Authorization")
// 	if token == "" {
// 		return c.Status(401).JSON(models.Response{
// 			Status:  false,
// 			Message: "Token no encontrado",
// 		})
// 	}

// 	tenantDB, err := database.GetTenantDB("default")
// 	if err != nil {
// 		return c.Status(500).JSON(models.Response{
// 			Status:  false,
// 			Message: "Error de conexión al tenant",
// 		})
// 	}

// 	ctx := c.UserContext()
// 	deps := ctx.Value(key.AppKey).(*dependencies.Application)
// 	deps.SetDBRepository(tenantDB)

// 	claims, err := utils.VerifyToken(token)
// 		if err != nil {
// 			return c.Status(401).JSON(models.Response{
// 				Status:  false,
// 				Message: "Token inválido",
// 			})
// 		}
	
// 	userId := claims.(jwt.MapClaims)["id"].(string)

// 	user, err := deps.AuthController.AuthService.GetCurrentUser(userId)

// 	if err != nil {
// 		return c.Status(500).JSON(models.Response{
// 			Status:  false,
// 			Message: "Error de conexión al tenant",
// 		})
// 	}

// 	if user == nil {
// 		return c.Status(500).JSON(models.Response{
// 			Status:  false,
// 			Message: "Error de conexión al tenant",
// 		})
// 	}

// 	requireAdmin := false
// 	if len(isAdmin) > 0 {
// 		requireAdmin = isAdmin[0]
// 	}

// 	if requireAdmin && !user.IsAdmin {
// 		return c.Status(403).JSON(models.Response{
// 			Status:  false,
// 			Message: "Acceso solo para administradores",
// 		})
// 	}

// 	c.Locals("user", user)
// 	return c.Next()
// }

// func AuthTenantMiddleware(deps *dependencies.Application) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		if isPublic, ok := c.Locals("is_public").(bool); ok && isPublic {
// 			tenantDB, err := database.GetTenantDB("default")
// 			if err != nil {
// 				return c.Status(500).JSON(models.Response{
// 					Status:  false,
// 					Message: "Error de conexión al tenant",
// 				})
// 			}
// 			deps.SetDBRepository(tenantDB)
// 			return c.Next()
// 		}
// 		// 1. Autenticación (usa DB principal)
// 		token := c.Get("Authorization")
// 		tenant := c.Get("X-Tenant")
// 		if token == "" {
// 			ConnectionDB, err := database.GetTenantDB("default")
// 			if err != nil {
// 				return c.Status(500).JSON(models.Response{
// 					Status:  false,
// 					Message: "Error de conexión al tenant",
// 				})
// 			}

// 			deps.SetDBRepository(ConnectionDB)
// 			return c.Next()
// 		}

// 		claims, err := utils.VerifyToken(token)
// 		if err != nil {
// 			return c.Status(401).JSON(models.Response{
// 				Status:  false,
// 				Message: "Token inválido",
// 			})
// 		}

// 		userId := claims.(jwt.MapClaims)["user_id"].(string)
// 		user, err := deps.AuthController.AuthService.GetCurrentUser(userId)

// 		if err != nil {
// 			return c.Status(500).JSON(models.Response{
// 				Status:  false,
// 				Message: "Error al obtener usuario",
// 			})
// 		}

// 		if user == nil {
// 			ConnectionDB, err := database.GetTenantDB("default")
// 			if err != nil {
// 				return c.Status(500).JSON(models.Response{
// 					Status:  false,
// 					Message: "Error de conexión a la db",
// 				})
// 			}
// 			deps.SetDBRepository(ConnectionDB)
// 			return c.Status(404).JSON(models.Response{
// 				Status:  false,
// 				Message: "Usuario no encontrado loguearse nuevamente",
// 			})
// 		}

// 		c.Locals("user", user)

// 		if tenant == ""  && token != "" {
			
			
// 			ConnectionDB, err := database.GetTenantDB("default")
// 			if err != nil {
// 				return c.Status(500).JSON(models.Response{
// 					Status:  false,
// 					Message: "Error de conexión a la db",
// 				})
// 			}

// 			deps.SetDBRepository(ConnectionDB)
// 			return c.Next()
// 		}

// 		// Valida token y obtiene tenant del usuario
// 		claims, err := utils.VerifyToken(token)
// 		if err != nil {
// 			return c.Status(401).JSON(models.Response{
// 				Status:  false,
// 				Message: "Token inválido",
// 			})
// 		}

// 		TenantID := claims.(jwt.MapClaims)["tenant_id"].(string)
// 		// UserID := claims.(jwt.MapClaims)["user_id"].(string)

// 		// 2. Obtiene conexión del tenant
// 		tenantDB, err := database.GetTenantDB(TenantID)
// 		if err != nil {
// 			return c.Status(500).JSON(models.Response{
// 				Status:  false,
// 				Message: "Error de conexión al tenant",
// 			})
// 		}

// 		deps.SetDBRepository(tenantDB)

// 		return c.Next()
// 	}
// }