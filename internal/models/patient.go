package models

import "time"

type Patient struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
	DoctorID    int       `json:"doctor_id"`
	DoctorNotes string    `json:"doctor_notes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
