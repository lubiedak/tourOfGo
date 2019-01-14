package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n, m, nextVal := 0, 1, 0
	return func() int {
		nextVal, m, n = n, nextVal, n+m
		return nextVal
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

