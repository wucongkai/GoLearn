package main

import (
	"fmt"
	"time"
)

const (
	TIME_FMT = "2006-01-02 15:04:05.000"
	DATE_FMT = "20060102"
)

func mainaa() {
	fmt.Println("start")
	fmt.Println(time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
	fmt.Println("bye-bye")

	t1 := time.Now()
	fmt.Printf("t1 = %s\n", t1.Format(TIME_FMT))
}

// 周期时执行任务
func ticker() {
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tk := time.NewTicker(time.Second)

	for i := 0; i < 10; i++ {
		<-tk.C
		fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	}

	tk.Stop()
}

// 定时执行任务
func timer() {
	fmt.Printf("现在的时间是%s\n", time.Now().Format(TIME_FMT))
	tm := time.NewTimer(time.Second * 10)
	<-tm.C
	defer tm.Stop()
	fmt.Printf("现在的时间是%s\n", time.Now().Format(TIME_FMT))
}
