package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// fmt.Println("Part 1 Solution:", part1("input-sample.txt"))
	fmt.Println("Part 1 Solution:", part1("input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	grid := GetData(filename)

	steps := 1

	for {
		moves := grid.MoveEast() + grid.MoveSouth()
		if moves == 0 {
			break
		}
		steps++
	}

	return steps
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
