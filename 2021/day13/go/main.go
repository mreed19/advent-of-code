package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day13/go/data"
)

func main() {
	// fmt.Println("Part 1 Solution:", part1("data/input-sample.txt"))
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	// fmt.Println("Part 2 Solution:")
	// part2("data/input-sample.txt")
	fmt.Println("Part 2 Solution:")
	part2("data/input.txt")
}

func part1(filename string) int {
	defer duration(track("part1"))

	points, folds := data.GetData(filename)

	// Only first fold for part 1
	folds = folds[:1]

	for _, fold := range folds {
		if fold.Direction == "x" {
			for x := range points {
				if x > fold.FoldLine {
					newX := fold.FoldLine - (x - fold.FoldLine)
					for y := range points[x] {
						if points[newX] == nil {
							points[newX] = make(map[int]bool)
						}
						points[newX][y] = true
					}
					points[x] = nil
				}
			}
		} else {
			for x := range points {
				for y := range points[x] {
					if y > fold.FoldLine {
						newY := fold.FoldLine - (y - fold.FoldLine)
						points[x][newY] = true
						delete(points[x], y)
					}
				}
			}
		}
	}

	count := 0
	for x := range points {
		for range points[x] {
			count++
		}
	}

	return count
}

func part2(filename string) {
	defer duration(track("part2"))

	points, folds := data.GetData(filename)

	for _, fold := range folds {
		if fold.Direction == "x" {
			for x := range points {
				if x > fold.FoldLine {
					newX := fold.FoldLine - (x - fold.FoldLine)
					for y := range points[x] {
						if points[newX] == nil {
							points[newX] = make(map[int]bool)
						}
						points[newX][y] = true
					}
					points[x] = nil
				}
			}
		} else {
			for x := range points {
				for y := range points[x] {
					if y > fold.FoldLine {
						newY := fold.FoldLine - (y - fold.FoldLine)
						points[x][newY] = true
						delete(points[x], y)
					}
				}
			}
		}
	}

	maxX := 0
	maxY := 0
	for x := range points {
		for y := range points[x] {
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	grid := make([][]string, maxX+1)
	for i := 0; i <= maxX; i++ {
		grid[i] = make([]string, maxY+1)
		for j := 0; j <= maxY; j++ {
			if points[i][j] {
				grid[i][j] = "X"
				fmt.Print("X")
			} else {
				grid[i][j] = "-"
				fmt.Print("-")
			}
		}
		fmt.Printf("\n")
	}

	grid = transpose(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Print("\n")
	}
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
