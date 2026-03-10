package main

import (
	"fmt"
	"time"
)

func mai() {
	c, d := 3, 5
	arg1(c, d)
}

// a， b是形参，形参是函数内部的局部变量，实参的值会拷贝给形参
func arg1(a int, b int) {
	a = a + b //函数内部修改形参的值，实参的值不受影响
	//return    //函数返回，return后面的语句不会再执行
	//fmt.Println("我不会被输出")
}

func arg2(a, b int) { //参数类型相同时可以只写一次
	a = a + b
	//不写return时,默认执行完最后一行代码返回
}

func arg3(a, b *int) { //如果想好通过函数修改实参，就需要指针类型
	*a = *a + *b
	*b = 888
}

func no_arg() { //函数可以没有参数，也没有返回值
	fmt.Println("hello golang")
}

func return1(a, b int) int { //函数需要返回一个int型数据
	a = a + b
	c := a //声明并初始化一个变量c
	return c
}

func return2(a, b int) (c int) { //返回变量c已经声明好了
	a = a + b
	c = a  //直接使用c
	return //由于函数要求有返回值，即使给c赋值了，也需要显示写return
}

func return3() (int, int) { //可以没有形参，可以返回多个参数
	now := time.Now()
	return now.Hour(), now.Minute()
}

// 不定长参数
func vaeiavle_ength_arg(a int, other ...int) int {
	//调用该函数时，other可以对应0个参数也可以对应多个参数
	sum := a
	//不定长参数实际上是slice类型
	for _, ele := range other {
		sum += ele
	}

	if len(other) > 0 {
		fmt.Printf("first ele %d len %d cap %d\n", other[0], len(other), cap(other))
	} else {
		fmt.Printf("len %d cap %d \n", len(other), cap(other))
	}

	return sum
}

// 不定长参数和递归
func sum(arr ...int) int {
	s := 0
	if len(arr) == 0 {
		return s
	}
	s += arr[0]
	s += sum(arr[1:]...)
	return s
}

//计算feibolaye
