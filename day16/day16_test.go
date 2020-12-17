package day16

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input           = utils.ReadStrings("day16_input.txt")
	validatorRegexp = regexp.MustCompile(`([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)
)

type validator struct {
	firstFrom, firstTo, secondFrom, secondTo int
}

func parseTicket(ticket string) []int {
	t := make([]int, 0)
	for _, f := range strings.Split(ticket, ",") {
		field, _ := strconv.Atoi(f)
		t = append(t, field)
	}
	return t
}

func Test_part1(t *testing.T) {
	sum := 0
	validators := map[string]validator{}

	for i := 0; i < 20; i++ {
		groups := validatorRegexp.FindStringSubmatch(input[i])
		firstFrom, _ := strconv.Atoi(groups[2])
		firstTo, _ := strconv.Atoi(groups[3])
		secondFrom, _ := strconv.Atoi(groups[4])
		secondTo, _ := strconv.Atoi(groups[5])
		validators[groups[1]] = validator{firstFrom, firstTo, secondFrom, secondTo}
	}

	for i := 25; i < len(input); i++ {
		for _, f := range strings.Split(input[i], ",") {
			field, _ := strconv.Atoi(f)
			isValid := false
			for _, v := range validators {
				if (field >= v.firstFrom && field <= v.firstTo) || (field >= v.secondFrom && field <= v.secondTo) {
					isValid = true
					break
				}
			}
			if !isValid {
				sum += field
			}
		}
	}
	t.Log("Result:", sum) // 21071
}

func Test_part0(t *testing.T) {
	sum := 0
	validators := map[string]validator{}

	for i := 0; i < 20; i++ {
		groups := validatorRegexp.FindStringSubmatch(input[i])
		firstFrom, _ := strconv.Atoi(groups[2])
		firstTo, _ := strconv.Atoi(groups[3])
		secondFrom, _ := strconv.Atoi(groups[4])
		secondTo, _ := strconv.Atoi(groups[5])
		validators[groups[1]] = validator{firstFrom, firstTo, secondFrom, secondTo}
	}

	yourTicket := parseTicket(input[22])

	validTickets := make([][]int, 0)
	validTickets = append(validTickets, yourTicket)

	for i := 25; i < len(input); i++ {
		validTicket := true
		ticket := parseTicket(input[i])

		for _, field := range ticket {
			valid := false

			for _, v := range validators {
				if field >= v.firstFrom && field <= v.firstTo {
					valid = true
					break
				}
				if field >= v.secondFrom && field <= v.secondTo {
					valid = true
					break
				}
			}

			if !valid {
				sum += field
				validTicket = false
				break
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	fieldPositions := map[string]int{}

	for i := 0; i < len(validators); i++ {
		for k, v := range validators {
			fitsCount := 0
			fitsPos := 0
			for i := 0; i < 20; i++ {
				fits := true
				for _, ticket := range validTickets {
					if (ticket[i] < v.firstFrom || ticket[i] > v.firstTo) && (ticket[i] < v.secondFrom || ticket[i] > v.secondTo) {
						fits = false
						break
					}
				}
				for _, v := range fieldPositions {
					if i == v {
						fits = false
					}
				}
				if fits {
					fitsCount++
					fitsPos = i
				}
			}
			if fitsCount == 1 {
				fieldPositions[k] = fitsPos
			}
		}
	}

	result := 1
	for k, v := range fieldPositions {
		if strings.HasPrefix(k, "departure") {
			result *= yourTicket[v]
		}
	}
	t.Log("Result:", result) // 3429967441937
}
