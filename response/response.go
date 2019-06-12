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

func RespView(w http.ResponseWriter,bytes []byte)  {
	w.WriteHeader(200)
	w.Header().Set("Content-Type","text/html")
	w.Write(bytes)
}

func RespMsg(w http.ResponseWriter,msg *defs.Message)  {
	w.WriteHeader(msg.Code)
	w.Header().Set("Content-Type","application/json")
	bytes, err := json.Marshal(msg.Resp)
	if err != nil {
		return
	}
	w.Write(bytes)
}

