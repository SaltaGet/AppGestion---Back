package dependencies

import (
	entCtrl "appGestion/cmd/api/controllers/entity"
	entRep "appGestion/pkg/repository/entity"
	entServ "appGestion/pkg/services/entity"
	userCtrl "appGestion/cmd/api/controllers/user"
	userRep "appGestion/pkg/repository/user"
	userServ "appGestion/pkg/services/user"
	authCtrl "appGestion/cmd/api/controllers/auth"
	authRep "appGestion/pkg/repository/auth"
	authServ "appGestion/pkg/services/auth"
	estCtrl "appGestion/cmd/api/controllers/establishment"
	estRep "appGestion/pkg/repository/establishment"
	estServ "appGestion/pkg/services/establishment"
	"database/sql"
)

var App *Application

type Application struct {
	EntityController *entCtrl.Controller
	UserController *userCtrl.Controller
	AuthController *authCtrl.Controller
	EstablishmentController *estCtrl.Controller
}

func NewApplication(db *sql.DB) *Application {
	entityRepo := &entRep.Repository{DB: db,}
	entityServ := &entServ.Service{EntityRepository: entityRepo}

	userRepo := &userRep.Repository{DB: db,}
	userServ := &userServ.Service{UserRepository: userRepo}

	establishmentRepo := &estRep.Repository{DB: db,}
	establishmentServ := &estServ.Service{EstablishmentRepository: establishmentRepo, EntityRepository: entityRepo}
	
	authRepo := &authRep.Repository{DB: db,}
	authServ := &authServ.Service{AuthRepository: authRepo, UserRepository: userRepo, EstablishmentRepository: establishmentRepo}

	return &Application{
		EntityController: &entCtrl.Controller{EntityService: entityServ},
		UserController: &userCtrl.Controller{UserService: userServ},
		AuthController: &authCtrl.Controller{AuthService: authServ},
		EstablishmentController: &estCtrl.Controller{EstablishmentService: establishmentServ},
	}
}

func (app *Application) SetDBRepository(db *sql.DB) {
	entityRepo := &entRep.Repository{DB: db,}
	entityServ := &entServ.Service{EntityRepository: entityRepo}

	userRepo := &userRep.Repository{DB: db,}
	userServ := &userServ.Service{UserRepository: userRepo}

	authRepo := &authRep.Repository{DB: db,}
	authServ := &authServ.Service{AuthRepository: authRepo, UserRepository: userRepo}

	establishmentRepo := &estRep.Repository{DB: db,}
	establishmentServ := &estServ.Service{EstablishmentRepository: establishmentRepo, EntityRepository: entityRepo}

	app.UserController.UserService = userServ
	app.EntityController.EntityService = entityServ
	app.AuthController.AuthService = authServ
	app.EstablishmentController.EstablishmentService = establishmentServ
}

