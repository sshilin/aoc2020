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
	rules = map[string]rule{}
	part2 = false
)

type rule struct {
	patterns [][]string
}

func resolve(name string) string {
	if rules[name].patterns[0][0] == "a" || rules[name].patterns[0][0] == "b" {
		return rules[name].patterns[0][0]
	}
	var sb strings.Builder

	sb.WriteString("(")

	if name == "8" && part2 {
		sb.WriteString(fmt.Sprintf("%s+", resolve("42")))
	} else if name == "11" && part2 {
		rule11 := make([]string, 0)
		for i := 1; i <= 5; i++ {
			rule11 = append(rule11, fmt.Sprintf("%s{%d}%s{%d}", resolve("42"), i, resolve("31"), i))
		}
		sb.WriteString(strings.Join(rule11, "|"))
	} else {
		subPatterns := make([]string, 0)
		for _, pattern := range rules[name].patterns {
			var ssb strings.Builder
			for _, name := range pattern {
				ssb.WriteString(resolve(name))
			}
			subPatterns = append(subPatterns, ssb.String())
		}
		sb.WriteString(strings.Join(subPatterns, "|"))
	}

	sb.WriteString(")")

	return sb.String()
}

func splitBySections() [][]string {
	sections := make([][]string, 0)
	section := make([]string, 0)
	for i := 0; i < len(input); i++ {
		if input[i] != "" {
			section = append(section, input[i])
		} else {
			sections = append(sections, section)
			section = make([]string, 0)
		}
	}
	sections = append(sections, section)
	return sections
}

func countValidMessages() int {
	sum := 0
	sections := splitBySections()
	for _, line := range sections[0] {
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

	rg := regexp.MustCompile("^" + resolve("0") + "$")

	for _, msg := range sections[1] {
		if rg.MatchString(msg) {
			sum++
		}
	}
	return sum
}
func Test_part1(t *testing.T) {
	t.Log("Result:", countValidMessages())
}

func Test_part2(t *testing.T) {
	part2 = true
	t.Log("Result:", countValidMessages())
}
