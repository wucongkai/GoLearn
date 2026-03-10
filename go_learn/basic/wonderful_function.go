package main

import (
	"fmt"
	"slices"
)

func mainp() {
	arr := []int{3, 7, 4, 6, 8, 1}
	slices.Sort(arr)
	fmt.Println(arr)

	slices.SortFunc(arr, func(a, b int) int { //自定义排序顺序
		return b - a
	})
	fmt.Println(arr)

	type User struct {
		Age    int
		Height float32
	}

	brr := []*User{{18, 1.8}, {25, 1.7}}

	slices.SortFunc(brr, func(a, b *User) int {
		// return int(b.Height - a.height)
		if b.Height > a.Height {
			return 1
		} else if b.Height < a.Height {
			return -1
		} else {
			return 0
		}
	})

	fmt.Println("最大者", slices.Max(arr))
	fmt.Println("最小者", slices.Min(arr))
	fmt.Println("包含", slices.Contains(arr, 5))

	crr := make([]int, len(arr))
	copy(crr, arr) //最多只能拷贝len（crr）个元素，性能比自己写的for循环要高很多
	fmt.Println(crr)

	fmt.Println("相等", slices.Equal(arr, crr)) //true
	arr[0]++
	fmt.Println("相等", slices.Equal(arr, crr))

	drr := brr
	fmt.Println("相等", slices.Equal(brr, drr)) //true
	brr[0].Age++
	fmt.Println("相等", slices.Equal(brr, drr)) //true
}
