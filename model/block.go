package model

import "github.com/golang-jwt/jwt"

type Block struct {
	Height  int    `db:"height" json:"height"`
	Hash    string `db:"hash" json:"hash"`
	Parent  string `db:"parent" json:"parent"`
	Encoded []byte `db:"encoded" json:"encoded"`
	Status  string `db:"status" json:"status"`
}

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string `db:"token"`
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
