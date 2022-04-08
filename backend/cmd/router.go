package cmd

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	// Post handler
	router.POST("/api/signup", SignUpHandler)
	router.POST("/api/signin", SignInHandler)
	router.POST("/api/signout", SignOutHandler)

	router.POST("/api/update/username", UpdateUsernameHandler)
	router.POST("/api/update/password", UpdatePasswordHandler)
	router.POST("/api/update/user", UpdateUserHandler)
	router.POST("/api/delete", DeleteUserHandler)

	// Get handler
	router.GET("/api/users", GetUsersHandler)
	router.GET("/api/profile", GetProfile)

	// Admin Handler
	router.POST("/api/admin/delete", AdminDeleteUserHandler)

	return router
}
