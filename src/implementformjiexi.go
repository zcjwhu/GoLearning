package main

import (
	"net/http"
	"fmt"
	"html/template"
)

const tpl=`
<html>
   <head>
      <title>Hey</title>
   </head>
   <body>
       <form method="post" action="/">
            username:<input type="text" name="uname">
            password:<input type="password" name="pwd">
            <button type="submit">Submit</button>
       </form>
   </body>
</html
`

func main(){
	http.HandleFunc("/",Hey)
	http.ListenAndServe(":9001",nil)
}

func Hey(w http.ResponseWriter,r *http.Request)  {
	if r.Method =="GET"{
		t:=template.New("Hey")
		t.Parse(tpl)
		t.Execute(w,nil)

	}else{
		fmt.Println(r.FormValue("uname"))
	}
}
