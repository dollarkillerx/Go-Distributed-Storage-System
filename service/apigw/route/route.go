/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 上午10:24
* */
package route

import (
	"Go-Distributed-Storage-System/service/apigw/container"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoute() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user/signup",container.DoSignupHandle)

	return router
}
