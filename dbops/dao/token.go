/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午4:19
* */
package dao

import (
	"Go-Distributed-Storage-System/dbops/mysql"
	"errors"
	"log"
)

func UpdateToken(username,token string) error {
	sql := "replace into `tbl_user_token`(`user_name`,`user_token`) VALUE(?,?)"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		log.Println(e.Error())
		return e
	}
	defer stmt.Close()
	result, e := stmt.Exec(username, token)
	if e != nil {
		log.Println(e.Error())
		return e
	}

	if i, e := result.RowsAffected();e != nil {
		log.Println(e.Error())
		return e
	}else{
		if i>0{
			return nil
		}
	}
	return errors.New("error")
}

func IsValidToken(token string) error {
	// 判断token的时效性
	// 数据库中查询是否存在此token
	//上面设计token就有点问题fuck,打脸了

	//就只是检测是否存在把
	sql := "select * from `tbl_user_token` WHERE user_token = ?"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		return e
	}
	var tokenn string
	e = stmt.QueryRow(token).Scan(&tokenn)
	if e != nil {
		return e
	}
	return nil
}