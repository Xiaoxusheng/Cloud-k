package controller

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

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

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功！",
		"data": gin.H{
			"token": uility.GetToken(user.Identity),
		},
	})
}

// 用户注册
func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if username == "" || password == "" || email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	if ok := models.GetEmail(email); ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "邮箱已经存在！",
		})
		return
	}

	if models.GetByUser(username) {
		c.JSON(http.StatusOK, gin.H{
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
	identity := c.MustGet("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
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
