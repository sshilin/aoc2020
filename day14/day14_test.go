package day14

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day14_input.txt")
)

func translate1(mask string, val int) int {
	for i := len(mask) - 1; i >= 0; i-- {
		pos := len(mask) - 1 - i
		switch mask[i] {
		case '0': // clear bit
			val &= ^(1 << pos)
		case '1': // set bit
			val |= (1 << pos)
		}
	}
	return val
}

func Test_part1(t *testing.T) {
	mem := map[int]int{}
	mask := ""

	for _, line := range input {
		parts := strings.Split(line, " = ")
		l := parts[0]
		r := parts[1]

		if l == "mask" {
			mask = r
		} else {
			addr, _ := strconv.Atoi(l[4 : len(l)-1])
			val, _ := strconv.Atoi(r)
			mem[addr] = translate1(mask, val)
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}

	t.Log("Result:", sum) // 9967721333886
}

func translate2(mask string, val int) []int {
	addr := make([]int, int(math.Pow(2, float64(strings.Count(mask, "X")))))

	for i := 0; i < len(addr); i++ {
		addr[i] = val
	}

	xNum := -1
	for i := len(mask) - 1; i >= 0; i-- {
		pos := len(mask) - 1 - i
		switch mask[i] {
		case 'X':
			xNum++
			for j := 0; j < len(addr); j++ {
				if j&(1<<xNum) == 0 { // test bit
					addr[j] &= ^(1 << pos) // clear bit
				} else {
					addr[j] |= (1 << pos) // set bit
				}
			}
		case '1': // set bit
			for j := 0; j < len(addr); j++ {
				addr[j] |= (1 << pos)
			}
		}
	}

	return addr
}

func Test_part2(t *testing.T) {
	mem := map[int]int{}
	mask := ""

	for _, line := range input {
		parts := strings.Split(line, " = ")
		l := parts[0]
		r := parts[1]

		if l == "mask" {
			mask = r
		} else {
			addr, _ := strconv.Atoi(l[4 : len(l)-1])
			val, _ := strconv.Atoi(r)
			for _, a := range translate2(mask, addr) {
				mem[a] = val
			}
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}

	t.Log("Result:", sum) // 4355897790573
}
