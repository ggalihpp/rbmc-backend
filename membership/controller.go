package membership

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func hashString(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}

func checkHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}

func getUsername(email string) (username string) {
	username = strings.Split(email, "@")[0]
	return
}

func testController() (output string, err error) {
	var x jwt.SigningMethod

	method := jwt.SigningMethodHS256
	token := jwt.New(method)

	claims := token.Claims.(jwt.MapClaims)

	claims["FOO"] = "BAR"
	claims["isAdmin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	output, _ = token.SignedString([]byte("SECRET"))

	fmt.Println(x.Verify(method.Alg(), output, []byte("SECRET")))

	return
}
