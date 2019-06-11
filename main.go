/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:06
* */
package main

import (
	"Go-Distributed-Storage-System/config"
	"Go-Distributed-Storage-System/router"
	"net/http"
)

func main() {
	app := router.RegisterRouter()

	if err := http.ListenAndServe(config.ConfigBase.Host, app);err != nil {
		panic(err.Error())
	}
}
