package main

import (
	"fmt"
)
/*
函数的使用

-go函数不支持嵌套，重载和默认参数
-但支持如下：
  无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包
-定义函数使用func关键字，左大括号不能另起一行
-函数也可以作为一种类型使用
 */
func main(){
     //匿名函数，将函数作为一种类型
	a:= func() {
		fmt.Println("匿名函数！")
	}
	a()
      b:=closure(10)
	fmt.Println(b(1))
	fmt.Println(b(2))
}

/*
对于简单类型，实际上是值传递，如果需要改变，可以使用指针
对于引用类型，是传递的已经时地址了
 */
func A(){

}

func B(a int ,b int){

}

func C(a,b,c int)(d int,e int){
	return 2,3
}

func e(a,b,c int)(d,e int){
	return d,e
}
/*
不定长变参
需要将它放到最后一个参数，如果有其它类型
 */
func f(a ...int){
	fmt.Println(a)
}

//闭包
func closure(x int) func(int) int{
	return func(y int) int {
		return x+y
	}
}
