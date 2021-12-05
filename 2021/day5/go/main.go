package main

import (
	"encoding/json"
	"fmt"

	"github.com/mreed19/advent-of-code/2021/day5/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1())
	fmt.Println("Part 2 Solution:", part2())
}

func part1() int {
	vents := data.ReadInput("data/input.txt", data.NewHZVent)
	return getDangerCount(vents)
}

func part2() int {
	vents := data.ReadInput("data/input.txt", data.NewVent)
	return getDangerCount(vents)
}

func getDangerCount(vents []*data.Vent) int {
	ventMap := make(map[int]map[int]int)
	dangerCount := 0

	for _, vent := range vents {
		xChange := getStep(vent.Point1.X, vent.Point2.X)
		yChange := getStep(vent.Point1.Y, vent.Point2.Y)

		x := vent.Point1.X
		y := vent.Point1.Y
		for {
			if ventMap[x] == nil {
				ventMap[x] = make(map[int]int)
			}

			if count, ok := ventMap[x][y]; ok {
				if count == 1 {
					dangerCount++
				}
				ventMap[x][y] = count + 1
			} else {
				ventMap[x][y] = 1
			}

			if x == vent.Point2.X && y == vent.Point2.Y {
				break
			}

			x += xChange
			y += yChange
		}
	}

	return dangerCount
}

func printVents(vents []*data.Vent) {
	jsonBytes, _ := json.MarshalIndent(vents, "", "  ")
	fmt.Println(string(jsonBytes))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getStep(start, end int) int {
	if start == end {
		return 0
	} else if start > end {
		return -1
	}
	return 1
}
