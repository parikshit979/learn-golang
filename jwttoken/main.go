package main

import (
	"log"

	"github.com/learn-golang/jwttoken/jwt"
)

func main() {
	log.Println("Starting server")
	jwt.InitJWT()
	tokenString, token, err := jwt.GenerateJWT("parikshit979")
	if err != nil {
		log.Println("Failed to generate token")
	}

	log.Printf("TokenString(%s), token(%v)", tokenString, token.Claims)
}
