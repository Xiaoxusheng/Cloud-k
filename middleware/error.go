package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Errors.Last(); err != nil {
			// 在这里你可以将错误记录到日志文件、数据库或通过网络发送
			log.Printf("发生错误: %v", err.Error())

			// 你还可以添加更多的处理逻辑，比如将错误写入数据库等
		}
		c.Next()
	}
}
