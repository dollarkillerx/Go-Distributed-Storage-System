/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-12
* Time: 下午12:32
* */
package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Enging *sql.DB
)

func init()  {
	//db, e := sql.Open()
	//db.SetConnMaxLifetime()
}
