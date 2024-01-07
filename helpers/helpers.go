package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// receives: err
// will be used to handle all of our errors and
// in order to obey the DRY principle
// returns: None
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// receives: byte
// method is used to hash passwords to cryptic form
// in order to prevent password leak incase of a data
// breach
// returns: string
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

// authenticate function is used to authenticate a user
// that logs in. It works by matching the username and
// password received from the frontend with the database
// If a match is found, then we generate a jwt, populate
// it with the user's detail and send to the client.
func GenerateJWT(Username string) string {
	var secretKey = []byte("kenechukwusecret")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = Username

	tokenString, err := token.SignedString(secretKey)
	HandleErr(err)

	return tokenString
}
