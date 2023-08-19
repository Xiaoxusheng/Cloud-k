package controller

import (
	"Cloud-k/uility"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken(c *gin.Context) {
	userIdentity := c.MustGet("UserIdentity").(string)

	token := uility.GetToken(userIdentity, 2)

	refresh_token := uility.GetToken(userIdentity, 24)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功!",
		"data": gin.H{
			"token":         token,
			"refresh_token": refresh_token,
		},
	})

}
