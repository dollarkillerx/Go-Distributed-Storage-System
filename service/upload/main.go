/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:06
* */
package upload

import (
	"Go-Distributed-Storage-System/config"
	"Go-Distributed-Storage-System/router"
	"fmt"
	"net/http"
)

func main() {
	app := router.RegisterRouter()

	fmt.Println("server is run http://127.0.0.1" + config.ConfigBase.Host)
	if err := http.ListenAndServe(config.ConfigBase.Host, app);err != nil {
		panic(err.Error())
	}
}
