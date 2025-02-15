package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	t.Run("Should calculate sum", func(t *testing.T) {

		total := sum(5, 5)
		if total != 10 {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
		}
	})

	t.Run("Should calculate sub", func(t *testing.T) {

		total := sub(5, 5)
		if total != 0 {
			t.Errorf("Sub was incorrect, got: %d, want: %d.", total, 0)
		}
	})

	t.Run("Should calculate times", func(t *testing.T) {

		total := times(5, 5)
		if total != 25 {
			t.Errorf("Times was incorrect, got: %d, want: %d.", total, 25)
		}
	})

	t.Run("Should print sum", func(t *testing.T) {
		
		result := sum(5, 5)
		_, err := fmt.Println(result)
		
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result != 10 {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 10)
		}
	})

	t.Run("Should print sub", func(t *testing.T) {
		
		result := sub(5, 5)
		_, err := fmt.Println(result)
		
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result != 0 {
			t.Errorf("Sub was incorrect, got: %d, want: %d.", result, 0)
		}
	})
	t.Run("Should print times", func(t *testing.T) {
		
		result := times(5, 5)
		_, err := fmt.Println(result)
		
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result != 25 {
			t.Errorf("Times was incorrect, got: %d, want: %d.", result, 25)
		}
	})


}
