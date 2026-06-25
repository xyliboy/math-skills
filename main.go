package main

import (
	"fmt"
	"io"
	"os"

	mathskills "math-skills/internal"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	if len(args) != 1 {
		err := fmt.Errorf("usage: go run . <data-file>")
		fmt.Fprintln(stderr, err)
		return err
	}

	numbers, err := mathskills.ReadNumbers(args[0])
	if err != nil {
		fmt.Fprintln(stderr, err)
		return err
	}

	stats, err := mathskills.CalculateStats(numbers)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return err
	}

	mathskills.PrintStats(stdout, stats)
	return nil
}
