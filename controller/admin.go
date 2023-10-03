package controller

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Banned 封禁用户
func Banned(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	//	判断是否存在
	ok := models.GetUserById(identity)
	if !ok {
		panic(uility.ErrorMessage{
			ErrorDescription: "用户不存在!",
		})
	}
	models.UpdateStatus(identity, 1)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "封禁用户成功",
	})
}

// 解封用户
func Unseal(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	//	判断是否存在
	ok := models.GetUserById(identity)
	if !ok {
		panic(uility.ErrorMessage{
			ErrorDescription: "用户不存在!",
		})
	}
	models.UpdateStatus(identity, 0)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "解封用户成功",
	})
}

//分配容量

//查看剩余容量

//查看用户数据

//查看用户访问信息

//系统操作日志查看
