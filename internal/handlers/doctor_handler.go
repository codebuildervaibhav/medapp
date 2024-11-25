package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/codebuildervaibhav/medapp/config"
	"github.com/codebuildervaibhav/medapp/internal/models"
	"github.com/codebuildervaibhav/medapp/internal/services"
	"github.com/codebuildervaibhav/medapp/pkg/utils"
	"github.com/gorilla/mux"
)

type DoctorHandler struct {
	cfg            *config.Config
	patientService services.PatientService
}

func NewDoctorHandler(cfg *config.Config, patientService services.PatientService) (*DoctorHandler, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	return &DoctorHandler{cfg: cfg, patientService: patientService}, nil
}

func (dh *DoctorHandler) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := dh.patientService.GetAllPatients()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error fetching patients")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, patients)
}

func (dh *DoctorHandler) GetPatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid patient ID")
		return
	}

	patient, err := dh.patientService.GetPatientByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Patient not found")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, patient)
}

func (dh *DoctorHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid patient ID")
		return
	}

	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}
	patient.ID = id

	err = dh.patientService.UpdatePatient(&patient)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error updating patient")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Patient updated successfully"})
}

func (dh *DoctorHandler) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid patient ID")
		return
	}

	patient, err := dh.patientService.GetPatientByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Patient not found")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, patient)
}

// delete patient
func (dh *DoctorHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid patient ID")
		return
	}

	err = dh.patientService.DeletePatient(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error deleting patient")
		return
	}

}
