package unit_test

import (
	"bytes"
	"math"
	"testing"

	mathskills "math-skills/internal"
)

func TestCalculateAverage(t *testing.T) {
	got := mathskills.CalculateAverage([]float64{1, 2, 3, 4})
	want := 2.5

	if got != want {
		t.Fatalf("CalculateAverage() = %v, want %v", got, want)
	}
}

func TestCalculateMedianOddCount(t *testing.T) {
	got := mathskills.CalculateMedian([]float64{9, 1, 4})
	want := 4.0

	if got != want {
		t.Fatalf("CalculateMedian() = %v, want %v", got, want)
	}
}

func TestCalculateMedianEvenCount(t *testing.T) {
	got := mathskills.CalculateMedian([]float64{10, 2, 4, 8})
	want := 6.0

	if got != want {
		t.Fatalf("CalculateMedian() = %v, want %v", got, want)
	}
}

func TestCalculateVariance(t *testing.T) {
	got := mathskills.CalculateVariance([]float64{1, 2, 3})
	want := 2.0 / 3.0

	if math.Abs(got-want) > 0.000001 {
		t.Fatalf("CalculateVariance() = %v, want %v", got, want)
	}
}

func TestCalculateStdDev(t *testing.T) {
	got := mathskills.CalculateStdDev(4)
	want := 2.0

	if got != want {
		t.Fatalf("CalculateStdDev() = %v, want %v", got, want)
	}
}

func TestCalculateStatsRejectsEmptyInput(t *testing.T) {
	_, err := mathskills.CalculateStats(nil)

	if err == nil {
		t.Fatal("CalculateStats() error = nil, want error")
	}
}

func TestCalculateStatsWithOneNumber(t *testing.T) {
	got, err := mathskills.CalculateStats([]float64{7})
	if err != nil {
		t.Fatalf("CalculateStats() error = %v, want nil", err)
	}

	if got.Average != 7 || got.Median != 7 || got.Variance != 0 || got.StandardDeviation != 0 {
		t.Fatalf("CalculateStats() = %+v, want all stats based on one value", got)
	}
}

func TestCalculateStatsWithNegativeNumbers(t *testing.T) {
	got, err := mathskills.CalculateStats([]float64{-5, -1, -3})
	if err != nil {
		t.Fatalf("CalculateStats() error = %v, want nil", err)
	}

	if got.Average != -3 || got.Median != -3 {
		t.Fatalf("CalculateStats() = %+v, want average and median -3", got)
	}

	wantVariance := 8.0 / 3.0
	if math.Abs(got.Variance-wantVariance) > 0.000001 {
		t.Fatalf("variance = %v, want %v", got.Variance, wantVariance)
	}
}

func TestCalculateStatsWithLargeNumbers(t *testing.T) {
	got, err := mathskills.CalculateStats([]float64{1000000000, 1000000002, 1000000004})
	if err != nil {
		t.Fatalf("CalculateStats() error = %v, want nil", err)
	}

	if got.Average != 1000000002 || got.Median != 1000000002 {
		t.Fatalf("CalculateStats() = %+v, want average and median 1000000002", got)
	}
}

func TestPrintStatsRoundsOutput(t *testing.T) {
	stdout := &bytes.Buffer{}

	mathskills.PrintStats(stdout, mathskills.StatsResult{
		Average:           2.5,
		Median:            3.49,
		Variance:          4.5,
		StandardDeviation: 2.49,
	})

	want := "Average: 3\nMedian: 3\nVariance: 5\nStandard Deviation: 2\n"
	if stdout.String() != want {
		t.Fatalf("stdout = %q, want %q", stdout.String(), want)
	}
}
