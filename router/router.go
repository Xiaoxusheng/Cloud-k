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
	//登录
	r.POST("/v1/user/UserLogin", controller.Login)
	//注册
	r.POST("/v1/user/UserRegister", controller.UserRegister)

	user := r.Group("/v1/user", middleware.ParseToken())
	//用户详情
	user.GET("/UserDetail", controller.UserDetail)
	//退出登录
	user.GET("/logout", controller.Logout)

	//文件
	file := r.Group("/v1/files", middleware.ParseToken())
	file.POST("/fileUpload", controller.UploadFile)
	//同步文件关联
	file.POST("/repositorySave", controller.RepositorySave)
	//获取文件列表
	file.GET("/fileList", controller.FileList)
	//获取文件夹列表
	file.GET("/folderList", controller.FolderList)
	//更新文件名称
	file.PUT("/fileNameUpdate", controller.UpdateFileName)
	//创建文件夹
	file.GET("/folderCreate", controller.CreateFolder)
	//删除文件
	file.DELETE("/fileDelete", controller.DeleteFile)
	//移动文件
	file.PUT("/fileMove", controller.MoveFile)
	//下载文件
	file.GET("/downloadFile", controller.DownloadFile)
	//修改文件夹名称
	file.GET("/updateFolder", controller.UpdateFolder)
	//分片上传
	file.POST("./FragmentUpload", controller.FragmentUpload)

	//资源分享
	ShareBasic := r.Group("/v1/files", middleware.ParseToken())
	//创建分享资源
	ShareBasic.GET("/ShareBasicCreate", controller.ShareBasicCreate)
	//资源详情
	ShareBasic.GET("/ShareBasicDetail", controller.ShareBasicDetail)
	//保存资源
	ShareBasic.POST("/ShareBasicSave", controller.ShareBasicSave)
	//刷新token
	r.GET("/v1/refresh/authorization", middleware.ParseToken(), controller.RefreshToken)

	return r

}
