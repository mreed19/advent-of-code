package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day18/go/data"
)

func main() {
	defer duration(track("Total"))
	numbers := data.GetData("data/input.txt")
	var numbers2 []*data.SnailNumber
	for _, num := range numbers {
		numbers2 = append(numbers2, num.Copy())
	}

	fmt.Println("Part 1 Solution:", part1(numbers))
	fmt.Println("Part 2 Solution:", part2(numbers2))
}

func part1(numbers []*data.SnailNumber) int {
	defer duration(track("part1"))

	curr := numbers[0]
	for i := 1; i < len(numbers); i++ {
		curr = data.AddSnailNumbers(curr, numbers[i])
	}

	return curr.Magnitude()
}

func part2(numbers []*data.SnailNumber) int {
	defer duration(track("part2"))

	maxMagnitude := 0

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				mag := data.AddSnailNumbers(numbers[i].Copy(), numbers[j].Copy()).Magnitude()
				if mag > maxMagnitude {
					maxMagnitude = mag
				}
			}
		}
	}

	return maxMagnitude
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
