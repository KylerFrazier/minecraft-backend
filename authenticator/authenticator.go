package authenticator

import (
	"minecraft-backend/databaseWrapper"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const jwtExpirationTime = 1 * time.Hour

var accessTokenSecret = []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
var refreshTokenSecret = []byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET"))

type Authenticator struct {
	database *databaseWrapper.DatabaseWrapper
}

func (auth *Authenticator) login(userName string, password string) (string, error) {
	userHash := auth.database.GetUserHash(userName)
	err := bcrypt.CompareHashAndPassword([]byte(auth.database.GetPasswordHash(userHash)), []byte(password))
	if err != nil {
		return "", err
	}

	type Claims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpirationTime)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "KaiNoSekai",
		Subject:   "User",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(accessTokenSecret)
	return signedString, err
}
