package main

import "fmt"

func calculator(op string, x float64, y float64) (res float64) {

	switch op {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		res = x / y
	default:
		fmt.Println("Invalid Opearator")
	}
	return
}

func main() {
	fmt.Println(calculator("+", 1, 2))
}
