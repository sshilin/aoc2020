package day12

import (
	"math"
	"strconv"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day12_input.txt")
)

type coord struct {
	east, west, north, south int
	direction                int
}

// east = 0
// south = 1
// west = 2
// north = 3

func Test_part1(t *testing.T) {
	steps := make([]coord, 0)
	steps = append(steps, coord{east: 0, west: 0, north: 0, south: 0, direction: 0})

	for _, line := range input {
		action := line[:1]
		arg, _ := strconv.Atoi(line[1:])

		lastStep := steps[len(steps)-1]
		switch action {
		case "F":
			if lastStep.direction == 0 {
				steps = append(steps, coord{east: lastStep.east + arg, west: lastStep.west, north: lastStep.north, south: lastStep.south, direction: lastStep.direction})
			} else if lastStep.direction == 1 {
				steps = append(steps, coord{east: lastStep.east, west: lastStep.west, north: lastStep.north, south: lastStep.south + arg, direction: lastStep.direction})
			} else if lastStep.direction == 2 {
				steps = append(steps, coord{east: lastStep.east, west: lastStep.west + arg, north: lastStep.north, south: lastStep.south, direction: lastStep.direction})
			} else if lastStep.direction == 3 {
				steps = append(steps, coord{east: lastStep.east, west: lastStep.west, north: lastStep.north + arg, south: lastStep.south, direction: lastStep.direction})
			}
		case "N":
			steps = append(steps, coord{east: lastStep.east, west: lastStep.west, north: lastStep.north + arg, south: lastStep.south, direction: lastStep.direction})
		case "S":
			steps = append(steps, coord{east: lastStep.east, west: lastStep.west, north: lastStep.north, south: lastStep.south + arg, direction: lastStep.direction})
		case "E":
			steps = append(steps, coord{east: lastStep.east + arg, west: lastStep.west, north: lastStep.north, south: lastStep.south, direction: lastStep.direction})
		case "W":
			steps = append(steps, coord{east: lastStep.east, west: lastStep.west + arg, north: lastStep.north, south: lastStep.south, direction: lastStep.direction})
		case "R":
			newDirection := (lastStep.direction + (arg / 90)) % 4
			steps = append(steps, coord{east: lastStep.east, west: lastStep.west, north: lastStep.north, south: lastStep.south, direction: newDirection})
		case "L":
			newDirection := (lastStep.direction - (arg / 90) + 4) % 4
			steps = append(steps, coord{east: lastStep.east, west: lastStep.west, north: lastStep.north, south: lastStep.south, direction: newDirection})
		}
	}

	lastStep := steps[len(steps)-1]
	ew := lastStep.east - lastStep.west
	ne := lastStep.north - lastStep.south
	res := math.Abs(float64(ew)) + math.Abs(float64(ne))

	t.Log("Result:", res) // 879
}

func Test_part2(t *testing.T) {
	steps := make([]coord, 1)
	waypoint := coord{east: 10, west: 0, north: 1, south: 0}

	for _, line := range input {
		action := line[:1]
		arg, _ := strconv.Atoi(line[1:])

		lastStep := steps[len(steps)-1]
		switch action {
		case "F":
			steps = append(steps, coord{
				east:  lastStep.east + waypoint.east*arg,
				west:  lastStep.west + waypoint.west*arg,
				north: lastStep.north + waypoint.north*arg,
				south: lastStep.south + waypoint.south*arg,
			})
		case "N":
			waypoint.north += arg
		case "S":
			waypoint.south += arg
		case "E":
			waypoint.east += arg
		case "W":
			waypoint.west += arg
		case "R":
			for i := 0; i < arg/90; i++ {
				tmp := waypoint.north
				waypoint.north = waypoint.west
				waypoint.west = waypoint.south
				waypoint.south = waypoint.east
				waypoint.east = tmp
			}
		case "L":
			for i := 0; i < arg/90; i++ {
				tmp := waypoint.north
				waypoint.north = waypoint.east
				waypoint.east = waypoint.south
				waypoint.south = waypoint.west
				waypoint.west = tmp
			}
		}
	}

	lastStep := steps[len(steps)-1]
	ew := lastStep.east - lastStep.west
	ne := lastStep.north - lastStep.south
	res := math.Abs(float64(ew)) + math.Abs(float64(ne))

	t.Log("Result:", res) // 18107
}
