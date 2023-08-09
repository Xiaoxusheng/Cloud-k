package router

import (
	"Cloud-k/controller"
	"Cloud-k/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Core())
	//r.Use(gin.Recovery())
	r.Use(middleware.Error())

	//v 0.1 版本
	r.POST("/user/v1/UserLogin", controller.Login)
	r.POST("/user/v1/UserRegister", controller.UserRegister)

	user := r.Group("/user/v1", middleware.ParseToken())
	user.GET("/UserDetail", controller.UserDetail)

	//文件
	file := r.Group("/files/v1", middleware.ParseToken())
	file.GET("/")

	return r

}
