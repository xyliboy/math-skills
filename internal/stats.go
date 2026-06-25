package mathskills

import (
	"fmt"
	"io"
	"math"
	"sort"
)

type StatsResult struct {
	Average           float64
	Median            float64
	Variance          float64
	StandardDeviation float64
}

func CalculateStats(numbers []float64) (StatsResult, error) {
	if len(numbers) == 0 {
		return StatsResult{}, fmt.Errorf("no numbers found")
	}

	variance := CalculateVariance(numbers)

	return StatsResult{
		Average:           CalculateAverage(numbers),
		Median:            CalculateMedian(numbers),
		Variance:          variance,
		StandardDeviation: CalculateStdDev(variance),
	}, nil
}

func CalculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}

func CalculateMedian(numbers []float64) float64 {
	sortedNumbers := append([]float64(nil), numbers...)
	sort.Float64s(sortedNumbers)

	middle := len(sortedNumbers) / 2
	if len(sortedNumbers)%2 == 1 {
		return sortedNumbers[middle]
	}

	return (sortedNumbers[middle-1] + sortedNumbers[middle]) / 2
}

func CalculateVariance(numbers []float64) float64 {
	average := CalculateAverage(numbers)
	sumSquares := 0.0

	for _, number := range numbers {
		difference := number - average
		sumSquares += difference * difference
	}

	return sumSquares / float64(len(numbers))
}

func CalculateStdDev(variance float64) float64 {
	return math.Sqrt(variance)
}

func PrintStats(writer io.Writer, stats StatsResult) {
	fmt.Fprintf(writer, "Average: %.0f\n", math.Round(stats.Average))
	fmt.Fprintf(writer, "Median: %.0f\n", math.Round(stats.Median))
	fmt.Fprintf(writer, "Variance: %.0f\n", math.Round(stats.Variance))
	fmt.Fprintf(writer, "Standard Deviation: %.0f\n", math.Round(stats.StandardDeviation))
}
