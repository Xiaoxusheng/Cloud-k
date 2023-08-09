package main

import (
	"Cloud-k/router"
)

func main() {
	r := router.Router()

	err := r.Run(":80")
	if err != nil {
		panic("服务器启动失败" + err.Error())
	}
}
