package day11

import (
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day11_input.txt")
)

type coord struct {
	x, y int
}

func countAdj1(c coord, grid map[coord]rune) int {
	count := 0
	for _, offset := range []coord{
		{-1, -1}, {+0, -1}, {+1, -1},
		{-1, +0}, {+1, +0},
		{-1, +1}, {+0, +1}, {+1, +1},
	} {
		if grid[coord{c.x + offset.x, c.y + offset.y}] == '#' {
			count++
		}
	}
	return count
}

func Test_part1(t *testing.T) {
	seats := map[coord]rune{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			seats[coord{x, y}] = rune(input[y][x])
		}
	}

	stop := false
	for !stop {
		stop = true
		newSeats := map[coord]rune{}
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[0]); x++ {
				c := coord{x, y}
				if seats[c] == 'L' && countAdj1(c, seats) == 0 {
					stop = false
					newSeats[c] = '#'
				} else if seats[c] == '#' && countAdj1(c, seats) >= 4 {
					stop = false
					newSeats[c] = 'L'
				} else {
					newSeats[c] = seats[c]
				}
			}
		}
		seats = newSeats
	}

	count := 0
	for _, v := range seats {
		if v == '#' {
			count++
		}
	}

	t.Log("Result:", count)
}

func countAdj2(c coord, grid map[coord]rune) int {
	count := 0
	for _, offset := range []coord{
		{-1, -1}, {+0, -1}, {+1, -1},
		{-1, +0}, {+1, +0},
		{-1, +1}, {+0, +1}, {+1, +1},
	} {
		z := c
		for {
			z.x += offset.x
			z.y += offset.y
			if grid[z] != '.' {
				if grid[z] == '#' {
					count++
				}
				break
			}
		}
	}
	return count
}

func Test_part2(t *testing.T) {
	seats := map[coord]rune{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			seats[coord{x, y}] = rune(input[y][x])
		}
	}

	stop := false
	for !stop {
		stop = true
		newSeats := map[coord]rune{}
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[0]); x++ {
				c := coord{x, y}
				if seats[c] == 'L' && countAdj2(c, seats) == 0 {
					stop = false
					newSeats[c] = '#'
				} else if seats[c] == '#' && countAdj2(c, seats) >= 5 {
					stop = false
					newSeats[c] = 'L'
				} else {
					newSeats[c] = seats[c]
				}
			}
		}
		seats = newSeats
	}

	count := 0
	for _, v := range seats {
		if v == '#' {
			count++
		}
	}

	t.Log("Result:", count)
}
