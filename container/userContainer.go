/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午3:22
* */
package container

import (
	"Go-Distributed-Storage-System/config"
	"Go-Distributed-Storage-System/dbops/dao"
	"Go-Distributed-Storage-System/defs"
	"Go-Distributed-Storage-System/response"
	"Go-Distributed-Storage-System/utils"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

// 用户注册页面view
func SignupHandlerView(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	bytes, e := ioutil.ReadFile("./static/view/file/signup.html")
	if e != nil {
		response.RespMsg(w,defs.ErrorBadView)
		return
	}
	response.RespView(w,bytes)
}

// 处理用户注册
func SignupHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	if len(username)<3 || len(password)<5 {
		log.Println("error len min")
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	pwd := password + config.ConfigBase.PwdSalt
	encode := utils.Sha1Encode(pwd)
	err := dao.UserSigrup(username, encode)
	if err != nil {
		log.Println(err.Error())
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	response.RespInputMsg(w,201,"sig OK!")
}

// 登陆接口
func SignlnHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	// 1.验证用户名以及密码
	err := dao.UserSigrup(username, password)
	if err != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	// 2.生成访问凭证(token)
	token := defs.GetToken(username)
	err = dao.UpdateToken(username,token)
	if err != nil {
		response.RespMsg(w,defs.ErrorBadServer)
		return
	}
	// 3.登陆成功后重定向到首页
	response.RespInputData(w,200,defs.Ic{
		"location":"http://" + r.Host + "home.html",
		"username":username,
		"token":token,
	})
}

func UserInfoHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	// 1.解析请求参数
	r.ParseForm()
	//username := r.Form.Get("username")
	token := r.Header.Get("token")
	// 2.验证token是否有效
	err := dao.IsValidToken(token)
	if err != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	// 3.查询用户信息

	// 4.组装并相应用户数据

	// 这里感觉不完善就重写想想怎么改造  感觉这个系统没有必要这个阿! 返回用户信息干嘛阿!主要作oss云存储阿
}