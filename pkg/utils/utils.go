package utils

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
)

// für Validierung
var Validate = validator.New()

// Hilf-Funktion für Erstellen des JSON
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Hilf-Funktion für Fehlermeldung
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// Hilf-Funktion für Abrufen aus JSON
func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

// Hilf-Funktion für das Abrufen des Tokens
func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	fmt.Print(tokenAuth)
	fmt.Print(tokenQuery)
	
	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}