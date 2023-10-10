package middleware

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Timeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		c.Next()
		fmt.Println(c.Request.Method, c.FullPath(), time.Now().Sub(t1).Milliseconds(), c.Writer.Status())
		go func() {
			models.InsertLog(&models.LogBasic{
				Identity:      uility.GetUuid(),
				Methods:       c.Request.Method,
				Path:          c.FullPath(),
				UserIdentity:  c.MustGet("identity").(string),
				StatusCode:    c.Writer.Status(),
				RequestTime:   time.Now(),
				TimeConsuming: time.Duration(time.Now().Sub(t1).Milliseconds()),
				Role:          c.MustGet("RuleId").(string),
			})
		}()
	}
}
