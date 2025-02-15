package main

import (
	"testing"
)

func Test(t *testing.T) {

	t.Run("Should calculate sum", func(t *testing.T) {

		total, err := sum(5, 5)

		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if total != 10 {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)

		}
		

	})

}
