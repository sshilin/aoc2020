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

func neighbors(c coord, grid map[coord]bool, mode string) ([]coord, []coord) {
	wRange := []int{0}
	if mode == "4d" {
		wRange = []int{1, 0, -1}
	}
	active := make([]coord, 0)
	inactive := make([]coord, 0)
	for _, x := range []int{1, 0, -1} {
		for _, y := range []int{1, 0, -1} {
			for _, z := range []int{1, 0, -1} {
				for _, w := range wRange {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					nc := coord{c.x + x, c.y + y, c.z + z, c.w + w}
					if _, ok := grid[nc]; ok {
						active = append(active, nc)
					} else {
						inactive = append(inactive, nc)
					}
				}
			}
		}
	}
	return active, inactive
}

func run(mode string) int {
	grid := map[coord]bool{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == '#' {
				grid[coord{x, y, 0, 0}] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		nextRound := map[coord]bool{}
		for ac := range grid {
			active, inactive := neighbors(ac, grid, mode)
			if len(active) == 2 || len(active) == 3 {
				nextRound[ac] = true
			}
			for _, ic := range inactive {
				active, _ := neighbors(ic, grid, mode)
				if len(active) == 3 {
					nextRound[ic] = true
				}
			}
		}
		grid = nextRound
	}

	return len(grid)
}

func Test_part1(t *testing.T) {
	t.Log("Result:", run("3d")) // 237
}

func Test_part2(t *testing.T) {
	t.Log("Result:", run("4d")) // 2448
}
