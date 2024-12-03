package solutions

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type DayTwo struct {
	inputPath string
}

func (d DayTwo) SolvePartOne() any {
	data, err := os.ReadFile(d.inputPath)
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	safeReadings := 0

	for _, line := range strings.Split(input, "\n") {
		if d.isLineSafe(line) {
			safeReadings++
		}
	}

	return safeReadings
}

func (d DayTwo) isLineSafe(line string) bool {
	levels := strings.Fields(line)
	decreasing, increasing := false, false
	for i := 0; i < len(levels)-1; i++ {
		prev := d.getLevelValue(levels[i])
		next := d.getLevelValue(levels[i+1])

		delta := math.Abs(float64(prev) - float64(next))

		if delta < 1 || delta > 3 {
			return false
		}

		if prev > next {
			decreasing = true
		}

		if prev < next {
			increasing = true
		}

		if increasing && decreasing {
			return false
		}
	}
	return true
}

func (d DayTwo) SolvePartTwo() any {
	data, err := os.ReadFile(d.inputPath)
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	safeReadings := 0

	for _, line := range strings.Split(input, "\n") {
		if d.isLineSafe(line) {
			safeReadings++
		} else {
			levels := strings.Fields(line)
			safe := false
			for idx := range levels {
				truncatedLevels := make([]string, 0, len(levels)-1)
				truncatedLevels = append(truncatedLevels, levels[:idx]...)
				truncatedLevels = append(truncatedLevels, levels[idx+1:]...)
				if d.isLineSafe(strings.Join(truncatedLevels, " ")) {
					safe = true
				} else {
				}
			}
			if safe {
				safeReadings++
			}
		}
	}

	return safeReadings
}

func (d DayTwo) getLevelValue(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return num
}

func (d DayTwo) GetDayNumber() int {
	return 2
}
