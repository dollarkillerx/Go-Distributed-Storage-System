/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午6:44
* */
package dao

import (
	"Go-Distributed-Storage-System/dbops/mysql"
	"Go-Distributed-Storage-System/utils"
	"errors"
)

type UserFile struct {
	UserName   string `json:"username"`
	FileSha1   string `json:"file_sha1"`
	FileSize   string `json:"file_size"`
	FileName   string `json:"file_name"`
	UploadAt   int64  `json:"upload_at"`
	LastUpdate int64  `json:"last_update"`
}

// 更新用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) error {
	sql := "INSERT IGNORE INTO tbl_user_file(`user_name`,`file_sha1`,`file_name`,`file_size`,`upload_at`,`status`) VALUE(?,?,?,?,?,0)"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		return e
	}
	defer stmt.Close()
	result, e := stmt.Exec(username, filehash, filename, filesize, utils.TimeGetNowTimeStr())
	if e != nil {
		return e
	}
	if i, e := result.RowsAffected(); e != nil {
		return e
	} else {
		if i > 0 {
			return nil
		}
	}
	return errors.New("error")
}

func QueryUserFileMetas(username string, limit int) ([]*UserFile, error) {
	sql := "SELECT file_sha1,file_name,file_size,upload_at form tbl_user_file where user_name = ? limit = ?"
	stmt, e := mysql.Engine.Prepare(sql)
	if e != nil {
		return nil, e
	}
	rows, e := stmt.Query(username, limit)
	if e != nil {
		return nil, e
	}

	files := []*UserFile{}
	for rows.Next() {
		file := &UserFile{}
		e := rows.Scan(&file.FileSha1, &file.FileName, &file.FileSize, &file.UploadAt)
		if e != nil {
			files = append(files, file)
		}
	}
	return files, nil
}
