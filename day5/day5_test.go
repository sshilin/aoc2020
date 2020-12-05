package day5

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day5_input.txt")
)

func Test_part1(t *testing.T) {
	maxID := 0
	for _, line := range input {
		line = strings.Map(func(r rune) rune {
			if r == 'F' || r == 'L' {
				return '0'
			}
			return '1'
		}, line)

		row, _ := strconv.ParseInt(line[:7], 2, 64)
		col, _ := strconv.ParseInt(line[7:], 2, 64)

		id := int(row*8 + col)

		if id > maxID {
			maxID = id
		}
	}

	t.Log("Result:", maxID) // 885
}

func Test_part2(t *testing.T) {
	seatIDs := make([]int, 0)

	for _, line := range input {
		line = strings.Map(func(r rune) rune {
			if r == 'F' || r == 'L' {
				return '0'
			}
			return '1'
		}, line)

		row, _ := strconv.ParseInt(line[:7], 2, 64)
		col, _ := strconv.ParseInt(line[7:], 2, 64)

		seatIDs = append(seatIDs, int(row*8+col))
	}

	sort.Ints(seatIDs)
	mySeatID := seatIDs[0]

	for i := 1; i < len(seatIDs); i++ {
		mySeatID++
		if seatIDs[i] != mySeatID {
			t.Log("Result:", mySeatID) // 623
			return
		}
	}
}
