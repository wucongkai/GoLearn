package main

import "fmt"

/**
Age和int可以互相做强制类型转换。存储的数据类型（成员变量）是一样的，但行为（成员）
方法是不一样的
**/

type Age int
type Tall = int //tall和int完全等价，不需要显式做类型转换

func mainC() {
	var b int
	var a Age
	var c Tall

	fmt.Println(a, b, c)
}
