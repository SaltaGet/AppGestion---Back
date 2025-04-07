// cmd/api/middleware/db_monitor.go
package middleware

import (
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DBStatsLogger(db *sql.DB, interval time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
			go func() {
					for {
							stats := db.Stats()
							log.Printf("DB Stats: MaxOpenConnections=%d, OpenConnections=%d, InUse=%d, Idle=%d",
									stats.MaxOpenConnections, stats.OpenConnections, stats.InUse, stats.Idle)
							time.Sleep(interval)
					}
			}()
			return c.Next()
	}
}