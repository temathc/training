package main

import "fmt"

func main() {
	fmt.Println("Hi")
	ss := sum(4, 5)
	fmt.Println("Sum =", ss)
	s := "привет"
	checkStr(s)
}

func sum(a int, b int) int {
	return a + b
}

func checkStr(str string) {
	for ind, val := range str {
		fmt.Printf("Индекс = %d, значение = %s\n", ind, string(val))
	}
}
