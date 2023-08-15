package main

import (
	"Cloud-k/router"
	"Cloud-k/uility"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.ForceConsoleColor()
	// 记录到文件。
	go uility.CreateLogFile()

	r := router.Router()

	err := r.Run(":80")
	if err != nil {
		log.Fatalln("服务器启动失败" + err.Error())
	}
}
