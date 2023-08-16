package uility

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"os"
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

// 记录日志文件
func CreateLogFile() {
	//现在的时间
	t := time.Now()
	//明天0点时间
	t1 := time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, time.Local)
	//t1 := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Local)
	times := time.NewTicker(t1.Sub(t))
	f, _ := os.OpenFile("./log/"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(f)

	for {
		select {
		case <-times.C:
			// 记录到文件。
			f, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			//f, err := os.Create("./log/" + time.Now().Format("2006-01-02") + ".log")
			// 如果需要同时将日志写入文件和控制台，请使用以下代码。

			if err == nil {
				log.Println(time.Now().Format("2006-01-02") + "log文件创建成功！")
				//现在的时间
				t = time.Now()
				//明天0点时间
				t1 = time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, time.Local)

				times = time.NewTicker(t1.Sub(t))
			}
			fmt.Println(err)
			f.Close()

		}
	}

}
