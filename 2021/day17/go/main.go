package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mreed19/advent-of-code/2021/day17/go/data"
)

func main() {
	// d := data.GetData("data/input-sample.txt")
	d := data.GetData("data/input.txt")

	fmt.Println("Part 1 Solution:", part1(d))
	fmt.Println("Part 2 Solution:", part2(d))
}

func part1(t *data.Target) int {
	defer duration(track("part1"))

	maxAbsY := int(math.Max(math.Abs(float64(t.MinY)), math.Abs(float64(t.MaxY))))

	for x := 1; x <= t.MaxX; x++ {
		for y := maxAbsY; y >= -maxAbsY; y-- {
			currX := 0
			currY := 0
			vx := x
			vy := y

			for !t.PastTargetArea(currY) {
				if t.InTargetArea(currX, currY) {
					return maxHeight(y)
				}
				currX += vx
				currY += vy
				if vx != 0 {
					vx--
				}
				vy--
			}

			if currX < t.MinX {
				break
			}
		}
	}

	return 0
}

func part2(t *data.Target) int {
	defer duration(track("part2"))

	maxAbsY := int(math.Max(math.Abs(float64(t.MinY)), math.Abs(float64(t.MaxY))))
	hitCount := 0

	for x := 1; x <= t.MaxX; x++ {
		for y := maxAbsY; y >= -maxAbsY; y-- {
			currX := 0
			currY := 0
			vx := x
			vy := y

			for !t.PastTargetArea(currY) {
				if t.InTargetArea(currX, currY) {
					hitCount++
					break
				}
				currX += vx
				currY += vy
				if vx != 0 {
					vx--
				}
				vy--
			}

			if currX < t.MinX {
				break
			}
		}
	}

	return hitCount
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func maxHeight(y int) int {
	return (y * (y + 1)) / 2
}
