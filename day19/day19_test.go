package day19

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day19_input.txt")
	// input = utils.ReadStrings("test.txt")
	rules = map[string]rule{}
)

type rule struct {
	patterns [][]string
}

func resolveRule2(r string) string {
	fmt.Println("====> rule ", r)
	patterns := rules[r].patterns
	if len(patterns) == 1 && (string(patterns[0][0]) == "a" || string(patterns[0][0]) == "b") {
		return string(patterns[0][0])
	}

	var sb strings.Builder
	sb.WriteString("(")
	subPatterns := make([]string, 0)
	for _, pattern := range rules[r].patterns {
		var ssb strings.Builder
		for _, r := range pattern {
			ssb.WriteString(resolveRule2(r))
		}
		fmt.Println("ssb:", ssb.String())
		subPatterns = append(subPatterns, ssb.String())
	}
	sb.WriteString(strings.Join(subPatterns, "|"))
	sb.WriteString(")")
	return sb.String()
}

func Test_part1(t *testing.T) {
	sum := 0
	for _, line := range input {
		if line == "" {
			break
		}
		line = strings.ReplaceAll(line, "\"a\"", "a")
		line = strings.ReplaceAll(line, "\"b\"", "b")
		parts1 := strings.Split(line, ":")
		ruleName := parts1[0]
		parts2 := strings.Split(parts1[1], "|")
		parts3 := strings.Fields(parts2[0])
		parts4 := make([]string, 0)
		if len(parts2) == 2 {
			parts4 = strings.Fields(parts2[1])
		}
		patterns := make([][]string, 0)
		if len(parts3) > 0 {
			patterns = append(patterns, parts3)
		}
		if len(parts4) > 0 {
			patterns = append(patterns, parts4)
		}
		rules[ruleName] = rule{patterns}
	}

	res := resolveRule2("0")

	// res = strings.ReplaceAll(res, "|)", ")")
	t.Log(res)

	rg := regexp.MustCompile("^" + res + "$")
	for i := 141; i < len(input); i++ {
		msg := input[i]

		if rg.MatchString(msg) {
			sum++
		}
	}

	t.Log("Result:", sum)
}
