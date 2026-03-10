package main

import (
	"fmt"
	"time"
)

//defer 典型的应用场景是释放资源，比如关闭文件句柄，释放数据库连接等

func basicDefer() {
	fmt.Println("A")
	defer fmt.Println(1) //defer用于注册一个延迟调用（在函数返回之前调用）
	fmt.Println("B")
	defer fmt.Println(2) //如果同一个函数里有多个defer,则后注册的先执行
	fmt.Println("C")
	defer fmt.Println(3)
	fmt.Println("D")

	//执行的顺序应该是ABCD321
}

func derfer_exe_time() (i int) {
	i = 9
	defer func() { //defer后可以跟一个func
		fmt.Printf("first i = %d\n", i)
	}()
	defer func(i int) {
		fmt.Printf("second i = %d\n", i)
	}(i)
	defer fmt.Printf("third i = %d\n", i)

	return 5

	// 5
}

func timeOfWork(arg int) int {
	begin := time.Now()
	defer func() { fmt.Printf("use time %f seconds\n", time.Since(begin).Seconds()) }()

	if arg > 10 {
		time.Sleep(2 * time.Second)
		return 100
	} else {
		time.Sleep(3 * time.Second)
		return 200
	}
}

func mainba() {
	basicDefer()

}
