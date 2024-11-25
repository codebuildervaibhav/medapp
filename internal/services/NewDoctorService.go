package services

import "github.com/codebuildervaibhav/medapp/internal/repositories"

// DoctorService defines operations related to doctors.
type DoctorService struct {
	doctorRepo repositories.DoctorRepository
}

// NewDoctorService creates a new instance of DoctorService.
func NewDoctorService(doctorRepo repositories.DoctorRepository) *DoctorService {
	return &DoctorService{
		doctorRepo: doctorRepo,
	}
}

// Add methods to handle doctor-related business logic as needed.
