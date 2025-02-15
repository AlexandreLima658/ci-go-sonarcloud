package main

import (
	"bytes"
	"os"
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

	t.Run("Should calculate times", func(t *testing.T) {
		
		total, err := times(5, 5)

		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if total != 25 {
			t.Errorf("Times was incorrect, got: %d, want: %d.", total, 25)

		}
	})
}

func TestMain(t *testing.T) {
	// Redireciona a saída padrão para um buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Chama a função main
	main()

	// Restaura a saída padrão
	w.Close()
	os.Stdout = oldStdout

	// Lê a saída do buffer
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Verifica se a saída corresponde ao esperado
	expectedOutput := "Result 10\n"
	if output != expectedOutput {
		t.Errorf("Saída inesperada: \nEsperado: %q\nObtido: %q", expectedOutput, output)
	}




}
