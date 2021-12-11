package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day11/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

var adj [][]int = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func validPoint(i int, j int) bool {
	return i >= 0 && j >= 0 && i < 10 && j < 10
}

func part1(filename string) int {
	defer duration(track("part1"))

	d := data.GetData(filename)

	flashCount := 0
	var flashed [][]int
	var flashFunc func(i int, j int)
	flashFunc = func(i int, j int) {
		flashCount++
		flashed = append(flashed, []int{i, j})
		for _, delta := range adj {
			iDelta := i + delta[0]
			jDelta := j + delta[1]
			if validPoint(iDelta, jDelta) {
				d[iDelta][jDelta]++
				if d[iDelta][jDelta] == 10 {
					flashFunc(iDelta, jDelta)
				}
			}
		}
	}

	for round := 0; round < 100; round++ {
		for i := 0; i < len(d); i++ {
			for j := 0; j < len(d[i]); j++ {
				d[i][j]++
				if d[i][j] == 10 {
					flashFunc(i, j)
				}
			}
		}
		for _, point := range flashed {
			d[point[0]][point[1]] = 0
		}
		flashed = [][]int{}
	}

	return flashCount
}

func part2(filename string) int {
	defer duration(track("part2"))

	d := data.GetData(filename)

	var flashed [][]int
	var flashFunc func(i int, j int)
	flashFunc = func(i int, j int) {
		flashed = append(flashed, []int{i, j})
		for _, delta := range adj {
			iDelta := i + delta[0]
			jDelta := j + delta[1]
			if validPoint(iDelta, jDelta) {
				d[iDelta][jDelta]++
				if d[iDelta][jDelta] == 10 {
					flashFunc(iDelta, jDelta)
				}
			}
		}
	}

	round := 1
	fullCount := len(d) * len(d[0])
	for {
		for i := 0; i < len(d); i++ {
			for j := 0; j < len(d[i]); j++ {
				d[i][j]++
				if d[i][j] == 10 {
					flashFunc(i, j)
				}
			}
		}
		if len(flashed) == fullCount {
			return round
		}
		for _, point := range flashed {
			d[point[0]][point[1]] = 0
		}
		flashed = [][]int{}
		round++
	}
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
