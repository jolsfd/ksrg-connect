package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// delete users and add in config
func adminDeleteUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeaders(w)

	var deleteUser struct {
		Username string `json:"deleteUsername"`
	}

	// decode request body
	err := json.NewDecoder(r.Body).Decode(&deleteUser)
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

	// authenticate admin
	if !checkAdmin(username, AppConfig.Admins) {
		errorJSON(w, "You are not an admin", http.StatusForbidden)
		return
	}

	// check if user exists
	userExists, err := CheckUsername(deleteUser.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !userExists {
		errorJSON(w, "User doesn't exist", http.StatusForbidden)
		return
	}

	// delete user from database
	err = DeleteAccount(deleteUser.Username)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode response to json
	if err := json.NewEncoder(w).Encode(ApiMessage{Success: true, Message: fmt.Sprintf("Account for %s is deleted", deleteUser.Username)}); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
