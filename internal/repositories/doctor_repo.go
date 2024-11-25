package repositories

import (
	"database/sql"

	"github.com/codebuildervaibhav/medapp/internal/models"
)

// DoctorRepository defines the interface for doctor-related data operations
type DoctorRepository interface {
	GetAllDoctors() ([]*models.Doctor, error)
	GetDoctorByID(id int) (*models.Doctor, error)
	AddDoctor(doctor *models.Doctor) error
	UpdateDoctor(doctor *models.Doctor) error
	DeleteDoctor(id int) error
}

// doctorRepository is the concrete implementation of DoctorRepository
type doctorRepository struct {
	db *sql.DB
}

// NewDoctorRepository initializes a new doctor repository
func NewDoctorRepository(db *sql.DB) DoctorRepository {
	return &doctorRepository{db: db}
}

func (d *doctorRepository) GetAllDoctors() ([]*models.Doctor, error) {
	query := "SELECT id, name, specialization FROM doctors"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []*models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		if err := rows.Scan(&doctor.ID, &doctor.Name, &doctor.Specialization); err != nil {
			return nil, err
		}
		doctors = append(doctors, &doctor)
	}
	return doctors, nil
}

func (d *doctorRepository) GetDoctorByID(id int) (*models.Doctor, error) {
	query := "SELECT id, name, specialization FROM doctors WHERE id = $1"
	row := d.db.QueryRow(query, id)

	var doctor models.Doctor
	if err := row.Scan(&doctor.ID, &doctor.Name, &doctor.Specialization); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &doctor, nil
}

func (d *doctorRepository) AddDoctor(doctor *models.Doctor) error {
	query := "INSERT INTO doctors (name, specialization) VALUES ($1, $2)"
	_, err := d.db.Exec(query, doctor.Name, doctor.Specialization)
	return err
}

func (d *doctorRepository) UpdateDoctor(doctor *models.Doctor) error {
	query := "UPDATE doctors SET name = $1, specialization = $2 WHERE id = $3"
	_, err := d.db.Exec(query, doctor.Name, doctor.Specialization, doctor.ID)
	return err
}

func (d *doctorRepository) DeleteDoctor(id int) error {
	query := "DELETE FROM doctors WHERE id = $1"
	_, err := d.db.Exec(query, id)
	return err
}
