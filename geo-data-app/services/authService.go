
package services

import (
	"geo-data-app/models"
	"geo-data-app/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	_, err := utils.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func AuthenticateUser(credentials models.User) (string, error) {
	var storedUser models.User
	err := utils.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", credentials.Email).Scan(&storedUser.ID, &storedUser.Password)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	token := utils.GenerateJWT(storedUser.ID)
	return token, nil
}
