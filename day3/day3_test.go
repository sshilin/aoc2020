package day3

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

func evaluate(slope Slope) int {
	col, trees := 0, 0
	for rowNum, line := range input {
		if rowNum%slope.y == 0 {
			if line[col] == '#' {
				trees++
			}
			col = (col + slope.x) % len(input[0])
		}
	}
	return trees
}

func Test_part1(t *testing.T) {
	t.Log("Result:", evaluate(Slope{3, 1})) // 265
}

func Test_part2(t *testing.T) {
	result := 1
	for _, slope := range []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		result *= evaluate(slope)
	}

	t.Log("Result:", result) // 3154761400
}
