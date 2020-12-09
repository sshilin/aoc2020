package day9

import (
	"sort"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadInts("day9_input.txt")
)

func searchInPreamble(preamble []int, x int) bool {
	for i := 0; i < len(preamble)-1; i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == x {
				return true
			}
		}
	}
	return false
}

func Test_part1(t *testing.T) {
	preamble := 25

	for i := preamble; i < len(input); i++ {
		if !searchInPreamble(input[i-preamble:i], input[i]) {
			t.Log("Result:", input[i]) // 2089807806
			return
		}
	}
}

func Test_part2(t *testing.T) {
	find := 2089807806

	for offset := 0; offset < len(input); offset++ {
		sum := 0
		set := make([]int, 0)
		for i := offset; i < len(input); i++ {
			sum += input[i]
			set = append(set, input[i])
			if sum == find {
				sort.Ints(set)
				t.Log("Result:", set[0]+set[len(set)-1]) // 245848639
				return
			}
		}
	}
}
