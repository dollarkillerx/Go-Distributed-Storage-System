/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午5:26
* */
package middleware

import (
	"Go-Distributed-Storage-System/dbops/dao"
	"Go-Distributed-Storage-System/defs"
	"Go-Distributed-Storage-System/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CheckToken(router httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
		token := r.Header.Get("token")
		err := dao.IsValidToken(token)
		if err != nil {
			response.RespMsg(w,defs.ErrorBadRequest)
			return
		}
		router(w,r,p)
	})
}
