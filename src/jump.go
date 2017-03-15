package main

import "fmt"
/**
跳转语句使用goto ,break,continue的使用
-三个语法都可以配合标签使用
-标签名区分大小写，若不是用会造成编译错误
-break与continue配合标签可以用于多层循环的跳出
-goto是调整程序执行位置，与其它2个语句配合标签的结果并不相同
 */

func main(){
	//配合标签使用的用法，其它用法和其它编程语言类似
	Label:
	for{
		for i:=0;i<10;i++{
			if i>3{
				//外层是一个死循环，使用这种方式和goto不一样，不会继续从Label开始重新执行，而是从面开始执行
				//使用continue只是跳出内层循环
				break Label
				//continue Label
			}
			fmt.Println(i)
		}
	}
	fmt.Println("ok")
}
