package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// delete users and add in config
func AdminDeleteUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	var deleteUser struct {
		Username string `json:"deleteUsername"`
	}

	// Decode request to struct
	err := json.NewDecoder(r.Body).Decode(&deleteUser)
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

	// Check if user is admin
	if !CheckAdmin(username, AppConfig.Admins) {
		ErrorJSON(w, "You are not an admin", http.StatusForbidden)
		return
	}

	userExists, err := CheckUsername(deleteUser.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !userExists {
		ErrorJSON(w, "User doesn't exist", http.StatusForbidden)
		return
	}

	// Delete user from database
	err = DeleteAccount(deleteUser.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode to json
	err = json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("Account for %s is deleted", deleteUser.Username)})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
