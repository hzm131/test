package data

import (
	"database/sql"
	"log"
)

/*
	在拥有了表之后，程序就必须考虑如何与数据库进行连接以及如何对表进行操作，
为此我们创建一个名为Db的全局变量，这个全局变量是一个指针，指向的是代表数据库连接池的sql.DB,
后续的代码则会使用Db变量来执行数据库查询操作
*/
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
