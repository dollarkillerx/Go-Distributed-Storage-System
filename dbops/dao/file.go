/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 上午10:20
* */
package dao

import (
	"Go-Distributed-Storage-System/dbops/mysql"
	"Go-Distributed-Storage-System/defs"
	"Go-Distributed-Storage-System/utils"
	"errors"
	"log"
)

// 对file文件相关的dao操作

// 入库
func OnFileUploadFinished(filehash,filename,fileaddr string,filesize int64) error {
	sql := "INSERT IGNORE INTO `tbl_file`(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`,`create_at`) VALUE(?,?,?,?,0,?)"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		log.Println("Prepare error")
		return errors.New(e.Error())
	}
	defer stmt.Close()
	result, e := stmt.Exec(filehash, filename, filesize, fileaddr,utils.TimeGetNowTimeStr())
	if e != nil {
		log.Println("exec error")
		return errors.New(e.Error())
	}

	if i, e := result.RowsAffected();e == nil {
		if i>0 {
			return nil
		}
	}else{
		log.Println(e.Error())
	}
	return errors.New("insert data error")
}

// 查询
func GetFileMetaDb(filehash string) (*defs.FileMeta,error)  {
	sql := "SELECT `file_sha1`,`file_size`,`file_addr`,`create_at`,`file_name` FROM `tbl_file` WHERE `file_sha1` = ? AND `status` = 0 LIMIT 1"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		log.Println("error mysql prepare")
		return nil,e
	}
	defer stmt.Close()
	meta := &defs.FileMeta{}
	e = stmt.QueryRow(filehash).Scan(&meta.FileSha1, &meta.FileSize, &meta.Location, &meta.UploadAt, &meta.FileName) // 查询单挑记录
	if e != nil {
		log.Println("error mysql query row")
		log.Println(e.Error())
		return nil,e
	}
	return meta,nil
}

func UpdateFileMetaToDb(fmeta *defs.FileMeta) error {
	err := OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.Location, fmeta.FileSize)
	if err != nil {
		return err
	}
	return nil
}
