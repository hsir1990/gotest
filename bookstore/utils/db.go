package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

//初始化数据库连接，获取数据库连接
func init() {
	// ?parseTime=true&loc=Local 设置本地时区
	Db, err = sql.Open("mysql",
		"root:Hsir1990!@tcp(39.103.175.224:3306)/bookstore?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
}
