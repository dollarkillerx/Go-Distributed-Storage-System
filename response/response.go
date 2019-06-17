/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:37
* */
package response

import (
	"Go-Distributed-Storage-System/defs"
	"encoding/json"
	"net/http"
)

func RespView(w http.ResponseWriter, bytes []byte) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html")
	w.Write(bytes)
}

func RespMsg(w http.ResponseWriter, msg *defs.Message) {
	w.WriteHeader(msg.Code)
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(msg.Resp)
	if err != nil {
		return
	}
	w.Write(bytes)
}

func RespInputMsg(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	data := &defs.Resp{
		Message: msg,
	}
	bytes, e := json.Marshal(data)
	if e != nil {
		RespMsg(w, defs.ErrorBadInternalInput)
		return
	}
	w.Write(bytes)
}

func RespInputData(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	datas := &defs.Resp{
		Code: "200",
		Data: data,
	}
	bytes, e := json.Marshal(datas)
	if e != nil {
		RespMsg(w, defs.ErrorBadInternalInput)
		return
	}
	w.Write(bytes)
}
