package day17

import (
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day17_input.txt")
)

type coord struct {
	x, y, z, w int
}

func neighbors(c coord, grid map[coord]byte, state byte, mode string) []coord {
	wRange := []int{0}
	if mode == "4d" {
		wRange = []int{1, 0, -1}
	}

	coords := make([]coord, 0)
	for _, x := range []int{1, 0, -1} {
		for _, y := range []int{1, 0, -1} {
			for _, z := range []int{1, 0, -1} {
				for _, w := range wRange {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					} else if grid[coord{x: c.x + x, y: c.y + y, z: c.z + z, w: c.w + w}] == state {
						coords = append(coords, coord{x: c.x + x, y: c.y + y, z: c.z + z, w: c.w + w})
					}
				}
			}
		}
	}
	return coords
}

func run(mode string) int {
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
			n := len(neighbors(k, grid, '#', mode))
			if v == '#' {
				if n == 2 || n == 3 {
					nextRound[k] = '#'
				}
				for _, ns := range neighbors(k, grid, 0, mode) {
					inactiveCellsToCheck[ns] = true
				}
			}
		}

		for k := range inactiveCellsToCheck {
			n := len(neighbors(k, grid, '#', mode))
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

	return count
}

func Test_part1(t *testing.T) {
	t.Log("Result:", run("3d")) // 237
}

func Test_part2(t *testing.T) {
	t.Log("Result:", run("4d")) // 2448
}
