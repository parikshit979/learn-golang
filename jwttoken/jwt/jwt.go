package jwt

import (
	"crypto/rsa"
	"errors"
	"os"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
)

const (
	authExpirationTime time.Duration = 15 * time.Minute
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

// InitJWT initializes the JWT by reading the RSA private and public keys from the specified file paths.
// It reads the private key from "path/to/private.key" and the public key from "path/to/public.key".
// If there is an error reading the files or parsing the keys, the function will panic.
func InitJWT() {
	privateKeyData, err := os.ReadFile("keys/rsa_private.key")
	if err != nil {
		panic(err)
	}

	privateKey, err = jwtGo.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		panic(err)
	}

	publicKeyData, err := os.ReadFile("keys/rsa_public.key")
	if err != nil {
		panic(err)
	}

	publicKey, err = jwtGo.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		panic(err)
	}
}

type Claims struct {
	Username string `json:"username"`
	jwtGo.StandardClaims
}

// GenerateJWT generates a JSON Web Token (JWT) for a given username.
// The token is signed using the RS256 signing method and includes an
// expiration time of 5 minutes.
func GenerateJWT(username string) (string, *jwtGo.Token, error) {
	expirationTime := time.Now().Add(authExpirationTime)
	claims := &Claims{
		Username: username,
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", nil, err
	}
	return tokenString, token, nil
}

// ParseJWT parses a JWT token string and returns the claims if the token is valid.
// It takes a token string as input and returns a pointer to Claims and an error.
// If the token is invalid or the signing method is unexpected, it returns an error.
func ParseJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwtGo.ParseWithClaims(tokenString, claims, func(token *jwtGo.Token) (any, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
