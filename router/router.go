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
	//登录
	r.POST("/v1/user/userLogin", controller.Login)
	//注册
	r.POST("/v1/user/userRegister", controller.UserRegister)
	r.Use(middleware.Timeout())

	user := r.Group("/v1/user", middleware.Timeout(), middleware.ParseToken())
	//用户详情
	user.GET("/userDetail", controller.UserDetail)
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
	//分片上传初始化
	file.POST("/uploadPart", controller.UploadPart)
	//分片上传
	file.POST("/fragmentUpload", controller.FragmentUpload)
	//分片上传结束
	file.POST("/uploadsCompletion", controller.UploadsCompletion)

	//资源分享
	ShareBasic := r.Group("/v1/files", middleware.Timeout(), middleware.ParseToken())
	//创建分享资源
	ShareBasic.GET("/shareBasicCreate", controller.ShareBasicCreate)
	//资源详情
	ShareBasic.GET("/shareBasicDetail", controller.ShareBasicDetail)
	//保存资源
	ShareBasic.POST("/shareBasicSave", controller.ShareBasicSave)
	//刷新token
	r.GET("/v1/refresh/authorization", middleware.ParseToken(), controller.RefreshToken)
	//管理

	Admin := r.Group("/v1/admin", middleware.Timeout(), middleware.ParseToken(), middleware.Casbin())

	//管理员封禁用户
	Admin.POST("/banned", controller.Banned)
	//解封
	Admin.POST("/unseal", controller.Unseal)
	//查看日志
	Admin.GET("/getLogList", controller.GetLogList)
	//分配容量
	Admin.GET("/divideCapacity", controller.DivideCapacity)
	//管理员所有查看剩余容量
	Admin.GET("/getResidualCapacity", controller.GetResidualCapacity)
	// 查看用户访问信息
	Admin.GET("/getUserLogs", controller.GetUserLogs)

	//超级管理员
	root := r.Group("/v1/root", middleware.Timeout(), middleware.ParseToken(), middleware.Casbin())
	//新增管理员
	root.POST("/add", controller.AddPermission)
	//修改管理员的权限
	root.POST("/update", controller.UpdatePermission)
	//新增资源
	root.POST("/updateAssets", controller.UpdateAssets)
	//获取资源列表
	root.POST("/getAssetsList", controller.GetAssetsList)
	//删除资源
	root.POST("/deleteAssets", controller.DeleteAssets)
	//新增资源
	root.POST("/addAssets", controller.AddAssets)
	return r
}
