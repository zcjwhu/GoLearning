package main

import (
	"io"
	"log"
	"net/http"
	"time"
)


//使用一个map来保存路由映射结构
var mux map[string]func(http.ResponseWriter,*http.Request)

func main()  {
	server:=http.Server{
		Addr:":9000",
		Handler:&myHandler{},
		ReadTimeout:5*time.Second,
	}

	mux=make(map[string]func(http.ResponseWriter,*http.Request))
	mux["/hello"]=sayHello

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal("ListenAndServe",err)
	}

}
type myHandler struct {

}

func (*myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//将路由注册到handlerde Servehttp中
	if h,ok :=mux[r.URL.String()];ok{
		//路由转发。比如/hello 调用sayhello()
		h(w,r)
		return
	}

	io.WriteString(w,"hello world,say handler third version")
}

func sayHello(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"hello world,function sayHello third version")
}


