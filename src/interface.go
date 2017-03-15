package main

import "fmt"
/*
接口interface

-接口是一个或多个方法签名的集合
-只要某个类型拥有该接口的所有方法签名，就算实现该接口，无需显示声明实现了哪一个接口
-接口只有方法声明，没有方法的实现，没有数据字段
-接口可以匿名嵌入其它接口
-将对象赋值给接口时，会发生拷贝，而接口内部的存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
-只有当将接口存储的类型和对象都为nil时，接口才等于nil
-接口调用不会做receiver的自动转换
-接口同样支持匿名字段方法
-接口可以实现oop中的多态
-空接口可以左任何类型数据的容器
 */

type empty interface {

}

type USB interface {
	Name() string
	Connector //嵌入其它接口
}

type Connector interface {
	Connect()
}

type PhoneConnetor struct {
	name string
}

func (pc PhoneConnetor) Name() string{
	return pc.name
}

func (pc PhoneConnetor) Connect() {
	fmt.Println("connect ...")
}

func main()  {
	a:=PhoneConnetor{name:"zcj"}
	b:=a.Name()
	a.Connect()
	fmt.Println(b)

	c:=Connector(a)
	c.Connect()

	Disconnect(a)

}

func Disconnect(usb interface{})  {
	switch v:=usb.(type) {
	case PhoneConnetor:
		fmt.Println("disconnect",v.name)
	default:
		fmt.Println("unknown device")
	}
}
