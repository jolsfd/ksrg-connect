package cmd

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares if a password and a hash matchs.
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Validate validates username.
func ValidateUsername(username string) (err error) {
	// Validate username
	if username == "" {
		return errors.New("you must provide an username")
	}

	if strings.Contains(username, " ") {
		return errors.New("no spaces are allowed in the username")
	}

	if len(username) > 20 {
		return errors.New("20 characters is the max length for usernames")
	}

	// Check username in database
	userExists, err := CheckUsername(username)
	if err != nil {
		return err
	}
	if userExists {
		return errors.New("username exists already")
	}

	return nil
}

// ValidateUser validates user information
func ValidateUser(user User) error {
	if len(user.FirstName) < 1 {
		return errors.New("you must provide a first name")
	}

	if len(user.LastName) < 1 {
		return errors.New("you must provide a last name")
	}

	return nil
}

// ValidatePassword validates if an user has provided a correct password.
func ValidatePassword(password string) (err error) {
	if password == "" {
		return errors.New("you must provide a password")
	}

	if len(password) < 8 {
		return errors.New("choose a password with at least 8 characters")
	}

	if len(password) > 40 {
		return errors.New("40 characters is the max length for passwords")
	}

	return nil
}

// CreateToken creates an jwt token with a username as claim. The token is 2 days valid.
func CreateToken(username string) (tokenString string, err error) {
	// Generate Unix time
	// tokenDuration, err := time.ParseDuration("24h")
	// if err != nil {
	// 	return tokenString, err
	// }

	validTime := time.Now().Add(time.Hour * time.Duration(AppConfig.TokenDuration))

	// Create json web token
	claims := CustomClaims{username, jwt.StandardClaims{ExpiresAt: validTime.Unix(), Issuer: AppConfig.Domain}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(AppConfig.SecretString))

	return tokenString, err
}

// AuthenticationFromToken validates an jwt and returns the username for the given token.
func AuthenticationFromToken(tokenString string) (string, error) {
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

// CheckAdmin checks if an user is admin.
func CheckAdmin(username string, admins []string) bool {
	for _, admin := range admins {
		if username == admin {
			return true
		}
	}

	return false
}
