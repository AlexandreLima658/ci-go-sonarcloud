package main

import "fmt"

func main() {

	result, err := sum(5, 5)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Result", result)
	}

}

func sum(a, b int) (int, error) {
	return a + b, nil
}
