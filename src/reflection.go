package main

import (
	"fmt"
	"reflect"
)

/*
反射 reflection

-反射可以大大提升程序的灵活性
-反射使用TypeOf和ValueOf函数从接口中获取目标对象的信息
-反射会将匿名字段作为独立字段
-想要利用反射修改对象状态，前提是interface。data是settable
-通过反射可以动态调用方法
 */

type User struct {
	Id int
	Name string
	Age int
}

func (user User) Hello(name string){
	fmt.Println("hello world",name,"my name is ",user.Name)
}

type Manager struct {
	User
	title string
}

func main(){
	user :=User{ 1,"zcj",3}
	Info(user)

	m:=Manager{User:User{1,"cyh",66},title:"123"}
	t:=reflect.TypeOf(m)

	//选取匿名字段
	fmt.Printf("%#v\n",t.Field(0))
	fmt.Printf("%#v\n",t.FieldByIndex([]int{0,0}))

	//通过反射来设置属性值
	Set(&user)
	fmt.Println(user)

	//通过反射动态调用方法
	v:=reflect.ValueOf(user)
	mv:=v.MethodByName("Hello")

	args:=[]reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)


}

func Info(o interface{}){
	t:=reflect.TypeOf(o)
	fmt.Println(t.Name())

	v:=reflect.ValueOf(o)
	fmt.Println("field:")

	//反射出字段
	for i:=0;i<t.NumField();i++{
		f:=t.Field(i)
		val:=v.Field(i).Interface()
		fmt.Printf("%6s: %v=%v\n",f.Name,f.Type,val)
	}
	//反射出方法
	for i:=0;i<t.NumMethod();i++{
		m:=t.Method(i)
		fmt.Printf("%6s: %v\n",m.Name,m.Type)
	}
}

func Set(o interface{})  {
	v:=reflect.ValueOf(o)

	if v.Kind()==reflect.Ptr&&!v.Elem().CanSet(){
		fmt.Println("xxx")
		return
	}else{
		v=v.Elem()
	}

	f:=v.FieldByName("Name")
	if !f.IsValid(){
		fmt.Println("bad")
		return
	}

	if f.Kind()==reflect.String{
		f.SetString("byebye")
	}
}