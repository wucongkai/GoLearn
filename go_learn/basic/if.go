package main

import "fmt"

func main0() {
	if 5 > 9 {
		fmt.Println("A")
	}

	var a int = 10

	if a < 5 {
		fmt.Println("B")
	}

	if b := 8; b > a {
		fmt.Print("b > a")
	}

}
