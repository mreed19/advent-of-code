package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/mreed19/advent-of-code/2021/day10/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

type symbolKind uint

const (
	OPENER symbolKind = iota
	CLOSER
)

type symbol struct {
	kind  symbolKind
	val   int
	match rune
}

var symbolMap map[rune]*symbol = map[rune]*symbol{
	'{': {
		kind: OPENER,
		val:  3,
	},
	'[': {
		kind: OPENER,
		val:  2,
	},
	'(': {
		kind: OPENER,
		val:  1,
	},
	'<': {
		kind: OPENER,
		val:  4,
	},
	'}': {
		kind:  CLOSER,
		val:   1197,
		match: '{',
	},
	']': {
		kind:  CLOSER,
		val:   57,
		match: '[',
	},
	')': {
		kind:  CLOSER,
		val:   3,
		match: '(',
	},
	'>': {
		kind:  CLOSER,
		val:   25137,
		match: '<',
	},
}

func part1(filename string) int {
	defer duration(track("part1"))

	lines := data.GetData(filename)

	errScore := 0

	for _, line := range lines {
		var stack []rune
		for _, char := range line {
			symbolObj := symbolMap[char]
			if symbolObj.kind == OPENER {
				stack = append(stack, char)
			} else {
				if symbolObj.match == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					errScore += symbolObj.val
					break
				}
			}
		}

	}

	return errScore
}

func part2(filename string) int {
	defer duration(track("part2"))

	lines := data.GetData(filename)

	var scores []int

	for _, line := range lines {
		var stack []rune
		score := 0
		skipLine := false
		for _, char := range line {
			symbolObj := symbolMap[char]
			if symbolObj.kind == OPENER {
				stack = append(stack, char)
			} else {
				if symbolObj.match == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					skipLine = true
					break
				}
			}
		}

		if skipLine {
			continue
		}

		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + symbolMap[stack[i]].val
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
