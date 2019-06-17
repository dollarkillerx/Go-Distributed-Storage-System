/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 上午10:09
* */
package container

import (
	"Go-Distributed-Storage-System/defs"
	"Go-Distributed-Storage-System/response"
	"Go-Distributed-Storage-System/service/account/proto"
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro"
	"net/http"
)

var (
	userCli proto.UserService
)

func init() {
	//service := micro.NewService(
	//	micro.Name("go.micro.api.user"),
	//	micro.RegisterTTL(time.Second*10),
	//	micro.RegisterInterval(time.Second*5),
	//)
	service := micro.NewService(
	//这个服务不需要访问就没有必要填写name了
	)
	// 初始化,解析命令行参数
	service.Init()

	// 创建rpcClient
	userCli = proto.NewUserService("go.micro.service", service.Client())
}

func DoSignupHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Header.Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	resp, e := userCli.Signup(context.TODO(), &proto.ReqSignup{
		Username: username,
		Password: password,
	})
	if e != nil {
		response.RespMsg(w, defs.ErrorBadServer)
		return
	}
	response.RespInputMsg(w, int(resp.Code), resp.Message)
}
