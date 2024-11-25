package repositories

import (
	"database/sql"

	"github.com/codebuildervaibhav/medapp/internal/models"
)

// PatientRepository is the interface for Patient repository
type PatientRepository interface {
	AddPatient(patient *models.Patient) error
	GetPatientByID(id int) (*models.Patient, error)
	GetAllPatients() ([]*models.Patient, error)
	UpdatePatient(patient *models.Patient) error
	DeletePatient(id int) error
}

// patientRepository is the concrete implementation of PatientRepository
type patientRepository struct {
	db *sql.DB
}

// NewPatientRepository initializes a new patient repository
func NewPatientRepository(db *sql.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (p *patientRepository) AddPatient(patient *models.Patient) error {
	query := "INSERT INTO patients (name, age, address) VALUES ($1, $2, $3)"
	_, err := p.db.Exec(query, patient.Name, patient.Age, patient.Address)
	return err
}

func (p *patientRepository) GetPatientByID(id int) (*models.Patient, error) {
	query := "SELECT id, name, age, address FROM patients WHERE id = $1"
	row := p.db.QueryRow(query, id)

	var patient models.Patient
	if err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.Address); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

func (p *patientRepository) GetAllPatients() ([]*models.Patient, error) {
	query := "SELECT id, name, age, address FROM patients"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []*models.Patient
	for rows.Next() {
		var patient models.Patient
		if err := rows.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.Address); err != nil {
			return nil, err
		}
		patients = append(patients, &patient)
	}
	return patients, nil
}

func (p *patientRepository) UpdatePatient(patient *models.Patient) error {
	query := "UPDATE patients SET name = $1, age = $2, address = $3 WHERE id = $4"
	_, err := p.db.Exec(query, patient.Name, patient.Age, patient.Address, patient.ID)
	return err
}

func (p *patientRepository) DeletePatient(id int) error {
	query := "DELETE FROM patients WHERE id = $1"
	_, err := p.db.Exec(query, id)
	return err
}
