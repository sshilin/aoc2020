package day8

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day8_input.txt")
)

type instruction struct {
	code string
	arg  int
}

func Test_part1(t *testing.T) {
	instrCounter := map[int]int{}
	program := make([]*instruction, 0)
	acc := 0
	ip := 0

	for _, line := range input {
		fields := strings.Fields(line)
		arg, _ := strconv.Atoi(fields[1])
		program = append(program, &instruction{fields[0], arg})
	}

	for {
		instrCounter[ip]++
		inst := program[ip]

		switch inst.code {
		case "nop":
			ip++
		case "acc":
			acc += inst.arg
			ip++
		case "jmp":
			ip += inst.arg
		}
		if instrCounter[ip] != 0 {
			break
		}
	}

	t.Log("Result:", acc) // 1675
}

func Test_part2(t *testing.T) {
	program := make([]*instruction, 0)
	jmps, nops := 0, 0

	for _, line := range input {
		fields := strings.Fields(line)
		arg, _ := strconv.Atoi(fields[1])
		program = append(program, &instruction{fields[0], arg})
		if fields[0] == "jmp" {
			jmps++
		} else if fields[0] == "nop" {
			nops++
		}
	}

	patchThisIP := -1
	for i := 0; i < jmps+nops; i++ {
		instrCounter := map[int]int{}
		acc := 0
		ip := 0

		for i, instr := range program {
			if (instr.code == "jmp" || instr.code == "nop") && i > patchThisIP {
				patchThisIP = i
				break
			}
		}

		for {
			instrCounter[ip]++
			inst := program[ip]

			if ip == patchThisIP {
				if inst.code == "jmp" {
					inst = &instruction{"nop", inst.arg}
				} else if inst.code == "nop" {
					inst = &instruction{"jmp", inst.arg}
				}
			}

			switch inst.code {
			case "nop":
				ip++
			case "acc":
				acc += inst.arg
				ip++
			case "jmp":
				ip += inst.arg
			}
			if instrCounter[ip] != 0 {
				break
			}
			if ip >= len(program) {
				t.Log("Result:", acc) // 1532
				return
			}
		}
	}
}
