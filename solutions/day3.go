package solutions

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type DayThree struct {
	inputPath string
}

func (d DayThree) SolvePartOne() any {
	data, err := os.ReadFile(d.inputPath)

	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`mul\(\d*,\d*\)`)
	matches := re.FindAll(data, -1)

	sum := 0
	for _, match := range matches {
		left, right := d.getNumsFromExpression(string(match))
		val := left * right
		sum += val
	}

	return sum
}

func (d DayThree) getNumsFromExpression(expr string) (int, int) {
	parts := strings.Split(strings.TrimSpace(expr), ",")
	left := d.getDigitsFromString(parts[0])
	right := d.getDigitsFromString(parts[1])
	return d.getNumFromString(left), d.getNumFromString(right)
}

func (d DayThree) getDigitsFromString(str string) string {
	var builder strings.Builder

	for _, r := range str {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func (d DayThree) getNumFromString(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func (d DayThree) SolvePartTwo() any {
	data, err := os.ReadFile(d.inputPath)

	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`do\(\)|don\'t\(\)|mul\(\d*,\d*\)`)
	matches := re.FindAll(data, -1)

	sum := 0
	execute := true

	for _, match := range matches {
		if strings.Contains(string(match), "do()") {
			execute = true
			continue
		}

		if strings.Contains(string(match), "don't()") {
			execute = false
			continue
		}

		if !execute {
			continue
		}

		left, right := d.getNumsFromExpression(string(match))
		val := right * left
		sum += val
	}

	return sum
}

func (d DayThree) GetDayNumber() int {
	return 3
}
