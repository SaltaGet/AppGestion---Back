package services

import(
	"api-stock/database"
	// "api-stock/models/client"
)

func GetCurrentClient(clientId string) (*cl.Client, error){
	db := database.GetDB()

	query := `SELECT * FROM clients WHERE id = ?`

	row := db.QueryRow(query, clientId,)

	var client cl.Client

	err := row.Scan(&client.Id, &client.Email, &client.CUIT, &client.Name,
		&client.Password, &client.Cellphone, &client.Role, &client.IsActive, &client.CreatedAt, &client.UpdatedAt)
	
	if err != nil {
		return nil, err
	}

	return &client, nil
}