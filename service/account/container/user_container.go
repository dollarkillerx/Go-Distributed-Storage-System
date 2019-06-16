/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午10:00
* */
package container

import (
	"Go-Distributed-Storage-System/config"
	"Go-Distributed-Storage-System/dbops/dao"
	"Go-Distributed-Storage-System/service/account/proto"
	"Go-Distributed-Storage-System/utils"
	"context"
	"log"
)

type User struct {}

// 处理用户主持请求
func (u *User) Signup(ctx context.Context,req *proto.ReqSignup,res *proto.RespSignup) error {
	username := req.Username
	password := req.Password

	// 参数校验
	if len(username) <3 || len(password) <5 {
		res.Code = 400
		res.Message = "注册参数无效"
		return nil
	}
	// 对秘密进行加密
	pwd := utils.Sha1Encode(password + config.ConfigBase.PwdSalt)
	if err := dao.UserSigrup(username, pwd);err == nil {
		res.Code = 200
		res.Message = "注册成功"
		return nil
	}else{
		log.Println(err.Error())
	}
	res.Code = 500
	res.Message = "注册失败"
	return nil

}