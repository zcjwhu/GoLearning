package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

//第一种实现，使用一个handler
type myHandler struct {

}

func (*myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"hello world,handler second version")
}

//第二种实现，利用一个函数代替handler
func sayHello(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"hello world,function second version")
}

func main()  {
	mux:=http.NewServeMux()
        //使用一个mux来控制路由,两种不同方式
	mux.Handle("/",&myHandler{})
	mux.HandleFunc("/sayHello",sayHello)

	wd,err:=os.Getwd()
	if err!=nil{
		log.Fatal(err)
	}

	//处理静态路由,实现静态文件服务器
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(wd))))

	err=http.ListenAndServe(":9000",mux)

	if err!=nil{
		log.Fatal("ListenAndServe",err)
	}

}


