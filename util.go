package main

import (
	"errors"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"
)

//session的工具函数，在各个处理器函数中复用它
func session(w http.ResponseWriter,r *http.Request) (sess data.Session,err error){
	cookie,err := r.Cookie("_cookie") //从请求头中取出cookie
	if err == nil {
		//访问数据库并核实会话的唯一ID是否存在
		sess = data.Session{ //取出会话
			Uuid: cookie.Value,
		}
		//检查是否存在
		if ok,_:=sess.Check(); !ok{
			err = errors.New("Invalid session") //无效的会话
		}
	}
	return
}
