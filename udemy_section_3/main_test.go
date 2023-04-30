package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "$457600.00") {
		t.Error("wrong balance returned!")
	}
}
