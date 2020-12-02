package day2

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day2_input.txt")
)

func Test_part1(t *testing.T) {
	valid := 0
	for _, line := range input {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == ':' || r == ' '
		})

		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		r := regexp.MustCompile(parts[2])
		res := r.FindAllString(parts[3], -1)
		if len(res) >= min && len(res) <= max {
			valid++
		}
	}
	t.Log("Result:", valid) // 422
}

func Test_part2(t *testing.T) {
	valid := 0
	for _, line := range input {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == ':' || r == ' '
		})

		pos1, _ := strconv.Atoi(parts[0])
		pos2, _ := strconv.Atoi(parts[1])

		pos1--
		pos2--

		if (string(parts[3][pos1]) == parts[2]) != (string(parts[3][pos2]) == parts[2]) {
			valid++
		}
	}
	t.Log("Result:", valid) // 451
}
