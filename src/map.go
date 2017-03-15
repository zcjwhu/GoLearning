package main

import (
	"fmt"
	"sort"
)
/*
map的用法

-类似其它语言中的哈希表，用于存储key-value类型的数据
-key必须是支持== ！=比较运算的类型，不可以是函数，map或slice
-map查找比线性搜索快很多
-map使用make（）创建

-make([keyType]varvalueType,cap)
-键值对不存在时自动添加，使用delete()删除某键值对
-使用for range对map和slice进行迭代操作
 */

func main(){
	/*
	简单map创建形式
	 */
	var m map[int]string=make(map[int]string)
	m[1]="ok"
	fmt.Println(m)
	delete(m,1) //删除key=1的键值对
	fmt.Println(m)

	/*
	两级map的创建
	 */
	mm:=make(map[int]map[int]string)
	mm[1]=make(map[int]string) //每一级的map都需要进行初始化，其中每一个都需要make
	mm[1][1]="ok"
	fmt.Println(mm)

	/*
	slice和map迭代操作
	slice :i表示索引，v是对应的值,只是对值的操作，并不会修改slice或map，如果需要修改可以通过索引来cauozuo
	map：i表示key,v表示value
	for i,v:=range slice{

	}
	 */
	sm:=make([]map[int]string,5)
	for i:=range sm{
		sm[i] =make(map[int]string,1)
		sm[i][1]="ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	mi:=map[int]string{1:"aa",2:"cc",3:"ee"}
	si:=make([]int,len(mi))
	i:=0
	for k,_:=range mi{
		si[i]=k
		i++
	}
	sort.Ints(si)
	fmt.Println(si)

}
