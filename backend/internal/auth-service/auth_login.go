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

func AuthentificateUser(db database.Service, userLogin types.UserLogin) (bool, error) {

	user, err := db.GetUserByEmail(userLogin.Email)
	if err != nil {
		return false, err
	}

	if !validatePassword(userLogin.Password, user.HashedPassword) {
		return false, errors.New("invalid password")
	}

	return true, nil
}
