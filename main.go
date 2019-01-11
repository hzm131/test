package main

import (
	"fmt"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"html/template"
	"net/http"
)

func main(){
	mux := http.NewServeMux() //创建多路复用器
	files := http.FileServer(http.Dir("/public")) //创建静态文件的处理器
	//将静态文件传递给多路复用处理器函数，并且使用StripPrefix方法移除请求中URL指定的前缀
	mux.Handle("/static/",http.StripPrefix("/static/",files))
	mux.HandleFunc("/",index) //重定向到处理器
	server := &http.Server{
		Addr : "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

//data空接口可以接受任何类型的参数
//...开头表示可以就收零个或任意多个值作为参数，这里我们可以同时传递任意多个模板给该函数
//注意：在go语言中，可变参数必须是可变函数参数的最后一个参数
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

//处理器实际上就是一个接受ResponseWriter和Request指针作为参数的go函数
/*func index(w http.ResponseWriter,r *http.Request){
	threads,err := data.Threads();if err == nil {
		//通过调用session函数可以获取一个存储了用户信息的Session结构，不过因为index函数目前并不需要这些信息，所以用
		//空白标识符忽略了这一结构
		//重点是err变量，程序会根据这个变量来判断用户是否已经登录，然后以此来选择是使用public还是private导航条
		_,err := session(w,r)
		//将模板文件放到了go的切片中，这些html文件都包含了特定的指令 {{}}
		public_tmpl_files := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		private_tmpl_files := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}
		var templates *template.Template
		if err != nil{
			//程序会调用ParseFiles函数对这些模板文件进行语法分析，并创建出相应的模板
			//为了捕捉语法分析过程中可能会产生的错误，程序使用了Must函数去包围ParseFiles函数的执行结果，这样当
			//ParseFiles函数返回错误的时候Must函数就会向用户返回相应的错误报告
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		}else{
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}

			//程序通过调用ExecuteTemplate函数，执行已经经过语法分析的layout模板，执行模板意味着把模板文件中的内容
			//和其他渠道的数据进行合并，然后生成最终的html内容

		templates.ExecuteTemplate(w,"layout",threads)
	}
}*/

func index(w http.ResponseWriter,r *http.Request){
	fmt.Println("连接啦！....")
	threads,err := data.Threads()
	if err == nil{
		_,err := session(w,r)
		if err != nil{
			fmt.Println("没有session")
			generateHTML(w,threads,"layout","public.navbar","index")
		}else{
			generateHTML(w,threads,"layout","private.navbar","index")
		}
	}else{
		generateHTML(w,threads,"layout","public.navbar","index")
	}
}