package main

import "fmt"

func maina() {
	var sum int
	for a := 6; a >= 0; a -= 1 {
		sum += a
	}

	fmt.Println(sum)
}
