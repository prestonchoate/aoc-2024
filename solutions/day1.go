package solutions

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DayOne struct{
	inputPath string
	data []byte
}

func (d DayOne) SolvePartOne() any {
	d.loadInputData()
	leftList, rightList := d.getSortedLists()

	sum := 0
	for i := 0; i < len(leftList); i++ {
		leftNum := leftList[i]
		rightNum := rightList[i]
		sum += int(math.Abs(float64(leftNum) - float64(rightNum)))
	}

	return sum
}

func (d DayOne) SolvePartTwo() any {
	d.loadInputData()
	leftList, rightList := d.getSortedLists()

	counts := map[int]int{}

	for i := 0; i < len(rightList); i++ {
		count, ok := counts[rightList[i]]
		if !ok {
			counts[rightList[i]] = 1
			continue
		}
		counts[rightList[i]] = count + 1
	}

	sum := 0

	for i := 0; i < len(leftList); i++ {
		num := leftList[i]
		count, ok := counts[num]
		if !ok {
			continue
		}
		sum += num * count
	}

	return sum
}

func (d *DayOne) loadInputData() {
	data, err := os.ReadFile(d.inputPath)
	if err != nil {
		panic(err)
	}

	d.data = data
}

func (d DayOne) GetDayNumber() int {
	return 1
}

func (d *DayOne) formatInput() string {
	return strings.TrimSpace(string(d.data))
}

func (d *DayOne) getSortedLists() ([]int, []int) {
	input := d.formatInput()

	leftList := []int{}
	rightList := []int{}

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Fields(line)

		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList

}
