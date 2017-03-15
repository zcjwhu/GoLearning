package main

import "fmt"
/**
循环语句for的使用
-go只有for一个循环语句，但支持3中形式
-初始化和步进表达式可以时多个值
-条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数，尽量提前计算好或以变量和常量来代替
-左大括号必须和条件语句在同一行
 */

func main(){
	//使用for来实现其它语言中的其它循环语句
	//第一种形式：没有条件语句,但在语句快中需要终止条件条件
	a:=1
	for{
	   	a++
		if a>5{
			break
		}
		fmt.Println(a)
	}
	//第二种形式：有条件语句
	b:=1
	for b<=3{
		fmt.Println(b)
		b++
	}
	//第三种形式：for的经典语句
	for i:=1;i<3;i++{
		fmt.Println(i)
	}
}
