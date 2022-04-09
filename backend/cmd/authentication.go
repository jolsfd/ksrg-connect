package cmd

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes a password.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash compares if a password and a hash matchs.
func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Validate validates username.
func validateUsername(username string) (err error) {
	// Validate username
	if username == "" {
		return errors.New("You must provide an username")
	}

	if strings.Contains(username, " ") {
		return errors.New("No spaces are allowed in the username")
	}

	if len(username) > 20 {
		return errors.New("20 characters is the max length for an username")
	}

	// Check username in database
	userExists, err := CheckUsername(username)
	if err != nil {
		return err
	}
	if userExists {
		return errors.New("Username exists already")
	}

	return nil
}

// validateUser validates user information
func validateUser(user User) error {
	if len(user.FirstName) < 1 {
		return errors.New("You must provide a first name")
	}

	if len(user.LastName) < 1 {
		return errors.New("You must provide a last name")
	}

	return nil
}

// validatePassword validates if an user has provided a correct password.
func validatePassword(password string) (err error) {
	if password == "" {
		return errors.New("You must provide a password")
	}

	if len(password) < 8 {
		return errors.New("Choose a password with at least 8 characters")
	}

	if len(password) > 40 {
		return errors.New("40 characters is the max length for passwords")
	}

	return nil
}

// createToken creates an jwt token with a username as claim. The token is 2 days valid.
func createToken(username string) (tokenString string, err error) {
	validTime := time.Now().Add(time.Hour * time.Duration(AppConfig.TokenDuration))

	// Create json web token
	claims := CustomClaims{username, jwt.StandardClaims{ExpiresAt: validTime.Unix(), Issuer: AppConfig.Domain}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(AppConfig.SecretString))

	return tokenString, err
}

// authenticationFromToken validates an jwt and returns the username for the given token.
func authenticationFromToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(AppConfig.SecretString), nil
	})

	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims.Username, nil
	} else {
		return "", err
	}
}

// checkAdmin checks if an user is admin.
func checkAdmin(username string, admins []string) bool {
	for _, admin := range admins {
		if username == admin {
			return true
		}
	}

	return false
}
