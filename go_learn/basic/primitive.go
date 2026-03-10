package main

import "fmt"

func main1() {
	var MyName int
	fmt.Println(MyName)

	var a int = 8
	fmt.Println(a)

	//自动推断b的数据类型
	var b = a
	_ = b //防止未使用报错
	fmt.Println(b)

	c := b //var和:= 二选一
	_ = c
}
