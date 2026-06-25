package mathskills

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumbers(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number on line %d: %q", lineNumber, line)
		}

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(numbers) == 0 {
		return nil, fmt.Errorf("no numbers found")
	}

	return numbers, nil
}
