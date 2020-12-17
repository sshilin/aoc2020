package day17

import (
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day17_input.txt")
)

type coord struct {
	x, y, z int
}

type coord4 struct {
	x, y, z, w int
}

func neighbors(c coord, grid map[coord]byte, state byte) []coord {
	coords := make([]coord, 0)
	for _, x := range []int{1, 0, -1} {
		for _, y := range []int{1, 0, -1} {
			for _, z := range []int{1, 0, -1} {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				if grid[coord{x: c.x + x, y: c.y + y, z: c.z + z}] == state {
					coords = append(coords, coord{x: c.x + x, y: c.y + y, z: c.z + z})
				}
			}
		}
	}
	return coords
}

func neighbors4(c coord4, grid map[coord4]byte, state byte) []coord4 {
	coords := make([]coord4, 0)
	for _, x := range []int{1, 0, -1} {
		for _, y := range []int{1, 0, -1} {
			for _, z := range []int{1, 0, -1} {
				for _, w := range []int{1, 0, -1} {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					if grid[coord4{x: c.x + x, y: c.y + y, z: c.z + z, w: c.w + w}] == state {
						coords = append(coords, coord4{x: c.x + x, y: c.y + y, z: c.z + z, w: c.w + w})
					}
				}
			}
		}
	}
	return coords
}

func Test_part1(t *testing.T) {
	grid := map[coord]byte{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			if input[x][y] == '#' {
				grid[coord{x: x, y: y, z: 0}] = input[x][y]
			}
		}
	}

	for i := 0; i < 6; i++ {
		nextRound := map[coord]byte{}
		inactiveCellsToCheck := map[coord]bool{}

		for k, v := range grid {
			n := len(neighbors(k, grid, '#'))
			if v == '#' {
				if n == 2 || n == 3 {
					nextRound[k] = '#'
				}
				for _, ns := range neighbors(k, grid, 0) {
					inactiveCellsToCheck[ns] = true
				}
			}
		}

		for k := range inactiveCellsToCheck {
			n := len(neighbors(k, grid, '#'))
			if n == 3 {
				nextRound[k] = '#'
			}
		}

		grid = nextRound
	}

	count := 0
	for _, v := range grid {
		if v == '#' {
			count++
		}
	}

	t.Log("Result:", count) // 237
}

func Test_part2(t *testing.T) {
	grid := map[coord4]byte{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			if input[x][y] == '#' {
				grid[coord4{x: x, y: y, z: 0, w: 0}] = input[x][y]
			}
		}
	}

	for i := 0; i < 6; i++ {
		nextRound := map[coord4]byte{}
		inactiveCellsToCheck := map[coord4]bool{}

		for k, v := range grid {
			n := len(neighbors4(k, grid, '#'))
			if v == '#' {
				if n == 2 || n == 3 {
					nextRound[k] = '#'
				}
				for _, ns := range neighbors4(k, grid, 0) {
					inactiveCellsToCheck[ns] = true
				}
			}
		}

		for k := range inactiveCellsToCheck {
			if len(neighbors4(k, grid, '#')) == 3 {
				nextRound[k] = '#'
			}
		}

		grid = nextRound
	}

	count := 0
	for _, v := range grid {
		if v == '#' {
			count++
		}
	}

	t.Log("Result:", count) // 2448
}
