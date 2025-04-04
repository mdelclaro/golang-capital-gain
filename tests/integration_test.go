package main_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests based on cases from challenge specification pdf file
func TestUseCases(t *testing.T) {
	t.Parallel()

	t.Run("should use case_1 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_1")
		assert.JSONEq(t,
			`[{"tax": 0},{"tax": 0},{"tax": 0}]`,
			stdout.String())
	})

	t.Run("should use case_2 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_2")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 10000.00},{"tax": 0.00}]`,
			stdout.String())
	})

	t.Run("should use case_2+1 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_2+1")
		assert.Equal(t,
			"[{\"tax\":0},{\"tax\":0},{\"tax\":0}]\n[{\"tax\":0},{\"tax\":10000},{\"tax\":0}]\n",
			stdout.String())
	})

	t.Run("should use case_3 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_3")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 0.00},{"tax": 1000.00}]`,
			stdout.String())
	})

	t.Run("should use case_4 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_4")
		assert.JSONEq(t,
			`[{"tax": 0},{"tax": 0},{"tax": 0}]`,
			stdout.String())
	})

	t.Run("should use case_5 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_5")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 10000.00}]`,
			stdout.String())
	})

	t.Run("should use case_6 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_6")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 3000.00}]`,
			stdout.String())
	})

	t.Run("should use case_7 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_7")
		assert.JSONEq(t,
			`[{"tax":0.00}, {"tax":0.00}, {"tax":0.00}, 
			 {"tax":0.00}, {"tax":3000.00}, {"tax":0.00}, 
			 {"tax":0.00}, {"tax":3700.00}, {"tax":0.00}]`,
			stdout.String())
	})

	t.Run("should use case_8 and validate output", func(t *testing.T) {
		stdout := executeProgram("case_8")
		assert.JSONEq(t,
			`[{"tax":0.00},{"tax":80000.00},{"tax":0.00},{"tax":60000.00}]`,
			stdout.String())
	})
}

func executeProgram(fileName string) bytes.Buffer {
	cmd := exec.Command("go", "run", "../cmd/main.go")
	cmd.Env = os.Environ()

	currentPathFile, _ := os.Getwd()
	data, err := os.ReadFile(fmt.Sprintf("%s/../payloads/%s", currentPathFile, fileName))
	if err != nil {
		panic(err)
	}

	cmd.Stdin = strings.NewReader(string(data))

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		panic(err)
	}

	return stdout
}
