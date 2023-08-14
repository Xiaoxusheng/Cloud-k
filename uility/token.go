package uility

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"time"
)

//生成token

type MyCustomClaims struct {
	Identity string
	jwt.RegisteredClaims
}

func GetToken(Identity string) string {
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		Identity,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySigningKey)
	if err != nil {
		panic(ErrorMessage{
			Error,
			err.Error(),
			"GetToken函数",
			time.Now(),
		})
	}
	//生成token同时存入redis
	//result, err := db.Rdb.Set(ctx, identification, ss, time.Hour*24).Result()
	//if err != nil {
	//	return ""
	//}
	//fmt.Println(result)

	fmt.Printf("%v ", ss)
	return ss
}

// uuid生成
func GetUuid() string {
	return uuid.NewV4().String()
}

func GetMd5(pwd string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}
