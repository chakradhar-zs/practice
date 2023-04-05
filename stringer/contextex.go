package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timed Out")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("Doing Something")
		}
	}
	time.Sleep(500 * time.Millisecond)
}

func main() {
	fmt.Println("Context Problem")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	doSomething(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Oh no i have exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}
