package main

import "fmt"

func main() {

	fmt.Println(sum(5, 5))
	fmt.Println(sub(5, 5))
	fmt.Println(times(5, 5))

}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}
