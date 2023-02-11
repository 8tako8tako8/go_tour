package main

import "fmt"

func fibonacci() func() int {
	prev, cur := 0, 1
	return func() int {
		fib := prev
		prev, cur = cur, prev + cur
		return fib
	}
}

// Exercise: Fibonacci closure
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
