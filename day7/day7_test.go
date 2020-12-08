package day7

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day7_input.txt")
	left  = regexp.MustCompile(`^(\w+ \w+) bag`)
	right = regexp.MustCompile(`(\d+) (\w+ \w+) bag`)
)

func Test_part1(t *testing.T) {
	bags := map[string][]string{}

	for _, line := range input {
		k := strings.Split(line, "contain")[0]
		k = strings.TrimSpace(k)
		k = k[:len(k)-1]
		bags[k] = []string{}

		v := strings.Split(line, "contain")[1]

		for _, s := range strings.Split(v, ",") {

			s = strings.TrimSpace(s)
			s = strings.Replace(s, ".", "", 1)
			s = strings.Replace(s, "bags", "bag", 1)
			s = s[2:]

			bags[k] = append(bags[k], s)
		}
	}

	result := 0
	contains := make([]string, 0)

	delete(bags, "shiny gold bag")

	for k, v := range bags {
		if isContain("shiny gold bag", v) {
			contains = append(contains, k)
			result++
		}
	}

	for i := 0; i < 20; i++ {
		for _, c := range contains {
			for k, v := range bags {
				if isContain(c, v) && !isContain(k, contains) {
					contains = append(contains, k)
					result++
				}
			}
		}
	}

	t.Log("Result:", result) // 372
}

func isContain(val string, strs []string) bool {
	for _, s := range strs {
		if s == val {
			return true
		}
	}
	return false
}

type bag struct {
	name     string
	quantity int
}

type bags map[string][]bag

func Test_part2(t *testing.T) {
	bags := map[string][]bag{}

	for _, line := range input {
		k := left.FindString(line)
		bags[k] = make([]bag, 0)

		v := strings.Split(line, "contain")[1]
		v = strings.TrimSpace(v)
		v = strings.TrimFunc(v, func(r rune) bool {
			return r == '.'
		})

		for _, s := range strings.Split(v, ",") {
			s = strings.TrimSpace(s)
			s = strings.Replace(s, ".", "", 1)
			s = strings.Replace(s, "bags", "bag", 1)

			if s != "no other bag" {
				q, _ := strconv.Atoi(s[:1])
				bags[k] = append(bags[k], bag{name: s[2:], quantity: q})
			}
		}
	}

	t.Log("Result:", traverse("shiny gold bag", bags, t)) // 8015
}

func traverse(currentBugName string, bs bags, t *testing.T) int {
	count := 0
	subBags := bs[currentBugName]
	for _, b := range subBags {
		count += b.quantity
		count += traverse(b.name, bs, t) * b.quantity
	}
	return count
}
