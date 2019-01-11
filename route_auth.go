package main

import (
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"fmt"
	"net/http"
	)

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	//通过给定的email获取与之对应的User结构
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		fmt.Println(err, "Cannot find user")
	}
	//data.Encrypt函数用于加密给定的字符串
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		//将设置到的cookie添加到响应头中
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}