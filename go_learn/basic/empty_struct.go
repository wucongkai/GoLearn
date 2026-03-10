package main

import (
	"fmt"
	"time"
	"unsafe"
)

type ETS struct {
}

// 所有空结构体指向同一个地址（内核完全一样的）
func allEmptyStructISSame() {
	var a ETS //等价于var c struct{}
	var b ETS //等价于var d struct{}

	fmt.Printf("Address of a %p b %p\n", &a, &b)
	fmt.Printf("size of a %d b %d\n", unsafe.Sizeof(a), unsafe.Sizeof(b))
}

// 空结构体的应用场景
func scenariosOfEmptyStruct() {
	set := map[int]struct{}{
		1: {},
		4: {},
		7: {},
	}

	if _, exists := set[5]; exists {
		fmt.Println("5是存在的")
	} else {
		fmt.Println("5是不存在的")
	}

	blocker := make(chan struct{}, 1)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("done")
		blocker <- struct{}{}
	}()
	<-blocker
}

func mainW() {
	allEmptyStructISSame()
}
