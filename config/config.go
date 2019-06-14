/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:09
* */
package config

import (
	"Go-Distributed-Storage-System/err"
	"encoding/json"
	"os"
)

type configStruct struct {
	Host string `json:"host"`
	MySQLDsn string `json:"mysql_dsn"`
	PwdSalt string `json:"pwd_salt"`
}

var (
	ConfigBase *configStruct
)

func init()  {
	path := "./config.json"
	file, e := os.Open(path)
	err.ErrPanic(e)
	decoder := json.NewDecoder(file)
	ConfigBase = &configStruct{}
	e = decoder.Decode(ConfigBase)
	err.ErrPanic(e)
}
