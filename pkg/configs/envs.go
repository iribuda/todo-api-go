package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

// globale Variable, die die Konfigurationseinstellungen lädt, wenn Package geladen wird
var Envs = initConfig()

// Typ, der die verschiedene Konfigurationseinstellungen enthält
type Configuration struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

// private Funktion für Initialisierung
func initConfig() Configuration {
	godotenv.Load()

	return Configuration{
		PublicHost:             getEnv("PUBLIC_HOST"),
		Port:                   getEnv("PORT"),
		DBUser:                 getEnv("DB_USER"),
		DBPassword:             getEnv("DB_PASSWORD"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:                 getEnv("DB_NAME"),
		JWTSecret:              getEnv("JWT_SECRET"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS"),
	}
}

// private Funktion für Abrufen von Einstellungen
func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		log.Fatalf("env for %q was not found", key)
		return ""
	}
}

func getEnvAsInt(key string) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Fatalf("env for %q was not found", key)
		}
		return i
	}
	return -1
}
