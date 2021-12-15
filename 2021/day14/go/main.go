package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mreed19/advent-of-code/2021/day14/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	return runPolymerization(filename, 10)
}

func part2(filename string) int {
	defer duration(track("part2"))

	return runPolymerization(filename, 40)
}

func runPolymerization(filename string, n int) int {
	polymers := data.GetData(filename)

	for i := 0; i < n; i++ {
		for _, polymer := range polymers {
			if polymer.Count > 0 {
				for _, next := range polymer.Next {
					polymers[next].NextCount += polymer.Count
				}
			}
		}
		for _, polymer := range polymers {
			polymer.Count = polymer.NextCount
			polymer.NextCount = 0
		}
	}

	counts := make(map[rune]int)
	for key, polymer := range polymers {
		for _, char := range key {
			counts[char] += polymer.Count
		}
	}

	minCount := math.MaxInt
	maxCount := math.MinInt
	for _, count := range counts {
		if count%2 == 1 {
			count = 1 + ((count - 1) / 2)
		} else {
			count = count / 2
		}

		if count < minCount {
			minCount = count
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount - minCount
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
