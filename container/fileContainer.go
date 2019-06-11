/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:16
* */
package container

import (
	"Go-Distributed-Storage-System/err"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	
}

func UploadHandlerView(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	bytes, e := ioutil.ReadFile("./static/view/file/upload.html")
	err.ErrPanic(e)
	w.Write(bytes)
}
