package main

import "fmt"

func rangeChannel() {
	var ch chan int //声明

	// if ch == nil {
	// 	fmt.Printf("ch is nil, ch len %d cap  %d\n", len(ch), cap(ch))
	// }

	//ch <- 2 **不能向nil chan里发送数据
	if len(ch) == 0 { //引用类型未初始化都是nil，可以对它们执行len（）函数，返回0
		fmt.Println("ch length is 0")
	}

	ch = make(chan int, 8) //初始化，环形队列里可以容纳8个int
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Printf("ch len %d cap %d\n", len(ch), cap(ch))
	v := <-ch //从管道里取走（recv）数据
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)
	fmt.Println()

	close(ch)
	//遍历并取走（receive）管道里的元素。当管道里已无剩余元素且没有close管道时， receive操作会
	//一直阻塞，最终报deadlock.且管道为空且被close后，for循环退出
	for ele := range ch {
		fmt.Println(ele)
	}

	c := make(chan int, 10)
	send(c)
	recv(c)
}

//只能向channel里写满数据, chan<- int只允许写入
func send(c chan<- int) {
	c <- 1
}

//只能取channel中的数据
func recv(c <-chan int) {
	v := <-c
	fmt.Printf("take %d from read-only channel\n", v)
}

func mainfsa() {
	rangeChannel()
}
