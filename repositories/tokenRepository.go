package repositories

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"storeApiRest/database"
	"storeApiRest/models"
	"time"
)

var (
	JwtKey = []byte("secret_key")
)

func AuthenticateUser(email, password string) (*models.AuthResponse, error) {
	query := `SELECT user_email, user_password 
				FROM users WHERE user_email = $1`
	db := database.GetConnection()
	defer db.Close()
	var user models.User
	err := db.QueryRow(query, email).Scan(&user.UserEmail, &user.UserPassword)
	if user.UserPassword != password {
		return nil, errors.New("invalid credentials")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{Token: tokenString}, nil
}
