package main

import "fmt"

//接口是一组行为规范的集合
type EmptyInterface interface{}

//标准库的空接口，any = interface{}
func SumI(args ...EmptyInterface) int {
	rect := 0

	for _, ele := range args {
		switch v := ele.(type) {
		case int:
			rect += v
		case float32:
			rect += int(v)
		default:
			fmt.Printf("不支持的数据类型 %T\n", ele)
		}
	}
	return rect
}

func mainQ() {
	fmt.Println()
	fmt.Println(1)
	fmt.Println("1")

	rect := SumI(1, float32(3.14), true)

	fmt.Println(rect)
}
