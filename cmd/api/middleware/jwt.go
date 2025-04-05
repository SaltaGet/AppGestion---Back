package middleware

import (
	"os"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/contrib/jwt"
)

func JWTProtected(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
		TokenLookup: "header:X-Client-Token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
	})(c)
}

// func JWTUserProtected(c *fiber.Ctx) error {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_USER_KEY"))},
// 		TokenLookup: "header:X-User-Token",
// 		ErrorHandler: func(c *fiber.Ctx, err error) error {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": true,
// 				"msg":   err.Error(),
// 			})
// 		},
// 	})(c)
// }

// func JWTMultiProtected(requiredTokens ...string) fiber.Handler {
// 	fmt.Println("verificando toen")
// 	return func(c *fiber.Ctx) error {
// 		// Mapa de claves secretas y cabeceras
// 		tokenConfigs := map[string]string{
// 			"client": os.Getenv("SECRET_CLIENT_KEY"),
// 			"user":   os.Getenv("SECRET_USER_KEY"),
// 			"tenant":   os.Getenv("SECRET_TENANT_KEY"),
// 		}
// 		headerNames := map[string]string{
// 			"client": "X-Client-Token",
// 			"user":   "X-User-Token",
// 			"tenant":   "X-Tenant-Token",
// 		}

// 		// Verificar cada token requerido
// 		for _, tokenType := range requiredTokens {
// 			secretKey, exists := tokenConfigs[tokenType]
// 			if !exists {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"error": true,
// 					"msg":   "Invalid token type",
// 				})
// 			}

// 			tokenHeader := headerNames[tokenType]
// 			jwtMiddleware := jwtware.New(jwtware.Config{
// 				SigningKey: jwtware.SigningKey{Key: []byte(secretKey)},
// 				TokenLookup: "header:" + tokenHeader,
// 				ErrorHandler: func(c *fiber.Ctx, err error) error {
// 					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 						"error": true,
// 						"msg":   "Invalid or missing " + tokenType + " token",
// 					})
// 				},
// 			})

// 			if err := jwtMiddleware(c); err != nil {
// 				return err
// 			}
// 		}

// 		return c.Next()
// 	}
// }

// func JWTMultiProtected(requiredTokens ...string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		fmt.Println("Verificando token...")

// 		tokenConfigs := map[string]string{
// 			"client": os.Getenv("SECRET_CLIENT_KEY"),
// 			"user":   os.Getenv("SECRET_USER_KEY"),
// 			"tenant": os.Getenv("SECRET_TENANT_KEY"),
// 		}
// 		headerNames := map[string]string{
// 			"client": "X-Client-Token",
// 			"user":   "X-User-Token",
// 			"tenant": "X-Tenant-Token",
// 		}

// 		var userID string

// 		// Verificar cada token requerido
// 		for _, tokenType := range requiredTokens {
// 			secretKey, exists := tokenConfigs[tokenType]
// 			if !exists || secretKey == "" {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"error": true,
// 					"msg":   "Invalid or missing token type: " + tokenType,
// 				})
// 			}

// 			tokenHeader := headerNames[tokenType]
// 			tokenString := c.Get(tokenHeader)

// 			if tokenString == "" {
// 				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 					"error": true,
// 					"msg":   "Missing " + tokenType + " token",
// 				})
// 			}

// 			// Parsear el token
// 			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 				return []byte(secretKey), nil
// 			})

// 			if err != nil || !token.Valid {
// 				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 					"error": true,
// 					"msg":   "Invalid " + tokenType + " token",
// 				})
// 			}

// 			// Extraer userID del token
// 			if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 				if id, exists := claims["user_id"].(string); exists {
// 					userID = id
// 				}
// 			}
// 		}

// 		// Verificar el usuario en la base de datos
// 		var user models.User
// 		result := database.DB.Where("id = ?", userID).First(&user)
// 		if result.Error != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": true,
// 				"msg":   "User not found",
// 			})
// 		}

// 		// Almacenar el usuario completo en el contexto
// 		c.Locals("user", user)

// 		return c.Next()
// 	}
// }

// func JWTMultiProtected(requiredTokens ...string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		fmt.Println("Verificando token...")

// 		// Mapa de claves secretas y cabeceras
// 		tokenConfigs := map[string]string{
// 			"client": os.Getenv("SECRET_CLIENT_KEY"),
// 			"user":   os.Getenv("SECRET_USER_KEY"),
// 			"tenant": os.Getenv("SECRET_TENANT_KEY"),
// 		}
// 		headerNames := map[string]string{
// 			"client": "X-Client-Token",
// 			"user":   "X-User-Token",
// 			"tenant": "X-Tenant-Token",
// 		}

// 		// Verificar cada token requerido
// 		for _, tokenType := range requiredTokens {
// 			secretKey, exists := tokenConfigs[tokenType]
// 			if !exists || secretKey == "" {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"error": true,
// 					"msg":   "Invalid or missing token type: " + tokenType,
// 				})
// 			}

// 			tokenHeader := headerNames[tokenType]

// 			// Crear middleware de JWT para este token
// 			jwtMiddleware := jwtware.New(jwtware.Config{
// 				SigningKey: jwtware.SigningKey{Key: []byte(secretKey)},
// 				TokenLookup: "header:" + tokenHeader,
// 				ErrorHandler: func(c *fiber.Ctx, err error) error {
// 					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 						"error": true,
// 						"msg":   "Invalid or missing " + tokenType + " token",
// 					})
// 				},
// 			})

// 			// Ejecutar el middleware de JWT
// 			handler := jwtMiddleware(c)
// 			if handler != nil {
// 				return handler // Si hay error, devuelve el error
// 			}
// 		}

// 		// Si pasa todas las validaciones, continuar
// 		return c.Next()
// 	}
// }