package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInts(name string) []int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	ints := make([]int, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			ints = append(ints, i)
		} else {
			log.Fatalln(err)
		}
	}
	return ints
}

func ReadStrings(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	lines := make([]string, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
