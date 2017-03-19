package main

import (
	"net/http"
	"io"
)

func main()  {
	http.HandleFunc("/",Cookie)
	http.HandleFunc("/2",Cookie2)
	http.ListenAndServe(":9001",nil)
}

func Cookie(w http.ResponseWriter,r *http.Request)  {
	ck:=&http.Cookie{
		Name:"myCookie",
		Value:"hello cookie1",
		Path:"/",
		Domain:"localhost",
		MaxAge:120,

	}

	http.SetCookie(w,ck)

	ck2,err:=r.Cookie("myCookie")
	if err!=nil{
		io.WriteString(w,err.Error())
		return
	}

	io.WriteString(w,ck2.Value)
}

func Cookie2(w http.ResponseWriter,r *http.Request)  {
	ck:=&http.Cookie{
		Name:"myCookie",
		Value:"hello cookie2",
		Path:"/",
		Domain:"localhost",
		MaxAge:120,

	}

	w.Header().Set("Set-Cookie",ck.String())

	ck2,err:=r.Cookie("myCookie")
	if err!=nil{
		io.WriteString(w,err.Error())
		return
	}

	io.WriteString(w,ck2.Value)

}