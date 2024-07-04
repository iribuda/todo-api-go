package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Funktion für Hashing des Passworts
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Vergleichen des Passworts, das Benutzer angegeben hat, mit dem, der in Datenbank gespeichert ist
func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}
