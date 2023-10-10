package middleware

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		defer func() {
			go func() {
				role, ok := c.Get("RuleId")
				identity, k := c.Get("identity")
				if ok && k {
					models.InsertLog(&models.LogBasic{
						Identity:      uility.GetUuid(),
						Methods:       c.Request.Method,
						Path:          c.FullPath(),
						UserIdentity:  identity.(string),
						StatusCode:    c.Writer.Status(),
						RequestTime:   time.Now(),
						TimeConsuming: time.Duration(time.Now().Sub(t1).Milliseconds()),
						Role:          role.(string),
					})
				}
			}()
		}()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err", err)
				Err := ""
				errorMessage := err.(uility.ErrorMessage)
				Err = errorMessage.ErrorDescription
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": Err,
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
		ip := uility.GetServerIP()
		uility.SendErrorEmail("3096407768@qq.com", ip, uility.ErrorMessage{ErrorDetails: errorMessage.ErrorDetails, ErrorDescription: errorMessage.ErrorDescription, ErrorType: "严重事故！", ErrorTime: time.Now()})
	}
	if strings.Contains(errorMessage.ErrorType, "400") {
		if uility.Count >= 10 {
			ip := uility.GetServerIP()
			uility.SendErrorEmail("3096407768@qq.com", ip, uility.ErrorMessage{ErrorDetails: errorMessage.ErrorDetails, ErrorDescription: errorMessage.ErrorDescription, ErrorType: "警告！", ErrorTime: time.Now()})
			uility.Count = 0
		}
		uility.Count++
		fmt.Println(uility.Count)
	}

}
