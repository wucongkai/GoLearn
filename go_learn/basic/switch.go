package main

import "fmt"

func switch_basic() {
	color := "yellow"
	//用switch-case-default模拟if-else
	switch color {
	case "green":
		fmt.Println("go")
	default:
		fmt.Println("stop")
	}
}
