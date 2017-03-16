package main

import (
	"fmt"
	"runtime"
)
//import "time"
/*
并发concurrency

-高并发特性，从源码解析，实质上goroutine只是由官方实现的超级线程池
-goroutine奉行通过通信来共享内存，而不是通过共享内存来通信

Channel

-Channel是goroutine沟通的桥梁，大都时阻塞同步的
-通过make来创建，close关闭
-Channel时引用类型
-可以通过for range来迭代不断操作channel
-可以设置单向或双向通道
-可以设置缓存大小，在没有被填满之前不会发生阻塞

Select

-可以处理一个或多个channel的发送和接收
-同时有多个channel按照随机顺序处理
-可用空的select来阻塞main函数
-可设置超时

 */
//func main(){
//	//启动 goroutine
//	go Go()
//	time.Sleep(2*time.Second)
//}
//
//func Go(){
//	fmt.Println("go go go")
//}

/*
简单使用channel通信
 */
//func main(){
//	//创建channel
//	c:=make(chan bool)
//	go func(){
//		fmt.Println("GO!!")
//		c<-true
//	}()
//	<-c
//}
/**
使用for range
有缓存时，异步的
无缓存时，同步的
 */
//func  main()  {
//	c:=make(chan bool)
//	go func(){
//		fmt.Println("GO!!")
//		c<-true
//		close(c)
//	}()
//
//	for v:=range c{
//		fmt.Println(v)
//	}
//}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	c:=make(chan bool,10)
	for i:=0;i<10;i++{
		go Go(c,i)

	}

	//<-c

	for i:=0;i<10;i++{
		<-c
	}
}

func Go(c chan bool ,index int)  {
	a:=1
	for i:=0;i<10000000;i++{
		a+=i
	}
	fmt.Println(index,a)

		c <-true
}
