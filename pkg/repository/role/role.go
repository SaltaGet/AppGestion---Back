package role

import (
	"appGestion/pkg/models/role"
	"appGestion/pkg/repository/database"
	"appGestion/pkg/utils"
)

func (r *Repository) GetById(id string) (*role.Role, error) {
	query := `SELECT * FROM roles WHERE id = ? LIMIT 1`
	row, err := database.GetRows(r.DB, query, id)

	if err != nil {
		return nil, err
	}

	var role role.Role
	err = utils.MapRowToStruct(row, &role)

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *Repository) ExistById(id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE id = ?)`

	row := database.GetRow(r.DB, query, id)

	var exist bool
	if err := row.Scan(&exist); err != nil {
		return false, err
	}

	return exist, nil
}

func (r *Repository) GetAll() (*[]role.Role, error) {
	query := `SELECT * FROM roles`
	rows, err := database.GetRows(r.DB, query)

	if err != nil {
		return nil, err
	}

	var roles []role.Role
	err = utils.MapRowsToStruct(rows, &roles)

	if err != nil {
		return nil, err
	}

	return &roles, nil
}
