package day22

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/sshilin/aoc2020/utils"
)

var (
	input = utils.ReadStrings("day22_input.txt")
)

type queue []int

func (q *queue) pop() int {
	n := (*q)[0]
	*q = (*q)[1:len((*q))]
	return n
}

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) equals(b queue) bool {
	if len(*q) != len(b) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if (*q)[i] != b[i] {
			return false
		}
	}
	return true
}
func Test_part1(t *testing.T) {
	sections := utils.SplitBySections(input)

	player1 := make(queue, 0)
	player2 := make(queue, 0)

	for i := 1; i < len(sections[0]); i++ {
		card, _ := strconv.Atoi(sections[0][i])
		player1.push(card)
	}

	for i := 1; i < len(sections[1]); i++ {
		card, _ := strconv.Atoi(sections[1][i])
		player2.push(card)
	}

	round := 0
	for len(player1) > 0 && len(player2) > 0 {
		p1 := player1.pop()
		p2 := player2.pop()
		if p1 > p2 {
			player1.push(p1)
			player1.push(p2)
		} else {
			player2.push(p2)
			player2.push(p1)
		}
		round++
	}

	var winner queue

	if len(player1) == 0 {
		winner = player2
	} else {
		winner = player1
	}

	m := len(winner)
	res := 0
	for _, c := range winner {
		res += c * m
		m--
	}

	t.Log("Result:", res)
}

func play(player1, player2 queue) int {
	p1 := make(queue, len(player1))
	p2 := make(queue, len(player2))

	copy(p1, player1)
	copy(p2, player2)

	prevRoundsDeck1 := map[int]queue{}
	prevRoundsDeck2 := map[int]queue{}

	var winnerDeck queue

	winner := -1
	round := 0
	for len(p1) > 0 && len(p2) > 0 {
		// ---------- write round ----------
		prevDeck1 := make(queue, len(p1))
		prevDeck2 := make(queue, len(p2))

		copy(prevDeck1, p1)
		copy(prevDeck2, p2)

		prevRoundsDeck1[round] = prevDeck1
		prevRoundsDeck2[round] = prevDeck2
		// ---------- write round ----------

		// ---------- check for termination ----------
		for i := 0; i < round; i++ {
			if p1.equals(prevRoundsDeck1[i]) || p2.equals(prevRoundsDeck2[i]) {
				return 1
			}
		}
		// ---------- check for termination ----------

		c1 := p1.pop()
		c2 := p2.pop()

		if c1 <= len(p1) && c2 <= len(p2) {
			if play(p1[:c1], p2[:c2]) == 1 {
				p1.push(c1)
				p1.push(c2)
			} else {
				p2.push(c2)
				p2.push(c1)
			}
		} else {
			if c1 > c2 {
				p1.push(c1)
				p1.push(c2)
			} else {
				p2.push(c2)
				p2.push(c1)
			}
		}
		if len(p1) == 0 {
			winner = 2
			winnerDeck = p2
			break
		}
		if len(p2) == 0 {
			winner = 1
			winnerDeck = p1
			break
		}
		round++
	}

	m := len(winnerDeck)
	res := 0
	for _, c := range winnerDeck {
		res += c * m
		m--
	}
	fmt.Println("Result:", res)

	return winner
}

func Test_part2(t *testing.T) {
	sections := utils.SplitBySections(input)

	player1 := make(queue, 0)
	player2 := make(queue, 0)

	for i := 1; i < len(sections[0]); i++ {
		card, _ := strconv.Atoi(sections[0][i])
		player1.push(card)
	}

	for i := 1; i < len(sections[1]); i++ {
		card, _ := strconv.Atoi(sections[1][i])
		player2.push(card)
	}

	play(player1, player2)
}
