package test

import (
	"Cloud-k/uility"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

type MyCustomClaims struct {
	Identity string
	jwt.RegisteredClaims
}

func TestToken(t *testing.T) {
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"nsiunfg",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(2) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("welcome to use Cloud-kAuth:Mr.Lei"))
	if err != nil {
		panic(uility.ErrorMessage{
			err.Error(),
			err.Error(),
			"GetToken函数",
			time.Now(),
		})
	}
	fmt.Println(ss)
}
