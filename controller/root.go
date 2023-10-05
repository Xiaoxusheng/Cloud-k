package controller

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	//验证是否存在
	ok := models.GetUserById(identity)
	if !ok {
		panic(uility.ErrorMessage{
			ErrorDescription: "用户不存在!",
		})
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

// UpdateAssets 修改用户可以访问的资源
func UpdateAssets(c *gin.Context) {
	id := c.Query("id")
	path := c.Query("path")
	methods := strings.ToUpper(c.Query("methods"))
	newPath := c.Query("newPath")
	newmethods := strings.ToUpper(c.Query("newmethods"))
	if id == "" || path == "" || methods == "" || newPath == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	ok := uility.E.HasPolicy(id, path, methods)
	if !ok {
		panic(uility.ErrorMessage{
			ErrorDescription: "策略不存在!",
		})
	}
	if newmethods == "" {
		newmethods = methods
	}

	fmt.Println(ok)
	_, err := uility.E.UpdatePolicy([]string{id, path, methods}, []string{id, newPath, newmethods})
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDescription: "修改失败!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功！",
	})
}

// GetAssetsList 查看所有权限
func GetAssetsList(c *gin.Context) {
	list := uility.E.GetNamedPolicy("p")
	fmt.Println(list)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功！",
		"data": gin.H{
			"assets_list": list,
		},
	})
}

// DeleteAssets 删除访问的路径资源
func DeleteAssets(c *gin.Context) {
	id := c.Query("id")
	path := c.Query("path")
	methods := strings.ToUpper(c.Query("methods"))
	if id == "" || path == "" || methods == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	ok, err := uility.E.RemovePolicy(id, path, methods)
	if err != nil || !ok {
		panic(uility.ErrorMessage{
			ErrorDescription: "删除资源失败!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除资源成功！",
	})
}

// AddAssets 添加用户的资源
func AddAssets(c *gin.Context) {
	path := c.Query("path")
	methods := strings.ToUpper(c.Query("methods"))
	if path == "" || methods == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空！",
		})
		return
	}
	//判断新增资源是否已经存在
	ok := uility.E.HasPolicy(uility.Admin, path, methods)
	if ok {
		panic(uility.ErrorMessage{
			ErrorDescription: "资源已经存在!",
		})

	}
	fmt.Println(ok)

	ok, err := uility.E.AddPolicy(uility.Admin, path, methods)
	if err != nil || !ok {
		panic(uility.ErrorMessage{
			ErrorDescription: "新增资源失败!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "新增资源成功！",
		"code": 200,
	})
}
