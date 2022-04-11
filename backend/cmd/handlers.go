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

// createJWTCookie creates a jwt and sets a cookie.
func createJWTCookie(username string) (http.Cookie, error) {
	// generate new json web token
	jwtToken, err := createToken(username)
	if err != nil {
		return http.Cookie{}, err
	}

	// create new cookie
	jwtCookie := &http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Hour * time.Duration(AppConfig.TokenDuration)),
		Path:     "/api",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	return *jwtCookie, nil
}

// checkJWTCookie checks if a cookie has the correct jwt.
func checkJWTCookie(r *http.Request) (username string, err error) {
	// get cookie from request
	jwtCookie, err := r.Cookie("token")
	if err != nil {
		return username, err
	}

	// validate json web token
	username, err = authenticationFromToken(jwtCookie.Value)
	if err != nil {
		log.Println(err)
		return username, errors.New("Authentication token isn't valid")
	}

	// check if user exists in database
	userExists, err := CheckUsername(username)
	if err != nil {
		log.Println(err)
		return username, err
	}

	if !userExists {
		return username, errors.New("User doesn't exist")
	}

	return username, nil
}

// signUpHandler handles user sign up.
func signUpHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	var userSignUp SignUp

	// decode request body
	err := json.NewDecoder(r.Body).Decode(&userSignUp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate username
	err = validateUsername(userSignUp.Username)
	if err != nil {
		errorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate user information
	err = validateUser(userSignUp.User)
	if err != nil {
		errorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate authentication password
	if userSignUp.AuthPassword != AppConfig.AuthenticationPassword {
		errorJSON(w, "Authentication password is incorrect", http.StatusForbidden)
		return
	}

	// validate password
	err = validatePassword(userSignUp.Password)
	if err != nil {
		errorJSON(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	// hash password
	hash, err := hashPassword(userSignUp.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userSignUp.Password = hash

	// add new user in database
	err = AddNewUser(userSignUp.Account)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("User %s was successfully registered", userSignUp.Username)}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// signInHandler handles login and returns a jwt for later authentication.
func signInHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	var credentials Credentials

	// decode request body
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check user credentials in database
	userExists, err := CheckUsername(credentials.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check password in database
	hash, err := GetPassword(credentials.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !checkPasswordHash(credentials.Password, hash) || !userExists {
		errorJSON(w, "Username or password is incorrect", http.StatusForbidden)
		return
	}

	// create new cookie
	jwtCookie, err := createJWTCookie(credentials.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &jwtCookie)

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "Successfully signed in"}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// signOutHandler overwrites the httpOnly cookie to sign out the user.
func signOutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	// new empty cookie
	newCookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/api",
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, &newCookie)

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "Successfully signed out"}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getUsersHandler checks authentication and returns all users from the database.
func getUsersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	// authenticate user
	_, err := checkJWTCookie(r)
	if err != nil {
		errorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// get users from database
	users, err := GetAllUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode response to json
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getAccountHandler checks authentication and returns one user from the username
func getAccountHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	// authenticate user
	username, err := checkJWTCookie(r)
	if err != nil {
		errorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// get user from database
	user, err := GetUser(username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode response to json
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// updatePasswordHandler updates an account in the database
func updatePasswordHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	var newPassword struct {
		Password string `json:"password"`
	}

	// decode request body
	err := json.NewDecoder(r.Body).Decode(&newPassword)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// authenticate user
	username, err := checkJWTCookie(r)
	if err != nil {
		errorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// validate new password
	err = validatePassword(newPassword.Password)
	if err != nil {
		errorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// hash new password
	hash, err := hashPassword(newPassword.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newPassword.Password = hash

	// update password in database
	err = UpdatePassword(newPassword.Password, username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: "Password succesfully updated"}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// updateUserHandler updates user information.
func updateUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	var newUser User

	// decode request body
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// authenticate user
	username, err := checkJWTCookie(r)
	if err != nil {
		errorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// validate new user information
	err = validateUser(newUser)
	if err != nil {
		errorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.Username = username

	// update user information in database
	err = UpdateUser(newUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("User %s was successfully updated", newUser.Username)}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// updateUsernameHandler updates the username.
func updateUsernameHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	var newUsername struct {
		Username string `json:"username"`
	}

	// decode request body
	err := json.NewDecoder(r.Body).Decode(&newUsername)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// authenticate user
	oldUsername, err := checkJWTCookie(r)
	if err != nil {
		errorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// validate new username
	err = validateUsername(newUsername.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update username in database
	err = UpdateUsername(newUsername.Username, oldUsername)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set new cookie for authentication
	jwtCookie, err := createJWTCookie(newUsername.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &jwtCookie)

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("Username was successfully changed from %s to %s", oldUsername, newUsername.Username)}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// deleteAccountHandler handles account deletions.
func deleteAccountHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	// authenticate user
	username, err := checkJWTCookie(r)
	if err != nil {
		errorJSON(w, "Access denied", http.StatusForbidden)
		return
	}

	// delete account in database
	err = DeleteAccount(username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// erase cookie for authentication
	newCookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/api",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(w, &newCookie)

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("Account for %s is deleted", username)}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
