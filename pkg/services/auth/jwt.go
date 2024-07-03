package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"guthub.com/iribuda/todo-api-go/pkg/configs"
	"guthub.com/iribuda/todo-api-go/pkg/models"
	"guthub.com/iribuda/todo-api-go/pkg/utils"
)

type contextKey string

const UserKey contextKey = "userID"

func WithJWTAuth(handlerFunc http.HandlerFunc, repository models.UserRepository) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		tokenString := utils.GetTokenFromRequest(r)
		
		token, err := validateJWT(tokenString)

		if err != nil{
			log.Printf("failed to validate token: %v", err)
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
			return
		}

		if !token.Valid{
			log.Println("invalid token")
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, err := strconv.Atoi(str)
		if err != nil{
			log.Printf("failed to convert userID to int: %v", err)
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
			return
		}

		u, err := repository.GetUserByID(userID)
		if err != nil{
			log.Printf("failed to get user by id: %v", err)
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
			return
		}

		// Benutzer wird dem Context hinzugefügt
		currentContext := r.Context()
		currentContext = context.WithValue(currentContext, UserKey, u.UserID)
		r = r.WithContext(currentContext)

		// Die Funktion wird aufgerufen, falls Token valid ist
		handlerFunc(w, r)
	}
}

func CreateJWT(secret []byte, userID int)(string, error){
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil{
		return "", err
	}

	return tokenString, err
}

func validateJWT(tokenString string) (*jwt.Token, error){
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signign method: %v", token.Header["alg"])
		}

		return []byte(configs.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(currentContext context.Context) int{
	userID, ok := currentContext.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userID
}