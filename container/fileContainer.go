/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:16
* */
package container

import (
	"Go-Distributed-Storage-System/defs"
	"Go-Distributed-Storage-System/response"
	"Go-Distributed-Storage-System/utils"
	"bufio"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	file, header, e := r.FormFile("file")

	if e != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	dir := "./file"
	e = utils.DirPing(dir)
	if e != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}

	// 创建一个用户接受的文件
	s, e := utils.FileGetPostfix(header.Filename) // 获取文件后缀
	if e != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	filename := dir +"/" + utils.FileGetRandomName(s) // 生成新的文件名

	log.Println(filename)

	newFile, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if e != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		log.Println("open file err")
		log.Println(e.Error())
		return
	}
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(newFile)
	bytes := make([]byte, 1024)
	for {
		_, e := reader.Read(bytes)
		if e != io.EOF{
			writer.Write(bytes)
		}else if e == io.EOF{
			break
		}else if e != nil{
			log.Println("write file err")
			log.Println(e.Error())
			response.RespMsg(w,defs.ErrorBadRequest)
			break
		}
	}

	response.RespMsg(w,&defs.Message{
		Code:200,
		Resp:&defs.Resp{
			Code:"001",
			Message:"upload Ok",
		},
	})
	defer func() {
		writer.Flush()
		file.Close()
		newFile.Close()
	}()

}

func UploadHandlerView(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	bytes, e := ioutil.ReadFile("./static/view/file/upload.html")
	if e != nil {
		response.RespMsg(w,defs.ErrorBadView)
		return
	}
	response.RespView(w,bytes)
}

