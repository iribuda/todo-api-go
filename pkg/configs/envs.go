package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// globale Variable, die die Konfigurationseinstellungen lädt, wenn Package geladen wird
var Envs = initConfig()

// Typ, der die verschiedene Konfigurationseinstellungen enthält
type Configuration struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

// private Funktion für Initialisierung
func initConfig() Configuration {
	godotenv.Load()

	return Configuration{
		PublicHost: getEnv("PUBLIC_HOST"),
		Port:       getEnv("PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:     getEnv("DB_NAME"),
	}
}

// private Funktion für Abrufen von Einstellungen
func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}	else{
		log.Fatalf("env for %q was not found", key)
		return ""
	}
}
