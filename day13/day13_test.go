package day13

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day13_input.txt")
)

func Test_part1(t *testing.T) {
	timestamp, _ := strconv.Atoi(input[0])

	candidateTime := math.MaxInt64
	candidateID := 0

	for _, bus := range strings.Split(input[1], ",") {
		if bus != "x" {
			id, _ := strconv.Atoi(bus)
			diff := id - (timestamp % id)
			if diff < candidateTime {
				candidateTime = diff
				candidateID = id
			}
		}
	}

	t.Log("Result:", candidateTime*candidateID) // 2845
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Test_part2(t *testing.T) {
	buses := map[int]int{}

	for i, bus := range strings.Split(input[1], ",") {
		if bus != "x" {
			id, _ := strconv.Atoi(bus)
			buses[id] = i
		}
	}

	t.Log(buses)

	step := 1
	timestamp := 0
	for k, v := range buses {
		id := k
		offset := v
		for {
			if (timestamp+offset)%id == 0 {
				break
			}
			timestamp += step
		}
		step = lcm(step, id)
	}

	t.Log("Result", timestamp) // 487905974205117
}
