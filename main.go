package main

import (
	"Cloud-k/router"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func main() {
	gin.ForceConsoleColor()
	// 记录到文件。
	f, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := router.Router()

	err = r.Run(":80")
	if err != nil {
		panic("服务器启动失败" + err.Error())
	}
}
