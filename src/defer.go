package main

import "fmt"
/*
defer
-执行方式类似析构函数，在函数体执行结束之后按照调用的相反顺序执行
-即使出现错误也能执行
-支持匿名函数调用
-常用于资源清理，文件关闭，解锁以及记录时间等操作
-通过与匿名函数配合，可以在return之后修改函数计算结果

-Go没有异常机制，但是有panic/recover模式来处理
-Panic可以在任何地方引发，recover只有在defer调用的函数中有效
 */
func main(){
	//defer  fmt.Println("a")
	//defer  fmt.Println("b")
	//for i:=0;i<5;i++{
	//	defer func() {
	//		fmt.Println(i)
	//	}()
	//}
	/*
	输出结果： 这里输出全是5,和js中的闭包原因一致
	5
	5
	5
	5
	5
	b
	a
	 */
	A()
	B()
	C()
}

func A(){
	fmt.Println("func A")
}

func B()  {
	/*
	下面这段将程序从panic中恢复出来
	 */
	defer func(){
		if err:=recover(); err!=nil{
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")

}

func C(){
	fmt.Println("func C")
}
