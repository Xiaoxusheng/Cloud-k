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

// 分配容量
func DivideCapacity(c *gin.Context) {
	//    查询所有用户
	userList := models.GetUserList()
	//	分配
	list := make([]*models.CapacityBasic, 0)
	for i := 0; i < len(userList); i++ {
		list = append(list, &models.CapacityBasic{
			TotalCapacity:    1,
			ResidualCapacity: 1,
			Recharge:         false,
			Identity:         uility.GetUuid(),
			UserIdentity:     userList[i].Identity,
		})
	}
	models.InsertCapacityBasic(list)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "分配完成！",
	})
}

// 管理员所有查看剩余容量
func GetResidualCapacity(c *gin.Context) {
	list := models.GetResidualCapacityList()
	c.JSON(http.StatusOK, gin.H{
		"cdoe": 200,
		"msg":  "获取成功！",
		"data": gin.H{
			"list": list,
		},
	})
}

//查看用户数据

//查看用户访问信息

// GetLogList 系统操作日志查看
func GetLogList(c *gin.Context) {
	list := models.GetLogBasicList()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取访问日志列表成功！",
		"data": gin.H{
			"log_list": list,
		},
	})
}
