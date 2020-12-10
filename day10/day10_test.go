package day10

import (
	"sort"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadInts("day10_input.txt")
)

func Test_part1(t *testing.T) {
	sort.Ints(input)

	ones, threes := 0, 0
	output := 0

	for _, adapter := range input {
		diff := adapter - output
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
		output = adapter
	}
	threes++
	t.Log("Result:", ones*threes) // 2312
}

func Test_part2(t *testing.T) {
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)
	combinations := map[int]int{0: 1}

	for _, adapter := range input {
		if prev, ok := combinations[adapter-1]; ok {
			combinations[adapter] += prev
		}
		if prev, ok := combinations[adapter-2]; ok {
			combinations[adapter] += prev
		}
		if prev, ok := combinations[adapter-3]; ok {
			combinations[adapter] += prev
		}
	}
	t.Log("Result:", combinations[input[len(input)-1]]) // 12089663946752
}
