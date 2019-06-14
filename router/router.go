/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:07
* */
package router

import (
	"Go-Distributed-Storage-System/container"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/file/upload",container.UploadHandlerView)
	router.POST("/file/upload",container.UploadHandler)
	router.GET("/file/query/:filehash",container.GetFileMetaHandler)
	router.GET("/file/querylimit/:limit",container.FileQueryHandler)
	router.GET("/file/download/:filehash",container.DownloadHandler)
	router.POST("/file/change",container.FileRename)
	router.DELETE("/file/delete/:filehash",container.FileDeleteHandler)

	router.GET("/user/signup",container.SignupHandlerView)
	router.POST("/user/signup",container.SignupHandler)
	router.POST("/user/signin",container.SignlnHandler)

	router.ServeFiles("/admin/*filepath",http.Dir("./static/view/file"))

	return router
}