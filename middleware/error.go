package middleware

import (
	"Cloud-k/uility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err", err)
				errorMessage := err.(uility.ErrorMessage)
				fmt.Println(errorMessage)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": errorMessage.ErrorDescription,
				})
				// 根据错误级别发送邮件
				determineErrorLevel(errorMessage)
			}
		}()
		c.Next()
	}
}

// 错误处理
func determineErrorLevel(errorMessage uility.ErrorMessage) {
	if strings.Contains(errorMessage.ErrorType, "500") {
		uility.SendErrorEmail("3096407768@qq.com", uility.ErrorMessage{ErrorDetails: errorMessage.ErrorDetails, ErrorDescription: errorMessage.ErrorDescription, ErrorType: "bug", ErrorTime: time.Now()})
	}
	if strings.Contains(errorMessage.ErrorType, "400") {
		if uility.Count >= 10 {
			uility.SendErrorEmail("3096407768@qq.com", uility.ErrorMessage{ErrorDetails: errorMessage.ErrorDetails, ErrorDescription: errorMessage.ErrorDescription, ErrorType: "bug", ErrorTime: time.Now()})

		}
		uility.Count++
	}
}
