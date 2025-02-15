package main

import "fmt"

func main() {

	fmt.Println(sum(5, 5))
	fmt.Println(sub(5, 5))
	fmt.Println(times(5, 5))

}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}
