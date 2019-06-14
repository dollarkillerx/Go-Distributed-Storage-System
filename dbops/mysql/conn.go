/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-12
* Time: 下午12:32
* */
package mysql

import (
	"Go-Distributed-Storage-System/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Engine *sql.DB
	e error
)

func init()  {
	Engine, e = sql.Open("mysql", config.ConfigBase.MySQLDsn)
	if e != nil {
		panic(e.Error())
	}
	Engine.SetMaxOpenConns(1000) // 设置数据库最大连接数
	err := Engine.Ping()
	if err != nil {
		panic(err.Error())
	}
}
