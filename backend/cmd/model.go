package cmd

import "github.com/golang-jwt/jwt"

// User struct handles user information.
type User struct {
	Username    string `json:"username"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Class       string `json:"schoolClass"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
}

type SignUp struct {
	Account
	AuthPassword string `json:"authPassword"`
}

// Account struct has a user struct with account related things like password.
type Account struct {
	User
	Password string `json:"password"`
}

// Credentials handles information for login.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CustomClaims handles data for json web tokens.
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type ApiMessage struct {
	Success bool   `json:"success"`
	Message string `json:"msg"`
}

// Struct for configuration
type Config struct {
	Port                   string   `mapstructure:"port"`
	Domain                 string   `mapstructure:"domain"`
	AllowedOrigins         []string `mapstructure:"origins"`
	PathToDB               string   `mapstructure:"database"`
	TokenDuration          int      `mapstructure:"duration"`
	SecretString           string   `mapstructure:"secret"`
	AuthenticationPassword string   `mapstructure:"password"`
	Admins                 []string `mapstructure:"admins"`
}
