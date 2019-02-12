package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main() {
	var n int
	fmt.Print("Enter number: ")
	fmt.Scan(&n)
	fmt.Println("Fibonacci Values:")
	c := make(chan int, n)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("%d ", i)
	}
}
