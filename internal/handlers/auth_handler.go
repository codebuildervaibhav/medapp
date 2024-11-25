// handlers/auth_handler.go
package handlers

import (
	"net/http"

	"github.com/codebuildervaibhav/medapp/config"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{cfg: cfg}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Login logic here
}
