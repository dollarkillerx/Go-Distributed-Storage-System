/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 上午10:26
* */
package main

import (
	"Go-Distributed-Storage-System/service/apigw/route"
	"net/http"
)

func main() {
	router := route.RegisterRoute()

	http.ListenAndServe(":8081", router)
}
