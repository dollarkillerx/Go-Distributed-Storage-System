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
	RedisHost string `json:"redisHost"`
}

type OssCon struct {
	OssBucket string `json:"oss_bucket"`
	OssEndpoint string `json:"oss_endpoint"`
	// oss访问key
	OssAccessKeyID string `json:"oss_access_key_id"`
	// oss访问key secret
	OssAccessKeySecret string `json:"oss_access_key_secret"`
}


var (
	ConfigBase *configStruct
	OssConfig *OssCon
)

func init()  {
	getBase()
	getOss()
}

func getBase()  {
	path := "./config.json"
	file, e := os.Open(path)
	err.ErrPanic(e)
	decoder := json.NewDecoder(file)
	ConfigBase = &configStruct{}
	e = decoder.Decode(ConfigBase)
	err.ErrPanic(e)
}

func getOss() {
	path := "./ossconf.json"
	file, e := os.Open(path)
	err.ErrPanic(e)
	decoder := json.NewDecoder(file)
	OssConfig = &OssCon{}
	e = decoder.Decode(OssConfig)
	err.ErrPanic(e)
}