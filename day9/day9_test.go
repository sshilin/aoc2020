package day9

import (
	"sort"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadInts("day9_input.txt")
)

func searchInPreamble(len, offset, x int) bool {
	for i := offset; i < offset+len-1; i++ {
		for j := i + 1; j < offset+len; j++ {
			if input[i]+input[j] == x {
				return true
			}
		}
	}
	return false
}

func Test_part1(t *testing.T) {
	offset := 0
	preambleLen := 25
	for i := preambleLen; i < len(input); i++ {
		if !searchInPreamble(preambleLen, offset, input[i]) {
			t.Log("Result:", input[i]) // 2089807806
			return
		}
		offset++
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

// func Test_part3(t *testing.T) {
// 	a := make([]int, 0)

// 	a = append(a, 23)
// 	a = append(a, 3)
// 	a = append(a, 10)

// 	sort.Ints(a)
// 	t.Log("====>", a)
// 	t.Log("====>", sort.SearchInts(a, 100) && a)
// }
