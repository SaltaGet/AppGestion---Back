package dependencies

import (
	entCtrl "api-stock/cmd/api/controllers/entity"
	entRep "api-stock/pkg/repository/entity"
	entServ "api-stock/pkg/services/entity"
	"database/sql"
)

type Application struct {
	EntityController *entCtrl.Controller
}

func NewApplication(db *sql.DB) *Application {
	entityRepo := &entRep.Repository{DB: db,}

	entityServ := &entServ.Service{EntityRepository: entityRepo}
	
	return &Application{
		EntityController: &entCtrl.Controller{EntityService: entityServ},
	}
}

