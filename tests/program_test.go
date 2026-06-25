package tests

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestProgramPrintsGoldenOutput(t *testing.T) {
	dataPath := filepath.Join("testdata", "basic.txt")
	command := exec.Command("go", "run", "..", dataPath)

	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\noutput: %s", err, output)
	}

	want := "Average: 2\nMedian: 2\nVariance: 1\nStandard Deviation: 1\n"
	if string(output) != want {
		t.Fatalf("output = %q, want %q", output, want)
	}
}

func TestProgramRequiresOneArgument(t *testing.T) {
	command := exec.Command("go", "run", "..")

	output, err := command.CombinedOutput()
	if err == nil {
		t.Fatalf("go run succeeded, want error")
	}

	if !strings.Contains(string(output), "usage") {
		t.Fatalf("output = %q, want usage message", output)
	}
}
