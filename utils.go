package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInts(name string) []int {
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
