package entity

import (
	"api-stock/pkg/ports"	
	"database/sql"
)

type Service struct {
	EntityRepository ports.EntityRepository	
}

func (s *Service) SetUserRepositoryDB(db *sql.DB) {
	if repo, ok := s.EntityRepository.(interface{ SetDB(*sql.DB) }); ok {
			repo.SetDB(db)
	}
}