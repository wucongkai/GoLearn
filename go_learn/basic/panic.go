package main

import "fmt"

func defer_panic() {
	defer fmt.Println(1)
	var arr []int
	n := 0
	// defer fmt.Println(1/n) 在注册时就要计算1/n， 发生panic
	defer func() {
		_ = arr[n]
		_ = 1 / n            //defer func 内部发生panic， main协程不会exit，其他defer还可以正常执行
		defer fmt.Println(2) //上面代码发生panic，所以本行的defer没有注册成功
	}()

	defer fmt.Println(3)
}

func mainaaa() {

}
