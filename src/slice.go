package main

import "fmt"

/*
切片slice

-其本身并不是数组，它是指向底层的数组
-作为变长数组的替代方案，可以关联底层数组的局部或全部
-是引用类型
-可以直接创建或从底层数组获得
-使用len()获取元素的个数，cap（）获取容量
-一般使用make()创建创建
-如果多个slice指向相同的底层数组，其中一个值的改变会影响全部

-make([]T,len,cap)
-其中cap可以省略，和len的值相同
-len表示存储元素的个数，cap表示容量
 */
func main(){
	/*
	slice的第一种声明方式
	 */
	var s1 []int  //slice不用指定大小
	fmt.Println(s1)

	a:=[10]int{1,2,3,4,5,6,7,8,9,10} //普通数组
	fmt.Println(a)

	s2:=a[1:10] //对数组进行切片 以下标1开始，以10结束但不包含10
	fmt.Println(s2)
	a[3]=33 //当修改了a[3]发现s2对应位置也发生了变化
	fmt.Println(s2)

	s2=append(s2,1,2,3) //这里增加三个元素之后已经超过底层数组a的容量，这是就会重新分配一个数组给s2,并追加元素
	fmt.Println(s2)
	fmt.Println(a)
	/*
	输出结果：
	[2 3 33 5 6 7 8 9 10 1 2 3]
        [1 2 3 33 5 6 7 8 9 10]   发现并不影响
	 */

	/*
	slice的第二种声明方式
	3表示存储的元素的个数
	10表示初始分配的连续空间大小，当元素个数超过分配空间，容量会加倍
	相当于在开始创建像内存索要空间
	 */
	s3:=make([]int,3,10)
	fmt.Println(len(s3))
	s3=append(s3,1,2,3)
	fmt.Println(s3)


	/*
	不管是长的拷贝到短的还是相反，都只会改变都有的索引部分
	 */
	s4:=[]int{1,2,3,4,5}
	s5:=[]int{6,7,8}
	copy(s4,s5)
	fmt.Println(s4,s5)
	copy(s5,s4)
	fmt.Println(s5,s4)

}