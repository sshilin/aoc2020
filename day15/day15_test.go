package day15

import (
	"fmt"
	"testing"
)

type pair struct {
	lastTurn, prevTurn int
}

func Test_part1and2(t *testing.T) {
	for _, turns := range []int{2020, 30000000} {
		t.Run(fmt.Sprintf("%d turns", turns), func(t *testing.T) {
			mem := map[int]pair{
				1:  {1, -1},
				0:  {2, -1},
				15: {3, -1},
				2:  {4, -1},
				10: {5, -1},
				13: {6, -1},
			}
			last := 13
			isNew := true
			for i := len(mem) + 1; i <= turns; i++ {
				if isNew {
					last = 0
					isNew = false
					mem[0] = pair{i, mem[0].lastTurn}
				} else {
					last = mem[last].lastTurn - mem[last].prevTurn
					if _, ok := mem[last]; ok {
						mem[last] = pair{i, mem[last].lastTurn}
						isNew = false
					} else {
						mem[last] = pair{i, -1}
						isNew = true
					}
				}
			}
			t.Log("Result:", last) // 211, 2159626
		})
	}
}
