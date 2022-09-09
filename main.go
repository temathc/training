package main

import "fmt"

func main() {
	fmt.Println("Hi")
	ss := sum(4, 5)
	fmt.Println("Sum =", ss)
}

func sum(a int, b int) int {
	return a + b
}
