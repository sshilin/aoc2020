package day1

import (
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadInts("day1_input.txt")
)

func Test_part1(t *testing.T) {
	for a := 0; a < len(input)-1; a++ {
		for b := a + 1; b < len(input); b++ {
			if input[a]+input[b] == 2020 {
				t.Log("Result:", input[a]*input[b])
				return
			}
		}
	}
}

func Test_part2(t *testing.T) {
	for a := 0; a < len(input)-2; a++ {
		for b := a + 1; b < len(input)-1; b++ {
			for c := b + 1; c < len(input); c++ {
				if input[a]+input[b]+input[c] == 2020 {
					t.Log("Result:", input[a]*input[b]*input[c])
					return
				}
			}
		}
	}
}
