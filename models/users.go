package models

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
)

type FormRegister struct {
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

type FormAuth struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" valid:"required"`
}

type GetUserScanner struct {
	ID       sql.NullInt64  `db:"id"`
	Username sql.NullString `db:"username"`
	Email    sql.NullString `db:"email"`
	Password sql.NullString `db:"password"`
}

type MyClaims struct {
	jwt.StandardClaims
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthResult struct {
	Token string `json:"token"`
}
