package day1

import (
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day3_input.txt")
)

type Slope struct {
	x, y int
}

type State struct {
	slope     Slope
	dx, trees int
}

func evaluate(slopes []Slope) []int {
	row := 0
	strLen := len(input[0])
	states := make([]*State, 0)
	trees := make([]int, 0)

	for _, slope := range slopes {
		states = append(states, &State{slope, 0, 0})
	}

	for _, line := range input {
		for _, state := range states {
			if row%state.slope.y == 0 {
				if line[state.dx] == '#' {
					state.trees++
				}
				state.dx = (state.dx + state.slope.x) % strLen
			}
		}
		row++
	}

	for _, s := range states {
		trees = append(trees, s.trees)
	}
	return trees
}

func Test_part1(t *testing.T) {
	trees := evaluate([]Slope{
		{3, 1},
	})

	t.Log("Result:", trees[0]) // 265
}

func Test_part2(t *testing.T) {
	trees := evaluate([]Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	})
	result := 1
	for _, t := range trees {
		result *= t
	}

	t.Log("Result:", result) // 3154761400
}
