package main

import "fmt"

func add(x, y int) int {
	return x * y
}

// Return multiple results
func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}