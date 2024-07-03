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

type UserController struct{
	repository models.UserRepository
}

func NewController(repository models.UserRepository) *UserController{
	return &UserController{repository: repository}
}

func (uc *UserController) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", uc.handleLogin).Methods("POST")
	router.HandleFunc("/register", uc.handleRegister).Methods("POST")
}

func (uc *UserController) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user models.LoginUserDTO
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := uc.repository.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (uc *UserController) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user models.RegisterUserDTO
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user credentials: %v", errors))
		return
	}

	// check if user exists
	_, err := uc.repository.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

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