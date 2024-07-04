package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"guthub.com/iribuda/todo-api-go/pkg/configs"
	"guthub.com/iribuda/todo-api-go/pkg/models"
	"guthub.com/iribuda/todo-api-go/pkg/services/auth"
	"guthub.com/iribuda/todo-api-go/pkg/utils"
)

// Controller für das Anmelden und die Registrierung
type UserController struct{
	repository models.UserRepository
}

// Konstruktor
func NewController(repository models.UserRepository) *UserController{
	return &UserController{repository: repository}
}

//  Deklaration der Routen
func (uc *UserController) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", uc.handleLogin).Methods("POST")
	router.HandleFunc("/register", uc.handleRegister).Methods("POST")
}

// Handler für das Anmelden
func (uc *UserController) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Aus JSON wird DTO abgerufen
	var user models.LoginUserDTO
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Die Daten werden validiert
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Abrufen des existierten Benutzer mit der Email aus der DB
	u, err := uc.repository.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	// Überprüfen, ob die Passworts stimmen
	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	// Token wird generiert
	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// Handler für die Registrierung
func (uc *UserController) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Aus JSON wird DTO abgerufen
	var user models.RegisterUserDTO
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Die Daten werden validiert
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user credentials: %v", errors))
		return
	}

	// Überpüfen, ob Benutzer mit solcher Email schon in der DB ist
	_, err := uc.repository.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// Passwort wird hashed
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Speicher des Benutzers
	err = uc.repository.CreateUser(models.User{
		Username:  user.Username,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}