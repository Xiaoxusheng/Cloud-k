package controller

import (
	"Cloud-k/db"
	"Cloud-k/uility"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 刷新token后必须使用最新token请求，否则认为在其他地方登录
func RefreshToken(c *gin.Context) {
	userIdentity := c.MustGet("UserIdentity").(string)
	token := uility.GetToken(userIdentity, 2)
	refresh_token := uility.GetToken(userIdentity, 24)
	ctx := context.Background()
	res, err := db.Rdb.HSet(ctx, userIdentity, "token", token, "refresh_token", refresh_token).Result()
	fmt.Println(res)
	db.Rdb.Expire(ctx, userIdentity, time.Hour*24)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDescription: "redis获取失败!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功!",
		"data": gin.H{
			"token":         token,
			"refresh_token": refresh_token,
		},
	})

}
