package model

import "github.com/dgrijalva/jwt-go"

// Login estructura para el inicio de sesión
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claim es el cuerpo del token
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
