package day18

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input      = utils.ReadStrings("day18_input.txt")
	precedence map[string]int
)

func rpn(tokens []string) []string {
	output := make([]string, 0)
	operator := make([]string, 0)
	for _, token := range tokens {
		switch token {
		case "+", "*":
			for len(operator) > 0 {
				op := operator[len(operator)-1]
				if op != "(" && precedence[op] >= precedence[token] {
					output = append(output, op)
					operator = operator[:len(operator)-1]
				} else {
					break
				}
			}
			operator = append(operator, token)
		case "(":
			operator = append(operator, token)
		case ")":
			for {
				op := operator[len(operator)-1]
				if op == "(" {
					operator = operator[:len(operator)-1]
					break
				}
				output = append(output, op)
				operator = operator[:len(operator)-1]
			}
		default:
			output = append(output, token)
		}
	}
	for i := len(operator) - 1; i >= 0; i-- {
		output = append(output, operator[i])
	}
	return output
}

func eval(expr []string) int {
	stack := make([]int, 0)
	for _, token := range expr {
		switch token {
		case "+":
			res := stack[len(stack)-1] + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case "*":
			res := stack[len(stack)-1] * stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		default:
			operand, _ := strconv.Atoi(token)
			stack = append(stack, operand)
		}
	}
	return stack[len(stack)-1]
}

func Test_part1(t *testing.T) {
	precedence = map[string]int{
		"+": 1,
		"*": 1,
	}
	sum := 0
	for _, line := range input {
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		sum += eval(rpn(strings.Fields(line)))
	}
	t.Log("Result:", sum)
}

func Test_part2(t *testing.T) {
	precedence = map[string]int{
		"+": 2,
		"*": 1,
	}
	sum := 0
	for _, line := range input {
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		sum += eval(rpn(strings.Fields(line)))
	}
	t.Log("Result:", sum)
}
