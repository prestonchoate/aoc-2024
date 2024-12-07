package solutions

import (
	"os"
	"strings"
)

type DayFour struct {
	inputPath string
}

func (d DayFour) SolvePartOne() any {
	data, err := os.ReadFile(d.inputPath)
	if err != nil {
		panic(err)
	}

	input := string(data)

	return d.GetTotalWordCount(strings.TrimSpace(input))
}

func (d DayFour) GetTotalWordCount(input string) int {
	lines := strings.Split(input, "\n")
	finds := 0
	for i, line := range lines {
		for j, char := range line {
			if string(char) != "X" {
				continue
			}
			if d.checkRight(line, j) {
				finds++
			}

			if d.checkLeft(line, j) {
				finds++
			}

			if d.checkUp(lines, j, i) {
				finds++
			}

			if d.checkDown(lines, j, i) {
				finds++
			}

			finds += d.checkDiags(lines, j, i)
		}
	}

	return finds
}

func (d DayFour) checkDiags(lines []string, xPos int, yPos int) int {
	count := 0

	dr, dl, ur, ul := true, true, true, true

	if yPos-3 < 0 {
		ur = false
		ul = false
	}

	if yPos+3 >= len(lines) {
		dr = false
		dl = false
	}

	if xPos-3 < 0 {
		ul = false
		dl = false
	}

	if xPos+3 >= len(lines[yPos]) {
		ur = false
		dr = false
	}

	if !dr && !dl && !ur && !ul {
		return 0
	}

	// check down right diagonal
	if dr && string(lines[yPos+1][xPos+1]) == "M" &&
		string(lines[yPos+2][xPos+2]) == "A" &&
		string(lines[yPos+3][xPos+3]) == "S" {
		count++
	}

	// check down left diagonal
	if dl && string(lines[yPos+1][xPos-1]) == "M" &&
		string(lines[yPos+2][xPos-2]) == "A" &&
		string(lines[yPos+3][xPos-3]) == "S" {
		count++
	}

	// check up right diagonal
	if ur && string(lines[yPos-1][xPos+1]) == "M" &&
		string(lines[yPos-2][xPos+2]) == "A" &&
		string(lines[yPos-3][xPos+3]) == "S" {
		count++
	}

	// check up left diagonal
	if ul && string(lines[yPos-1][xPos-1]) == "M" &&
		string(lines[yPos-2][xPos-2]) == "A" &&
		string(lines[yPos-3][xPos-3]) == "S" {
		count++
	}

	return count
}

func (d DayFour) checkRight(line string, pos int) bool {
	if pos+3 >= len(line) {
		return false
	}

	if string(line[pos+1]) != "M" {
		return false
	}

	if string(line[pos+2]) != "A" {
		return false
	}

	if string(line[pos+3]) != "S" {
		return false
	}

	return true
}

func (d DayFour) checkLeft(line string, pos int) bool {
	if pos-3 < 0 {
		return false
	}

	if string(line[pos-1]) != "M" {
		return false
	}

	if string(line[pos-2]) != "A" {
		return false
	}

	if string(line[pos-3]) != "S" {
		return false
	}

	return true
}

func (d DayFour) checkUp(lines []string, xPos int, yPos int) bool {
	if yPos-3 < 0 {
		return false
	}

	m := string(lines[yPos-1][xPos])
	a := string(lines[yPos-2][xPos])
	s := string(lines[yPos-3][xPos])

	if m != "M" || a != "A" || s != "S" {
		return false
	}

	return true
}

func (d DayFour) checkDown(lines []string, xPos int, yPos int) bool {
	if yPos+3 >= len(lines) {
		return false
	}

	m := string(lines[yPos+1][xPos])
	a := string(lines[yPos+2][xPos])
	s := string(lines[yPos+3][xPos])

	if m != "M" || a != "A" || s != "S" {
		return false
	}

	return true
}

func (d DayFour) SolvePartTwo() any {

	return 0
}

func (d DayFour) GetDayNumber() int {
	return 4
}
