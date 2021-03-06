/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:24
* */
package defs

import "net/http"

var (
	ErrorBadView            = &Message{Code: http.StatusInternalServerError, Resp: &Resp{Message: "Status Internal View Error", Code: "0051"}}
	ErrorBadRequest         = &Message{Code: http.StatusBadRequest, Resp: &Resp{Message: "Status Bad Request", Code: "0040"}}
	ErrorBadServer          = &Message{Code: http.StatusInternalServerError, Resp: &Resp{Message: "Status Internal Server Error", Code: "0050"}}
	ErrorBadQueryDatabase   = &Message{Code: http.StatusInternalServerError, Resp: &Resp{Message: "Database Query Error", Code: "0052"}}
	ErrorBadInternalInput   = &Message{Code: http.StatusInternalServerError, Resp: &Resp{Message: "Bad Internal input error", Code: "0053"}}
	ErrorBadInternalStorage = &Message{Code: http.StatusInternalServerError, Resp: &Resp{Message: "Bad Internal Storage", Code: "0054"}}
)
