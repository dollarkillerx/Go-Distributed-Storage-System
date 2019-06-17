/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:16
* */
package container

import (
	"Go-Distributed-Storage-System/dbops/dao"
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
	"strconv"
)

// 上传文件
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file, header, e := r.FormFile("file")

	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	dir := "./file"
	e = utils.DirPing(dir)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}

	// 创建一个用户接受的文件
	s, e := utils.FileGetPostfix(header.Filename) // 获取文件后缀
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	filename := utils.FileGetRandomName(s)
	filepath := dir + "/" + filename // 生成新的文件名

	log.Println(filepath)

	newFile, e := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		log.Println("open file err")
		log.Println(e.Error())
		return
	}
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(newFile)
	bytes := make([]byte, 1024)
	for {
		_, e := reader.Read(bytes)
		if e != io.EOF {
			writer.Write(bytes)
		} else if e == io.EOF {
			break
		} else if e != nil {
			log.Println("write file err")
			log.Println(e.Error())
			response.RespMsg(w, defs.ErrorBadRequest)
			break
		}
	}

	i, _ := strconv.Atoi(utils.TimeGetNowTimeStr())
	writer.Flush()
	//newFile.Seek(0,0)
	newFile.Close()
	open, e := os.Open(filepath)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	sha1 := utils.FileGetSha1(open)
	defer open.Close()

	meta := &defs.FileMeta{
		FileName: header.Filename,
		Location: filepath,
		FileSize: header.Size,
		UploadAt: i,
		FileSha1: sha1,
	}
	//defs.UpdateFileMeta(meta)
	// 写唯一文件表
	e = dao.UpdateFileMetaToDb(meta)
	if e != nil {
		response.RespInputMsg(w, 400, e.Error())
		return
	}
	// 写用户文件表
	//从token里面获取username
	r.ParseForm()
	username := r.PostForm.Get("username")
	e = dao.OnUserFileUploadFinished(username, meta.FileSha1, meta.FileName, meta.FileSize)
	if e != nil {
		response.RespInputMsg(w, 400, e.Error())
		return
	}

	response.RespMsg(w, &defs.Message{
		Code: 200,
		Resp: &defs.Resp{
			Code:    "001",
			Message: "upload Ok",
			Data: defs.Ic{
				"hash": sha1,
			},
		},
	})
	defer func() {
		file.Close()
	}()

}

// 上传文件view展示页面
func UploadHandlerView(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bytes, e := ioutil.ReadFile("./static/view/file/upload.html")
	if e != nil {
		response.RespMsg(w, defs.ErrorBadView)
		return
	}
	response.RespView(w, bytes)
}

// 获取上传文件信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("filehash")
	if name == "" {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	//meta, e := defs.GetFileMeta(name)
	meta, e := dao.GetFileMetaDb(name)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	response.RespMsg(w, &defs.Message{
		Code: 200,
		Resp: &defs.Resp{
			Code:    "001",
			Message: "get Ok",
			Data: defs.Ic{
				"data": meta,
			},
		},
	})
}

// 批量获取上传文件信息
func FileQueryHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	limit := p.ByName("limit")
	i, e := strconv.Atoi(limit)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	metas := defs.GetLastFileMetas(i)
	response.RespMsg(w, &defs.Message{
		Code: 200,
		Resp: &defs.Resp{
			Code:    "001",
			Message: "get Ok",
			Data: defs.Ic{
				"data": metas,
			},
		},
	})
}

func DownloadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	hash := p.ByName("filehash")
	if hash == "" {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}

	meta, e := defs.GetFileMeta(hash)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	file, e := os.Open(meta.Location)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	bytes := make([]byte, 1024)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\""+meta.FileName+"\"")
	for {
		_, e := reader.Read(bytes)
		if e == io.EOF {
			break
		} else if e != nil {
			response.RespMsg(w, defs.ErrorBadServer)
			return
		} else {
			w.Write(bytes)
		}
	}
}

// 更新元信息接口(rename)
// 注意以下修改只是修改用户认知的名称,而不是真实存储的名称
func FileRename(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	opType := r.PostForm.Get("op")         // 文件类型
	fileHash := r.PostForm.Get("filehash") // 文件hash
	newName := r.PostForm.Get("newname")   // 文件名称

	if opType == "" || fileHash == "" || newName == "" {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}

	meta, e := defs.GetFileMeta(fileHash)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	meta.FileName = newName
	defs.UpdateFileMeta(meta)
	response.RespMsg(w, &defs.Message{
		Code: 200,
		Resp: &defs.Resp{
			Code:    "001",
			Message: "update Ok",
			Data: defs.Ic{
				"data": meta,
			},
		},
	})
}

// 文件删除
func FileDeleteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	hash := p.ByName("filehash")
	if hash == "" {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	err := defs.FileDeleteHandler(hash)
	if err != nil {
		response.RespMsg(w, defs.ErrorBadInternalStorage)
		return
	}
	response.RespInputMsg(w, 200, "del ok!")
}

func TryFastUploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 1.解析请求参数
	r.ParseForm()
	username := r.PostForm.Get("username")
	fileHash := r.PostForm.Get("filehash")
	fileName := r.PostForm.Get("filename")
	fileSize := r.PostForm.Get("filesize")
	// 2.从文件表中查询相同hash的文件记录
	// 3.查不到记录则返回秒传失败
	meta, e := dao.GetFileMetaDb(fileHash)
	if e != nil || meta == nil {
		response.RespInputMsg(w, 400, "秒传失败 文件不存在")
		return
	}
	// 4.上传城关则将文件信息写入到用户文件表中去
	i, e := strconv.ParseInt(fileSize, 10, 64)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadRequest)
		return
	}
	e = dao.OnUserFileUploadFinished(username, fileHash, fileName, i)
	if e != nil {
		response.RespMsg(w, defs.ErrorBadQueryDatabase)
		return
	}
	response.RespInputMsg(w, 200, "秒传成功")
}
