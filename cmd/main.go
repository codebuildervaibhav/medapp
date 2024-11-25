package main

import (
	"log"
	"net/http"
	"time"

	"github.com/codebuildervaibhav/medapp/config"
	"github.com/codebuildervaibhav/medapp/db"
	"github.com/codebuildervaibhav/medapp/internal/handlers"
	"github.com/codebuildervaibhav/medapp/internal/repositories"
	"github.com/codebuildervaibhav/medapp/internal/services"
	"github.com/codebuildervaibhav/medapp/pkg/middleware"
	"github.com/codebuildervaibhav/medapp/pkg/utils"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Run database migrations
	utils.RunMigrations(cfg.DatabaseDSN(), "./migrations")

	// Get DB connection
	dbConn, err := db.Connect(cfg.DatabaseDSN())
	if err != nil {
		log.Fatalf("Failed to get DB connection: %v", err)
	}

	// Initialize repositories
	patientRepo := repositories.NewPatientRepository(dbConn)
	//doctorRepo := repositories.NewDoctorRepository(dbConn)

	// Initialize services
	patientService := services.NewPatientService(patientRepo)
	//doctorService := services.NewDoctorService(doctorRepo)

	// Initialize handlers
	receptionistHandler, err := handlers.NewReceptionistHandler(cfg, *(*services.PatientService)(patientService))
	if err != nil {
		log.Fatalf("Failed to initialize receptionist handler: %v", err)
	}
	doctorHandler, err := handlers.NewDoctorHandler(cfg, *(*services.PatientService)(patientService))
	//doctorHandler, err := handlers.NewDoctorHandler(cfg, *(*services.PatientService)(doctorService))
	if err != nil {
		log.Fatalf("Failed to initialize doctor handler: %v", err)
	}

	// Initialize router
	r := mux.NewRouter()

	// Middleware setup
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Authentication routes
	authHandler := handlers.NewAuthHandler(cfg)
	api.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)

	// Receptionist routes
	api.HandleFunc("/receptionist/patient", receptionistHandler.CreatePatient).Methods(http.MethodPost)
	api.HandleFunc("/receptionist/patients", receptionistHandler.GetAllPatients).Methods(http.MethodGet)
	api.HandleFunc("/receptionist/patient/{id}", receptionistHandler.GetPatient).Methods(http.MethodGet)
	api.HandleFunc("/receptionist/patient/{id}", receptionistHandler.UpdatePatient).Methods(http.MethodPut)
	api.HandleFunc("/receptionist/patient/{id}", receptionistHandler.DeletePatient).Methods(http.MethodDelete)
	api.HandleFunc("/receptionist/patient/register", receptionistHandler.RegisterPatient).Methods(http.MethodPost)

	// Doctor routes
	api.HandleFunc("/doctor/patients", doctorHandler.GetAllPatients).Methods(http.MethodGet)
	api.HandleFunc("/doctor/patients/{id}", doctorHandler.GetPatient).Methods(http.MethodGet)
	api.HandleFunc("/doctor/patients/{id}", doctorHandler.UpdatePatient).Methods(http.MethodPut)
	api.HandleFunc("/doctor/patients/{id}", doctorHandler.DeletePatient).Methods(http.MethodDelete)

	// Start the server
	server := &http.Server{

		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server on :8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
