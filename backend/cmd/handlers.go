package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// ErrorJSON returns an api message with the error as json
func ErrorJSON(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiMessage{Success: false, Message: error})
}

// CreateJWTCookie creates a jwt and sets a cookie.
func CreateJWTCookie(username string) (http.Cookie, error) {
	// Generate jwt
	jwtToken, err := CreateToken(username)
	if err != nil {
		return http.Cookie{}, err
	}

	// Create jwtCookie
	jwtCookie := &http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * time.Duration(AppConfig.TokenDuration)),
		Path:     "/api",
		Secure:   false,
		HttpOnly: true,
	}

	return *jwtCookie, nil
}

// CheckJWTCookie checks if a cookie has the correct jwt.
func CheckJWTCookie(r *http.Request) (username string, err error) {
	jwtCookie, err := r.Cookie("token")
	if err != nil {
		return username, err
	}

	username, err = AuthenticationFromToken(jwtCookie.Value)
	if err != nil {
		log.Println(err)
		return username, errors.New("json web token isn't valid")
	}

	userExists, err := CheckUsername(username)
	if err != nil {
		log.Println(err)
		return username, err
	}

	if !userExists {
		return username, errors.New("user doesn't exist")
	}

	return username, nil
}

// SignUpHandler handles user sign up.
func SignUpHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userSignUp SignUp

	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Decode request to struct
	err := json.NewDecoder(r.Body).Decode(&userSignUp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Make Account struct from request
	userAccount := userSignUp.Account

	// Validate user information
	err = ValidateUsername(userAccount.Username)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusNotAcceptable)
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user information
	err = ValidateUser(userAccount.User)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check AuthPassword
	if userSignUp.AuthPassword != AppConfig.AuthenticationPassword {
		ErrorJSON(w, "authentication password is incorrect", http.StatusForbidden)
		return
	}

	err = ValidatePassword(userAccount.Password)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	// Hash password
	hash, err := HashPassword(userAccount.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set new password in account
	userAccount.Password = hash

	// Insert user struct in the database
	err = AddNewUser(userAccount)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return message for correct signup
	err = json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("user %s was successfully registered", userAccount.Username)})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// SignInHandler handles login and returns a jwt for later authentication.
func SignInHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var credentials Credentials

	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check user credentials from database
	userExists, err := CheckUsername(credentials.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !userExists {
		ErrorJSON(w, "User does not exist", http.StatusForbidden)
		return
	}

	hash, err := GetPassword(credentials.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !CheckPasswordHash(credentials.Password, hash) {
		ErrorJSON(w, "password does not match", http.StatusForbidden)
		return
	}

	jwtCookie, err := CreateJWTCookie(credentials.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &jwtCookie)

	// Return response
	err = json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "Successfully signed in"})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetUsersHandler checks authentication and returns all users from the database.
func GetUsersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Validate authentication jwt
	_, err := CheckJWTCookie(r)
	if err != nil {
		ErrorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// Get users from database
	users, err := GetAllUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode to json
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetProfile checks authentication and returns one user from the username
func GetProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Validate authentication jwt
	username, err := CheckJWTCookie(r)
	if err != nil {
		ErrorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// Get user from database
	user, err := GetUser(username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode to json
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdatePasswordHandler updates an account in the database
func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var newPassword struct {
		Password string `json:"password"`
	}

	// Decode request to struct
	err := json.NewDecoder(r.Body).Decode(&newPassword)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check user credentials user from jwt
	username, err := CheckJWTCookie(r)
	if err != nil {
		ErrorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// Check Password
	err = ValidatePassword(newPassword.Password)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	hash, err := HashPassword(newPassword.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set new password in account
	newPassword.Password = hash

	// Update Account in the database
	err = UpdatePassword(newPassword.Password, username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode to json
	err = json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "Password succesfully updated"})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateUserHandler updates user information.
func UpdateUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var userUpdate User

	// Decode request to struct
	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate authentication jwt
	username, err := CheckJWTCookie(r)
	if err != nil {
		ErrorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	userUpdate.Username = username

	// Update in database
	err = UpdateUser(userUpdate)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode to json
	err = json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "User was successfully updated"})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateUsernameHandler updates the username.
func UpdateUsernameHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var newUsername struct {
		Username string `json:"username"`
	}

	// Decode request to struct
	err := json.NewDecoder(r.Body).Decode(&newUsername)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate authentication jwt
	oldUsername, err := CheckJWTCookie(r)
	if err != nil {
		ErrorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// Validate new username
	err = ValidateUsername(newUsername.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update username in database
	err = UpdateUsername(newUsername.Username, oldUsername)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set new cookie
	jwtCookie, err := CreateJWTCookie(newUsername.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &jwtCookie)

	// Encode to json
	err = json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "Username was successfully updated"})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteUserHandler handles account deletions.
func DeleteUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check user credentials user from jwt
	username, err := CheckJWTCookie(r)
	if err != nil {
		ErrorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// Delete user from database
	err = DeleteAccount(username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newCookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/api",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(w, &newCookie)

	// Encode to json
	err = json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("Account for %s is deleted", username)})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// SignOutHandler overwrites the httpOnly cookie to sign out the user.
func SignOutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	// Create empty jwtCookie
	newCookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/api",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, &newCookie)

	err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "Successfully signed out"})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
