package controller

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddPermission 添加用户权限
func AddPermission(c *gin.Context) {
	//用户的唯一id
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	//添加权限
	models.InsertAdmin(identity)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "增加权限成功！",
	})
	fmt.Println(c.Request.Method)

}

// UpdatePermission 修改用户权限
func UpdatePermission(c *gin.Context) {
	identity := c.Query("identity")
	status := c.Query("status")
	s, err := strconv.Atoi(status)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDescription: "转换失败!",
		})
		return
	}
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	list := map[int]string{
		0: uility.User,
		1: uility.Admin,
		2: uility.Root,
	}
	//	判断id是否存在
	models.GetUserPermission(identity)
	//    更新
	models.UpdatePermission(identity, list[s])
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改权限成功！",
	})
}
