package auth

import (
	"backend/internal/database"
	"backend/internal/types"
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func validateUserRegistration(db database.Service, user types.UserRegistration) error {
	ok, err := isPasswordMatch(user.Password, user.CopyPassword)
	if !ok {
		return err
	}

	ok, err = isPasswordStrongEnough(user.Password)
	if !ok {
		return err
	}

	exists, err := emailExists(db, user.Email)
	log.Print(exists)
	if exists {
		return errors.New("email already exists")
	}
	if err != nil {
		return errors.New("error checking if email exists")
	}
	err = isValidEmail(user.Email)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(db database.Service, user types.UserRegistration) error {
	err := validateUserRegistration(db, user)
	if err != nil {
		return err
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	err = db.CreateUser(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func isPasswordMatch(password string, copyPassword string) (bool, error) {

	if password != copyPassword {
		return false, errors.New("the passwords do not match")

	}
	return true, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func isPasswordStrongEnough(password string) (bool, error) {
	if len(password) < 8 {
		return false, errors.New("password is not strong enough , please make sure that the password is at least 8 characters long")
	}
	return true, nil
}

func emailExists(db database.Service, email string) (bool, error) {
	user, err := db.GetUserByEmail(email)

	if user.Email == email {
		return true, errors.New("email already exists")
	}
	if err != nil {
		log.Println(err)
		return false, err
	}

	return false, nil
}

func isValidEmail(email string) error {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return nil
	}
	return errors.New("email is not a valid email")
}
