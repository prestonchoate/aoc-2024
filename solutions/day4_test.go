package solutions

import (
	"testing"
)

func TestDayFourPartOne(t *testing.T) {
	input :=
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	expect := 18

	d4 := DayFour{}

	got := d4.GetTotalWordCount(input)

	if got != expect {
		t.Fatalf("Expected %v, Got: %v\n", expect, got)
	}

}
