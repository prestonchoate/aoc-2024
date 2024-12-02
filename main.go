package main

import (
	"flag"
	"fmt"

	"github.com/prestonchoate/aoc-2024/solutions"
)

func main() {

  var dayFlag = flag.Int("d", 1, "day number to solve")

  flag.Parse()

  pool := solutions.GetSolutionPool()

  solution := pool.GetSolutionForDay(*dayFlag)

  if solution == nil {
    fmt.Println("No solution exists for day: ", *dayFlag)
    return
  }

  fmt.Printf("Part one: %v\nPart two:: %v\n", (*solution).SolvePartOne(), (*solution).SolvePartTwo())
}

