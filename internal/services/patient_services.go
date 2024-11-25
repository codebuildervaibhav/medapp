package services

import (
	"errors"
	"log"

	"github.com/codebuildervaibhav/medapp/internal/models"
	"github.com/codebuildervaibhav/medapp/internal/repositories"
)

// PatientService defines methods for managing patient data
type PatientService struct {
	repo repositories.PatientRepository
}

// NewPatientService creates a new PatientService instance
func NewPatientService(repo repositories.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

// AddPatient adds a new patient to the database
func (ps *PatientService) AddPatient(patient *models.Patient) error {
	if patient == nil {
		return errors.New("patient cannot be nil")
	}

	if patient.Name == "" {
		return errors.New("patient name is required")
	}

	if err := ps.repo.AddPatient(patient); err != nil {
		log.Printf("Error adding patient: %v\n", err)
		return err
	}
	return nil
}

// GetAllPatients retrieves all patients from the database using service layer
func (ps *PatientService) GetAllPatients() ([]*models.Patient, error) {
	patients, err := ps.repo.GetAllPatients()
	if err != nil {
		log.Printf("Error fetching all patients: %v\n", err)
		return nil, err
	}
	return patients, nil
}

// GetPatientByID retrieves a single patient by their ID
func (ps *PatientService) GetPatientByID(id int) (*models.Patient, error) {
	if id <= 0 {
		return nil, errors.New("invalid patient ID")
	}

	patient, err := ps.repo.GetPatientByID(id)
	if err != nil {
		log.Printf("Error fetching patient by ID: %v\n", err)
		return nil, err
	}
	if patient == nil {
		log.Printf("Patient with ID %d not found\n", id)
		return nil, errors.New("patient not found")
	}
	return patient, nil
}

// DeletePatient removes a patient from the database
func (ps *PatientService) DeletePatient(id int) error {
	if id <= 0 {
		return errors.New("invalid patient ID")
	}

	if err := ps.repo.DeletePatient(id); err != nil {
		log.Printf("Error deleting patient with ID %d: %v\n", id, err)
		return err
	}
	return nil
}

// UpdatePatient updates a patient's information
func (ps *PatientService) UpdatePatient(patient *models.Patient) error {
	if patient == nil {
		return errors.New("patient cannot be nil")
	}

	if patient.ID <= 0 {
		return errors.New("invalid patient ID")
	}

	if err := ps.repo.UpdatePatient(patient); err != nil {
		log.Printf("Error updating patient with ID %d: %v\n", patient.ID, err)
		return err
	}
	return nil
}
