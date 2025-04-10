// cmd/api/middleware/db_monitor.go
package middleware

import (
	"database/sql"
	"log"
	"time"
)

func StartDBStatsLogger(db *sql.DB, interval time.Duration) {
	go func() {
		for {
			stats := db.Stats()
			log.Printf("DB Stats: MaxOpenConnections=%d, OpenConnections=%d, InUse=%d, Idle=%d",
				stats.MaxOpenConnections, stats.OpenConnections, stats.InUse, stats.Idle)
			time.Sleep(interval)
		}
	}()
}