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
	r.POST("/v1/user/UserLogin", controller.Login)
	r.POST("/v1/user/UserRegister", controller.UserRegister)

	user := r.Group("/v1/user", middleware.ParseToken())
	user.GET("/UserDetail", controller.UserDetail)

	//文件
	file := r.Group("/v1/files", middleware.ParseToken())
	file.GET("/fileDetail")
	file.POST("/fileupload", controller.UploadFile)

	return r

}
