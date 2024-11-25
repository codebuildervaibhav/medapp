package repositories

import (
	"github.com/codebuildervaibhav/medapp/config"
	"github.com/codebuildervaibhav/medapp/db"
	"github.com/codebuildervaibhav/medapp/internal/models"
)

// UserRepository is the interface for User repository
type UserRepository interface {
	// GetUser returns a user by the given ID
	GetUserByID(id int) (*models.User, error)
	// AddUser adds a new user
	AddUser(user *models.User) error
	// DeleteUser deletes a user by the given ID
	DeleteUser(id int) error
	// UpdateUser updates a user
	UpdateUser(user *models.User) error
}

func AddPatient(cfg *config.Config, patient *models.Patient) error {
	conn, err := db.Connect(cfg.DatabaseDSN())
	if err != nil {
		return err
	}

	query := "INSERT INTO patients (name, age, address) VALUES ($1, $2, $3)"
	_, err = conn.Exec(query, patient.Name, patient.Age, patient.Address)
	return err
}
