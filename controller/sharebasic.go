package controller

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ShareBasicCreate(c *gin.Context) {
	repository_identity := c.Query("repository_identity")
	user_repository_identity := c.Query("user_repository_identity")
	expired_time := c.DefaultQuery("expired_time", "0")
	if repository_identity == "" || user_repository_identity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	userIdentity := c.MustGet("UserIdentity").(string)
	//查询资源是否存在
	f, _ := models.GetByIdentityUserIdentity(user_repository_identity, userIdentity)
	if !f {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "资源不存在!",
		})
		return
	}
	times, err := strconv.Atoi(expired_time)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorTime:        time.Now(),
			ErrorDescription: err.Error(),
		})
	}
	uuid := uility.GetUuid()
	models.InsertShareBasic(uuid, userIdentity, user_repository_identity, repository_identity, times)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "分享成功!",
		"data": gin.H{
			"identity": uuid,
		},
	})
}

func ShareBasicDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	f := models.GetIdentity(identity)
	if !f {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件不存在!",
		})
		return
	}
	//count++
	models.UpdateClickNum(identity)
	//查询
	ok, data := models.GetShareBasicDetail(identity)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取成功!",
			"data": gin.H{
				"data": data,
			},
		})
	}
}

func ShareBasicSave(c *gin.Context) {

}
