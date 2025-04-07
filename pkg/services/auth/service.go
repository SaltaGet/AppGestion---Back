package auth

import (
	"api-stock/pkg/ports"	
	"database/sql"
)

type Service struct {
	AuthRepository ports.AuthRepository	
	UserRepository ports.UserRepository
}

func (s *Service) SetAuthRepositoryDB(db *sql.DB) {
	if repo, ok := s.UserRepository.(interface{ SetDB(*sql.DB) }); ok {
			repo.SetDB(db)
	}
}