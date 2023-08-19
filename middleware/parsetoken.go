package middleware

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// 解析token
func ParseToken() gin.HandlerFunc {
	ErrorMessage := uility.ErrorMessage{}
	return func(c *gin.Context) {
		// Token from another example.  This token is expired
		tokens := c.GetHeader("Authorization")
		if tokens == "" {
			c.Abort()
			ErrorMessage.ErrorDescription = "token不能为空"
			panic(ErrorMessage)
		}
		//fmt.Println(strings.Contains(tokens, "Bearer"))
		if !strings.Contains(tokens, "Bearer") {
			c.Abort()
			ErrorMessage.ErrorDescription = "token格式不对!"
			panic(ErrorMessage)
			return
		}
		tokenString := strings.Split(tokens, "Bearer ")
		fmt.Println(tokenString)
		user := uility.MyCustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString[len(tokenString)-1], &user, func(token *jwt.Token) (interface{}, error) {
			return uility.MySigningKey, nil
		})
		//if err != nil {
		//	c.Abort()
		//	ErrorMessage.ErrorDescription = "令牌格式不对!"
		//	panic(ErrorMessage)
		//}

		if token.Valid {
			fmt.Println("验证通过")
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			fmt.Println("That's not even a token")
			c.Abort()
			ErrorMessage.ErrorDescription = "这不是一个token!"
			panic(ErrorMessage)

		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			// Invalid signature
			fmt.Println("Invalid signature")
			c.Abort()
			ErrorMessage.ErrorDescription = "无效的签名!"
			panic(ErrorMessage)
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			c.Abort()
			ErrorMessage.ErrorDescription = "token失效或过期!"
			panic(ErrorMessage)
		} else {
			fmt.Println("Couldn't handle this token:", err)
			c.Abort()
			ErrorMessage.ErrorDescription = "未知错误!"
			panic(ErrorMessage)
		}

		c.Set("UserIdentity", user.Identity)
		//fmt.Println("username", user.Identification)
		ctx := context.Background()
		exit, _ := db.Rdb.Exists(ctx, user.Identity).Result()
		fmt.Println("exit", exit)
		if exit != 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  "已经退出登录!",
			})
			c.Abort()
			return
		}
		t, err := db.Rdb.HGet(ctx, user.Identity, "token").Result()
		rt, err := db.Rdb.HGet(ctx, user.Identity, "refresh_token").Result()
		if err != nil {
			ErrorMessage.ErrorDescription = "redis获取失败!"
			panic(ErrorMessage)
		}
		//fmt.Println(t, "\n", rt)
		if t != tokenString[len(tokenString)-1] && rt != tokenString[len(tokenString)-1] {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "账号在其他地方登录，只允许一台设备登录！",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
