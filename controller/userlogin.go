package controller

import (
	"Cloud-k/db"
	"Cloud-k/models"
	"Cloud-k/uility"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(c.RemoteIP())

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	user, err := models.GetUser(username, uility.GetMd5(password))
	fmt.Println(user, err != nil)

	if user.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户不存在！",
		})
		return
	}
	m := 2
	n := 24
	token := uility.GetToken(user.Identity, user.CasbinIdentity, m)
	refresh_token := uility.GetToken(user.Identity, user.CasbinIdentity, n)
	//设置用户唯一identity
	ctx := context.Background()
	res, err := db.Rdb.HSet(ctx, user.Identity, "token", token, "refresh_token", refresh_token).Result()
	fmt.Println(res)
	db.Rdb.Expire(ctx, user.Identity, time.Hour*24)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDescription: "redis获取失败!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功！",
		"data": gin.H{
			"token":         token,
			"refresh_token": refresh_token,
		},
	})

}

// 用户注册
func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if username == "" || password == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	if ok := models.GetEmail(email); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "邮箱已经存在！",
		})
		return
	}

	if models.GetByUser(username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "用户名已经存在！",
		})
		return
	}

	models.InsertUser(username, uility.GetMd5(password), uility.GetUuid(), email)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功！",
	})

}

// 用户详情
func UserDetail(c *gin.Context) {
	identity, ok := c.Get("UserIdentity")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "获取用户详情失败！",
		})
		return
	}
	userDetil := models.GetUserDetail(identity.(string))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功！！",
		"data": gin.H{
			"userdetil": userDetil,
		},
	})

}

// 退出登录
func Logout(c *gin.Context) {
	ctx := context.Background()
	UserIdentity := c.MustGet("UserIdentity").(string)
	result, err := db.Rdb.Del(ctx, UserIdentity).Result()
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDescription: "redis获取失败!",
		})
	}
	if result == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "退出成功!",
		})
	}

}
