/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:07
* */
package router

import (
	"Go-Distributed-Storage-System/container"
	"Go-Distributed-Storage-System/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/file/upload", container.UploadHandlerView)
	router.POST("/file/upload", container.UploadHandler)
	router.GET("/file/query/:filehash", container.GetFileMetaHandler)
	router.GET("/file/querylimit/:limit", container.FileQueryHandler)
	router.GET("/file/download/:filehash", container.DownloadHandler)
	router.POST("/file/change", container.FileRename)
	router.DELETE("/file/delete/:filehash", container.FileDeleteHandler)
	router.POST("/file/fastpload", middleware.CheckToken(container.UserInfoHandler))

	router.GET("/user/signup", container.SignupHandlerView)
	router.POST("/user/signup", container.SignupHandler)
	router.POST("/user/signin", container.SignlnHandler)

	// 分块上传通用接口
	// 初始化分块信息
	router.POST("/file/mpupload/init", container.InitialMultipartUploadHandler)
	// 上传分块
	router.POST("/file/mpupload/uppart", container.UploadPartHandler)
	// 通知分块上传完成
	router.POST("/file/mpupload/complete", container.CompleteUploadHandler)
	// 取消上传分块
	router.POST("/file/mpupload/cancel", nil)
	// 查看分块上传的整体状态
	router.POST("/file/mpupload/status", nil)

	router.ServeFiles("/admin/*filepath", http.Dir("./static/view/file"))

	return router
}
