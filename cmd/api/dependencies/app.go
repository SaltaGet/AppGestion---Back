package dependencies

import (
	entCtrl "api-stock/cmd/api/controllers/entity"
	entRep "api-stock/pkg/repository/entity"
	entServ "api-stock/pkg/services/entity"
	userCtrl "api-stock/cmd/api/controllers/user"
	userRep "api-stock/pkg/repository/user"
	userServ "api-stock/pkg/services/user"
	authCtrl "api-stock/cmd/api/controllers/auth"
	authRep "api-stock/pkg/repository/auth"
	authServ "api-stock/pkg/services/auth"
	"database/sql"
)

type Application struct {
	EntityController *entCtrl.Controller
	UserController *userCtrl.Controller
	AuthController *authCtrl.Controller
}

func NewApplication(db *sql.DB) *Application {
	entityRepo := &entRep.Repository{DB: db,}
	entityServ := &entServ.Service{EntityRepository: entityRepo}

	userRepo := &userRep.Repository{DB: db,}
	userServ := &userServ.Service{UserRepository: userRepo}

	authRepo := &authRep.Repository{DB: db,}
	authServ := &authServ.Service{AuthRepository: authRepo, UserRepository: userRepo}
	
	return &Application{
		EntityController: &entCtrl.Controller{EntityService: entityServ},
		UserController: &userCtrl.Controller{UserService: userServ},
		AuthController: &authCtrl.Controller{AuthService: authServ},
	}
}

func (app *Application) SetDBRepository(db *sql.DB) {
	entityRepo := &entRep.Repository{DB: db,}
	entityServ := &entServ.Service{EntityRepository: entityRepo}

	userRepo := &userRep.Repository{DB: db,}
	userServ := &userServ.Service{UserRepository: userRepo}

	authRepo := &authRep.Repository{DB: db,}
	authServ := &authServ.Service{AuthRepository: authRepo, UserRepository: userRepo}

	app.UserController.UserService = userServ
	app.EntityController.EntityService = entityServ
	app.AuthController.AuthService = authServ
}

