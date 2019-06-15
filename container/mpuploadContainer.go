/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午9:55
* */
package container

import (
	"Go-Distributed-Storage-System/cache/redis"
	"Go-Distributed-Storage-System/defs"
	"Go-Distributed-Storage-System/response"
	"Go-Distributed-Storage-System/utils"
	"bufio"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 初始化信息
type MultipartUploadInfo struct {
	FileHash string `json:"file_hash"`
	FileSize int64	`json:"file_size"`
	UploadID string	`json:"upload_id"`
	ChunkSize int `json:"chunk_size"`// 每一个分块大小
	ChunkCount int `json:"chunk_count"`// 分块数量
}

// 初始化分块上传
func InitialMultipartUploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	// 1.解析用户请求信息
	r.ParseForm()
	username := r.PostForm.Get("username")
	filehash := r.PostForm.Get("filehash")
	filesize, err := strconv.ParseInt(r.PostForm.Get("filesize"), 10, 64)
	if err != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}

	// 2.获得redis的连接
	redisConn := redis.RedisConn.Get()
	defer redisConn.Close()

	// 3.生成分块上传的初始化信息
	info := &MultipartUploadInfo{
		FileHash:filehash,
		FileSize:filesize,
		UploadID:username + utils.TimeGetNowTimeStr(),
		ChunkSize:5*1024*1024,// 5MD
		ChunkCount:int(math.Ceil(float64(filesize)/(5*1024*1024))),// 转float64除法在向上取整
	}
	// 4.将初始化信息写入到redis缓存
	redisConn.Do("set","name","age")
	redisConn.Do("HSET","MP_" + info.UploadID,"chunkcount",info.ChunkCount)
	redisConn.Do("HSET","MP_" + info.UploadID,"filehash",info.FileHash)
	redisConn.Do("HSET","MP_" + info.UploadID,"filesize",info.FileSize)
	redisConn.Do("HSET","MP_" + info.UploadID,"chunksize",info.ChunkSize)
	// 5.将相应信息初始化数据返回到客户端
	response.RespInputData(w,200,info)
}

// 上传文件分块
func UploadPartHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	// 1.解析用户请求参数
	r.ParseForm()
	//username := r.Form.Get("username")
	uploadId := r.Form.Get("uploadid")
	chunkIndex := r.Form.Get("index")

	// 2.获得redis连接池中的一个连接
	redisConn := redis.RedisConn.Get()
	defer redisConn.Close()

	// 3.获得文件句柄,用于存储分块内容
	path := "data/" + uploadId
	err := utils.DirPing(path)
	if err != nil {
		response.RespMsg(w,defs.ErrorBadServer)
		return
	}
	file := path + "/" + chunkIndex
	openFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 00666)
	if err != nil {
		response.RespMsg(w,defs.ErrorBadServer)
		return
	}
	defer openFile.Close()
	data, _, err := r.FormFile("file")
	defer data.Close()
	if err != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	reader := bufio.NewReader(data)
	writer := bufio.NewWriter(openFile)
	buf := make([]byte, 1024*1024) // 1M buf
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			break
		}else if err != nil {
			response.RespMsg(w,defs.ErrorBadRequest)
			return
		}else{
			writer.Write(buf)
		}
	}
	writer.Flush()
	// 4.更新redis缓存状态
	redisConn.Do("HSET","MP_"+uploadId,"chkidx_"+chunkIndex,1)
	// 5.返回处理结果到客户端
	response.RespInputMsg(w,200,"ok")

	// 不足之处,客户端上传需要携带当前分块的hash,服务端校验确保文件的完整性
}

// 通知上传合并接口
func CompleteUploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	// 1.解析请求参数
	r.ParseForm()
	upid := r.Form.Get("uploadid")
	//username := r.Form.Get("username")
	//filehash := r.Form.Get("filehash")
	//filesize := r.Form.Get("filesize")
	//filename := r.Form.Get("filename")

	// 2.获得redis连接池的一个连接
	redisConn := redis.RedisConn.Get()
	defer redisConn.Close()

	// 3.通过uploadid查询redis判断是否所有分块上传完成
	values, e := redis2.Values(redisConn.Do("HGETALL", "MP_"+upid))
	if e != nil {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}
	totalCount := 0 // 上传完成数量
	chunkCount := 0 // 总数量
	for i:=0;i<len(values);i+=2{
		k := string(values[i].([]byte))
		v := string(values[i+1].([]byte))
		if k == "chunkcount" {
			totalCount,_=strconv.Atoi(v)
		}else if strings.HasPrefix(k,"chkidx_") && v == "1" {
			chunkCount += 1
		}
	}
	// 不等就是上传没有完成
	if totalCount != chunkCount {
		response.RespMsg(w,defs.ErrorBadRequest)
		return
	}

	// 4.合并分块
	// 5.更新唯一文件表,更新用户文件表

	// 6.相应处理结果

}