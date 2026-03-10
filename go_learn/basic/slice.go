package main

import "fmt"

func slice_init() {
	var s []int //切片声明， len=cap=0、
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{} //初始化， len= cap = 0

	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3) //初始化， len= cap = 3
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3, 5) // len = 3, cap = 5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{1, 2, 3, 4, 5}

	fmt.Println("==================")

	//二维切片初始的一种方式
	s2d := [][]int{
		{1},
		{2, 3}, //二维数组各行的列数是相等的，但二维切片各行的len可以不等
	}

	fmt.Printf("s2d len %d cap %d\n", len(s2d), cap(s2d))
	fmt.Printf("s2d[0] len %d cap %d\n", len(s2d[0]), cap(s2d[0]))
	fmt.Printf("s2d[1] len %d cap %d\n", len(s2d[1]), cap(s2d[1]))
	fmt.Println("===========")
}

//探究capacity扩容规律
func expansion() {
	s := make([]int, 0, 3)
	prevCap := cap(s)

	for i := 0; i < 100; i++ {
		s = append(s, i)
		currCap := cap(s)

		if currCap > prevCap {
			//每次扩容都是扩到原先的2倍
			fmt.Printf("capacity从%d变成%d\n", prevCap, currCap)
			prevCap = currCap
		}
	}
}

func maip() {
	slice_init()
	expansion()
}
