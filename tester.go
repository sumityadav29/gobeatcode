package main

import "fmt"

func addNums(x int, y int) int {
	var sum int
	sum = x + y
	return sum
}

func main() {
	result := addNums(5, 10)
	fmt.Println(result)
}

// package should be main ?
// main function should always be there
// main function can be present only in main package ?
