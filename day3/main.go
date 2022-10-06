package main

import "fmt"

func main() {
	fmt.Println(Fib100())
}

func Fib100() []int {
	fib := []int{0, 1}

	for i := 2; i < 100; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib
}
