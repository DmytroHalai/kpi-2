package main

import (
	"fmt"
	"os"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

func main() {
	handler := lab2.ComputeHandler{
		Input:  os.Stdin,
		Output: os.Stdout,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
