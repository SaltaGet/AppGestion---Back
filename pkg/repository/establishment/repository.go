package establishment

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) SetDB(db *sql.DB) {
	r.DB = db
}