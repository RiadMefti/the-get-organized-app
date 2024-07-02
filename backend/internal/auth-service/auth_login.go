package auth

import (
	"backend/internal/database"
	"backend/internal/types"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func validatePassword(password string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func AuthentificateUser(db database.Service, userLogin types.UserLogin) (string, error) {

	user, err := db.GetUserByEmail(userLogin.Email)
	if err != nil {
		return "", err
	}

	if !validatePassword(userLogin.Password, user.HashedPassword) {
		return "", errors.New("invalid password")
	}

	return user.ID.String(), nil
}
