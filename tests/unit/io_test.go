package unit_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	mathskills "math-skills/internal"
)

func TestReadNumbers(t *testing.T) {
	path := writeTestFile(t, "numbers.txt", "1\n\n-2\n3.5\n")

	got, err := mathskills.ReadNumbers(path)
	if err != nil {
		t.Fatalf("ReadNumbers() error = %v, want nil", err)
	}

	want := []float64{1, -2, 3.5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ReadNumbers() = %v, want %v", got, want)
	}
}

func TestReadNumbersRejectsInvalidLine(t *testing.T) {
	path := writeTestFile(t, "invalid.txt", "1\nabc\n3\n")

	_, err := mathskills.ReadNumbers(path)
	if err == nil {
		t.Fatal("ReadNumbers() error = nil, want error")
	}
}

func TestReadNumbersRejectsEmptyFile(t *testing.T) {
	path := writeTestFile(t, "empty.txt", "\n\n")

	_, err := mathskills.ReadNumbers(path)
	if err == nil {
		t.Fatal("ReadNumbers() error = nil, want error")
	}
}

func writeTestFile(t *testing.T, name string, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("os.WriteFile() error = %v, want nil", err)
	}

	return path
}
