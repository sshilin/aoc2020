package day6

import (
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day6_input.txt")
)

func Test_part1(t *testing.T) {
	counts := 0

	answers := map[rune]bool{}
	for _, line := range input {
		if line != "" {
			for _, r := range line {
				answers[r] = true
			}
		} else {
			counts += len(answers)
			answers = map[rune]bool{}
		}
	}
	counts += len(answers)
	t.Log("Result:", counts)
}

func Test_part2(t *testing.T) {
	counts := 0
	groupSize := 0

	answers := map[rune]int{}
	for _, line := range input {
		if line != "" {
			groupSize++
			for _, r := range line {
				if _, ok := answers[r]; ok {
					answers[r]++
				} else {
					answers[r] = 1
				}
			}
		} else {
			for _, v := range answers {
				if v == groupSize {
					counts++
				}
			}
			answers = map[rune]int{}
			groupSize = 0
		}
	}

	for _, v := range answers {
		if v == groupSize {
			counts++
		}
	}

	t.Log("Result:", counts) // 3276
}
