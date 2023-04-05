package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func divide(a int, b int) int {
	if a == 1000 {
		fmt.Println("Hello")
	}
	if b == 0 {
		return -1
	}
	return a / b
}
