package cmd

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	// Post handler
	router.POST("/api/signup", signUpHandler)
	router.POST("/api/signin", signInHandler)
	router.POST("/api/signout", signOutHandler)

	router.POST("/api/update/username", updateUsernameHandler)
	router.POST("/api/update/password", updatePasswordHandler)
	router.POST("/api/update/user", updateUserHandler)
	router.POST("/api/delete", deleteAccountHandler)

	// Get handler
	router.GET("/api/users", getUsersHandler)
	router.GET("/api/profile", getAccountHandler)

	// Admin Handler
	router.POST("/api/admin/delete", adminDeleteUserHandler)

	return router
}
