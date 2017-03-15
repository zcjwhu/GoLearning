package main

import "fmt"
/*
为结构编写方法
 */
//使用嵌入结构，同名字段的读取处理
type A struct {
	B
	name string
}
type B struct {
	name string
}

//为结构写方法
type student struct {
	name string
}

type teacher struct {
	name string
}

func main(){
	//嵌入结构测试
      a:=A{
	      name:"A",
	      B:B{name:"B"},
      }
	fmt.Println(a.name,a.B.name)

	s:=student{name:"zcj"}
	t:=teacher{name:"wxp"}

	s.Print()
	fmt.Println(s.name)
	t.Print()
	fmt.Println(t.name)

	(*student).Print(&s)
}

/*
方法的绑定
对于字段的公有和私有，小写的对于同一包都是可以访问的，对于包外不可访问，所以下面结构的方法可以访问私有字段
和函数的定义类似，只是强制增加一个recieve 部分
 */
//为student写方法,
func (s *student) Print(){
	s.name="hahh"
	fmt.Println("student")
}

//为teacher写方法
func (t teacher) Print(){
	fmt.Println("teacher")
}
