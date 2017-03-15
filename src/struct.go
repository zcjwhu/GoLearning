package main

import "fmt"
/*
结构struct

-go中的struct和c语言中的struct非常相似，并且go没有class
-使用type <Name> struct{} 定义数据结构，名称遵循可见性规则
-支持指向自身类型的指针类型成员
-支持匿名结构，可用作成员或者定义成员变量
-可以使用字面值对结构进行初始化
-允许直接通过指针来读写结构成员
-相同类型的成员可进行直接拷贝赋值
-支持==和！=比较运算符，不支持<和>
-支持匿名字段

 */

/*
go没有继承的概念
对于一些共用的属性，可以通过组合嵌入的方式，如下
 */
type human struct {
	sex int
}

type teacher struct {
	human
	name string
	age int
}
type student struct {
	human
      name string
      age int
}

/*
	匿名字段
	进行初始化需要参数类型一样
	 */
type dog struct {
	string
	int
}

func main(){
	a:=student{}
	a.name="zcj"  //操作struct中的成员使用 .
	a.age=25
	b:=student{
		name:"cyh",
		age:25,
	}
	fmt.Println(a)
	fmt.Println(b)
	//值传递
	A(a)
	fmt.Println(a)
	//引用传递
	/*
	为了实现结构操作和对结构进行操作的方便性，比如函数B的调用就太麻烦了，为了减少取地址符的书写，a:=&student{ ... }
	并且使用指针a来操作属性也只需要通过. 的操作
	 */
	B(&a)
	fmt.Println(a)

	/*
	匿名结构
	 */
	c:= struct {
		color string
	}{
		color:"red",
	}
	fmt.Println(c.color)

	//c:= &struct {
	//	color string
	//}{
	//	color:"red",
	//}
	//fmt.Println(c.color)

	p:=teacher{name:"wxp",age:40,human:human{sex:1}}
	q:=student{name:"zcj",age:25,human:human{sex:1}}
	p.name="wxp1"
	p.age=13
	p.sex=0 //这里可以直接获取sex属性
	fmt.Println(p,q)

}

func A(s student){
	s.age=20
}

func B(s *student){
	s.age=15
}