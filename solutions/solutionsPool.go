package solutions

type Solution interface {
	SolvePartOne() any
	SolvePartTwo() any
	GetDayNumber() int
}

type SolutionPool struct {
	instances []Solution
}

func GetSolutionPool() SolutionPool {
	return SolutionPool{
		instances: []Solution{
			DayOne{
				inputPath: "./inputs/day1.txt",
			},
			DayTwo{
				inputPath: "./inputs/day2.txt",
			},
		},
	}
}

func (Sp *SolutionPool) GetSolutionForDay(day int) *Solution {
	for _, solution := range Sp.instances {
		if solution.GetDayNumber() == day {
			return &solution
		}
	}
	return nil
}
