package main

import "fmt"

func main() {
	//fmt.Sprint(math.Sqrt())
	switch {
	case 1 < 2:
		fmt.Println(1)
	case 2 < 2:
		fmt.Println(2)
	default:
		fmt.Println(3)

	}
}
