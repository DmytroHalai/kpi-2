package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/DmytroHalai/kpi-2"
)

func main() {
	exprFlag := flag.String("e", "", "Postfix expression to convert")
	fileFlag := flag.String("f", "", "Input file containing the postfix expression")
	outputFlag := flag.String("o", "", "Output file to store the result")

	flag.Parse()

	if *exprFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: specify only one input source (-e or -f)")
		os.Exit(1)
	}

	var input io.Reader
	if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening input file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else if *exprFlag != "" {
		input = strings.NewReader(*exprFlag)
	} else {
		fmt.Fprintln(os.Stderr, "Error: no input provided (-e or -f required)")
		os.Exit(1)
	}

	var output io.Writer = os.Stdout
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

}
