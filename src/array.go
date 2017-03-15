package main

import "fmt"
/*
数组array
-定义数组的格式：var varname [n]type n>=0
-数组的长度也是类型的一部分，因此具有不同长度的数组为不同的类型
-区分指向数组的指针和指针数组
-数组在go语言中为值类型,在传递时需要区别
-数组之间可以使用==或！=进行比较，但是不可以用<或>
-可以使用new来创建数组，此方法返回一个指向数组的指针
-go支持多为数组
 */

func main(){
	//基本应用
	var a [2]int
	var b [1]int
	c:=[2]int{8}
	d:=[20]int{19:1}//将索引为19的设置为1,其它为默认值
	e:=[...]int{6,3,1}
	//b=a直接赋值时不可以的，因为长度不同属于不同类型
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//指向数组的指针
	m:=[...]int{99:2}
	var pp *[100]int=&m
	fmt.Println(pp)

	//指针数组
	x,y:=5,2
	p:=[...]*int{&x,&y}
	fmt.Println(*p[0])

	//使用new,返回一个指向数组的指针
	pn:=new([10]int)
	fmt.Println(pn[1])

	//多维数组
	ma:=[2][3]int{
		{1,2,3},
		{4,5,6}}
	fmt.Println(ma)

	bubble:=[...]int{5,1,4,8,4,2}
	fmt.Println(bubble)
	length:=len(bubble)
	//冒泡排序
	for i :=length-1;i>=1;i--{
		for j:=0;j<i;j++{
			if bubble[j]>bubble[i] {
				temp:=bubble[j]
				bubble[j]=bubble[i]
				bubble[i]=temp
			}
		}
	}
	fmt.Println(bubble)
}
/*
	输出结果：
	[0 0]
	[0]
	[8 0]
	[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	[6 3 1]
	&[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]
	5
	0
	[[1 2 3] [4 5 6]]
	[5 1 4 8 4 2]
	[1 2 4 4 5 8]
	 */

