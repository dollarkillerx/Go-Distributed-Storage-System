/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午12:26
* */
package dao

import (
	"Go-Distributed-Storage-System/config"
	"Go-Distributed-Storage-System/dbops/mysql"
	"Go-Distributed-Storage-System/utils"
	"errors"
	"fmt"
	"log"
	"strings"
)

func UserSigrup(username, password string) error {
	sql := "INSERT IGNORE INTO `tbl_user`(`user_name`,`user_pwd`,`signup_at`,`status`) VALUE(?,?,?,0)"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		log.Println(e.Error())
		return e
	}
	defer stmt.Close()

	fmt.Println(username)
	fmt.Println(password)

	result, e := stmt.Exec(username, password, utils.TimeGetNowTimeStr())
	if e != nil {
		log.Println(e.Error())
		return e
	}
	if i, e := result.RowsAffected(); e == nil {
		if i > 0 {
			return nil
		}
	} else {
		log.Println(e.Error())
	}
	return errors.New("sigrup err")
}

func UserSignin(username, password string) error {
	pwd := password + config.ConfigBase.PwdSalt
	encode := utils.Sha1Encode(pwd)
	// 着会有两种 我选择最安全的那种把
	sql := "SELECT `user_pwd` FROM `tbl_user` WHERE `user_name` = ? AND `status` = 0"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		return e
	}
	defer stmt.Close()
	var uspwd string
	e = stmt.QueryRow(username).Scan(&uspwd)
	if e != nil {
		return e
	}
	// true 匹配正确
	if strings.EqualFold(encode, uspwd) {
		return nil
	}
	return errors.New("pwd not")
}
