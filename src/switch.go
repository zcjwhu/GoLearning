package main

import "fmt"
/**
选择语句使用switch的使用
-可以使用任何类型或表达式作为条件语句
-不需要写break,一旦条件符合自动终止
-如果希望继续执行下一个case,需要使用fallthrough语句
-支持一个初始化表达式（可以时并行方式），右侧用分号
-左大括号必须与条件语句在同一行
 */

func main(){
	a:=1
	switch a{
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	default:
		fmt.Println("none")
	}


	switch b:=3;{
	case b<0:
		fmt.Println("b<0")
	case b>=0:
		fmt.Println("b>=0")
		fallthrough
	case b>=1:
		fmt.Println("b>=1")
	default:
		fmt.Println("none")
	}
}
