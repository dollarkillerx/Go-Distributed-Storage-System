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
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/file/upload",container.UploadHandlerView)
	router.POST("/file/upload",container.UploadHandler)
	//router.GET("/file/query")
	//router.GET("/file/:filename/download")
	//router.DELETE("/file/:filename/delete")
	//router.PUT("/file/:filename/update")

	return router
}