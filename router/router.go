package router

import (
	"Cloud-k/controller"
	"Cloud-k/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Core())
	r.Use(middleware.ParseTakeToken())
	//r.Use(gin.Recovery())
	r.Use(middleware.Error())

	//v 0.1 版本
	r.POST("/v1/user/UserLogin", controller.Login)
	r.POST("/v1/user/UserRegister", controller.UserRegister)

	user := r.Group("/v1/user", middleware.ParseToken())
	user.GET("/UserDetail", controller.UserDetail)
	user.GET("/logout", controller.Logout)

	//文件
	file := r.Group("/v1/files", middleware.ParseToken())
	file.POST("/fileUpload", controller.UploadFile)
	file.POST("/repositorySave", controller.RepositorySave)
	file.GET("/fileList", controller.FileList)
	file.GET("/folderList", controller.FolderList)
	file.PUT("/fileNameUpdate", controller.UpdateFileName)
	file.GET("/folderCreate", controller.CreateFolder)
	file.DELETE("/fileDelete", controller.DeleteFile)
	file.PUT("/fileMove", controller.MoveFile)
	file.GET("/downloadFile", controller.DownloadFile)

	//资源分享
	ShareBasic := r.Group("/v1/files", middleware.ParseToken())
	ShareBasic.GET("/ShareBasicCreate", controller.ShareBasicCreate)
	ShareBasic.GET("/ShareBasicDetail", controller.ShareBasicDetail)
	ShareBasic.POST("/ShareBasicSave", controller.ShareBasicSave)

	r.GET("/v1/refresh/authorization", middleware.ParseToken(), controller.RefreshToken)

	return r

}
