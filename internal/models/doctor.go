package models

import "time"

// Doctor represents a doctor
type Doctor struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Specialization string    `json:"specialization"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

