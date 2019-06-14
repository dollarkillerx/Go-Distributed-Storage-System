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
	"github.com/julienschmidt/httprouter"
	"math"
	"net/http"
	"strconv"
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
	redisConn.Do("HSET","MP_" + info.UploadID,"chunkcount",)
	// 5.将相应信息初始化数据返回到客户端
}
