package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func main()  {
	http.HandleFunc("/",sayHelloName) //设置访问的路由

	err:=http.ListenAndServe(":9000",nil) //设置监听的端口,这里handler使用了默认的
	if err!=nil{
		log.Fatal("ListenAndServe",err)
	}
}

func sayHelloName(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm() //解析请求参数
	fmt.Println(r.Form) //这些是输出到服务端打印的信息
	fmt.Println("path",r.URL.Path) //打印请求路径
	fmt.Println("scheme",r.URL.Scheme) //打印请求路径
	fmt.Println(r.Form["url_long"])
	for k,v:= range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello zcj,firset version!")
}
