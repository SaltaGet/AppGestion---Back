package user

import (
	"api-stock/pkg/ports"	
	"database/sql"
)

type Service struct {
	UserRepository ports.UserRepository	
}

func (s *Service) SetUserRepositoryDB(db *sql.DB) {
	if repo, ok := s.UserRepository.(interface{ SetDB(*sql.DB) }); ok {
			repo.SetDB(db)
	}
}