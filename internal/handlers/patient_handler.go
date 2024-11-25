package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/codebuildervaibhav/medapp/config"
	"github.com/codebuildervaibhav/medapp/internal/models"
	"github.com/codebuildervaibhav/medapp/internal/repositories"
	"github.com/codebuildervaibhav/medapp/internal/services"
	"github.com/gorilla/mux"
)

type PatientHandler struct {
	repo repositories.PatientRepository
}

// NewPatientHandler initializes a new PatientHandler
func NewPatientHandler(repo repositories.PatientRepository) *PatientHandler {
	return &PatientHandler{repo: repo}
}

// GetAllPatients retrieves all patients
func (ph *PatientHandler) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := ph.repo.GetAllPatients()
	if err != nil {
		http.Error(w, "Failed to retrieve patients: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(patients); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetPatientByID retrieves a patient by ID
func (ph *PatientHandler) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := ph.repo.GetPatientByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if patient == nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

type ReceptionistHandler struct {
	cfg            *config.Config
	patientService services.PatientService
}

func NewReceptionistHandler(cfg *config.Config, patientService services.PatientService) (*ReceptionistHandler, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	return &ReceptionistHandler{
		cfg:            cfg,
		patientService: patientService,
	}, nil
}

// RegisterPatient registers a new patient
func (rh *ReceptionistHandler) RegisterPatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := rh.patientService.AddPatient(&patient); err != nil {
		http.Error(w, "Error registering patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

// AddPatient adds a new patient to the system
func (rh *ReceptionistHandler) AddPatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := rh.patientService.AddPatient(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

// GetAllPatients retrieves all patients
func (rh *ReceptionistHandler) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := rh.patientService.GetAllPatients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patients)
}

// GetPatient retrieves a patient by ID
func (rh *ReceptionistHandler) GetPatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := rh.patientService.GetPatientByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patient)
}

// DeletePatient deletes a patient by ID
func (rh *ReceptionistHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	err = rh.patientService.DeletePatient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreatePatient is a helper method that delegates to AddPatient
func (rh *ReceptionistHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	rh.AddPatient(w, r)
}

// UpdatePatient updates a patient
func (rh *ReceptionistHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	patient.ID = id

	err = rh.patientService.UpdatePatient(&patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

