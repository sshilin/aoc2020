package day4

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day4_input.txt")

	hgtRegexp = regexp.MustCompile(`^(\d+)(\w+)$`)
	hclRegexp = regexp.MustCompile(`^#[[:xdigit:]]{6}$`)
	eclRegexp = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRegexp = regexp.MustCompile(`^\d{9}$`)
)

type Validator func(Passport) bool
type Passport map[string]string

func (p Passport) validate(validator Validator) bool {
	return validator(p)
}

func parsePassports() []Passport {
	passports := make([]Passport, 0)

	passport := make(Passport)
	for _, line := range input {
		if line != "" {
			for _, field := range strings.Fields(line) {
				kv := strings.Split(field, ":")
				passport[kv[0]] = kv[1]
			}
		} else {
			passports = append(passports, passport)
			passport = make(Passport)
		}
	}
	passports = append(passports, passport)

	return passports
}

func validatorPart1(passport Passport) bool {
	if _, ok := passport["byr"]; !ok {
		return false
	}
	if _, ok := passport["iyr"]; !ok {
		return false
	}
	if _, ok := passport["eyr"]; !ok {
		return false
	}
	if _, ok := passport["hgt"]; !ok {
		return false
	}
	if _, ok := passport["hcl"]; !ok {
		return false
	}
	if _, ok := passport["ecl"]; !ok {
		return false
	}
	if _, ok := passport["pid"]; !ok {
		return false
	}

	return true
}

func validatorPart2(passport Passport) bool {
	if val, ok := passport["byr"]; ok {
		v, _ := strconv.Atoi(val)
		if v < 1920 || v > 2002 {
			return false
		}
	} else {
		return false
	}

	if val, ok := passport["iyr"]; ok {
		v, _ := strconv.Atoi(val)
		if v < 2010 || v > 2020 {
			return false
		}
	} else {
		return false
	}

	if val, ok := passport["eyr"]; ok {
		v, _ := strconv.Atoi(val)
		if v < 2020 || v > 2030 {
			return false
		}
	} else {
		return false
	}

	if val, ok := passport["hgt"]; ok {
		res := hgtRegexp.FindStringSubmatch(val)
		if res[2] == "cm" {
			v, _ := strconv.Atoi(res[1])
			if v < 150 || v > 193 {
				return false
			}
		} else if res[2] == "in" {
			v, _ := strconv.Atoi(res[1])
			if v < 59 || v > 76 {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}

	if val, ok := passport["hcl"]; ok {
		if !hclRegexp.MatchString(val) {
			return false
		}
	} else {
		return false
	}

	if val, ok := passport["ecl"]; ok {
		if !eclRegexp.MatchString(val) {
			return false
		}
	} else {
		return false
	}

	if val, ok := passport["pid"]; ok {
		if !pidRegexp.MatchString(val) {
			return false
		}
	} else {
		return false
	}

	return true
}

func Test_part1(t *testing.T) {
	valid := 0
	for _, passport := range parsePassports() {
		if passport.validate(validatorPart1) {
			valid++
		}
	}
	t.Log("Result:", valid) // 190
}

func Test_part2(t *testing.T) {
	valid := 0
	for _, passport := range parsePassports() {
		if passport.validate(validatorPart2) {
			valid++
		}
	}
	t.Log("Result:", valid) // 121
}
