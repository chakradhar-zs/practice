package main

import "fmt"

func Sum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func main() {
	fmt.Println(Sum(10))
}