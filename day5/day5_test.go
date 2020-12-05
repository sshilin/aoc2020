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
	// input = utils.ReadStrings("test.txt")
)

func Test_part1(t *testing.T) {
	result := 0
	for _, line := range input {
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "L", "0")
		line = strings.ReplaceAll(line, "R", "1")
		row, _ := strconv.ParseInt(line[:7], 2, 64)
		col, _ := strconv.ParseInt(line[7:], 2, 64)

		id := row*8 + col

		if int(id) > result {
			result = int(id)
		}
	}
	t.Log("Result:", result)
}

func Test_part2(t *testing.T) {

	ids := make([]int, 0)

	for _, line := range input {
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "L", "0")
		line = strings.ReplaceAll(line, "R", "1")
		row, _ := strconv.ParseInt(line[:7], 2, 32)
		col, _ := strconv.ParseInt(line[7:], 2, 32)

		id := row*8 + col

		ids = append(ids, int(id))

		// t.Log(line, "[row=", row, " col=", col, "]")
	}

	sort.Ints(ids)
	seat := ids[0]
	for i := 1; i < len(ids); i++ {
		seat++
		t.Log(ids[i])
		if ids[i] != seat {
			t.Log("================>", seat) // 623?
		}
	}

	t.Log(ids)
}
