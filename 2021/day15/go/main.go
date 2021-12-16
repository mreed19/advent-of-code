package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mreed19/advent-of-code/2021/day15/go/data"
)

func main() {
	// grid := data.GetData("data/input-sample.txt")
	grid := data.GetData("data/input.txt")

	fmt.Println("Part 1 Solution:", part1(grid))
	fmt.Println("Part 2 Solution:", part2(grid))
}

func part1(grid [][]int) int {
	defer duration(track("part1"))

	return shortestPath(grid)
}

func part2(grid [][]int) int {
	defer duration(track("part2"))

	rows := len(grid)
	cols := len(grid[0])

	for count := 0; count < 4; count++ {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				grid[i] = append(grid[i], (grid[i][j]+count)%9+1)
			}
		}
	}

	for count := 0; count < 4; count++ {
		for i := 0; i < rows; i++ {
			var newRow []int
			for j := 0; j < 5*cols; j++ {
				newRow = append(newRow, (grid[i][j]+count)%9+1)
			}
			grid = append(grid, newRow)
		}
	}

	return shortestPath(grid)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var directions = [][]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func isValidMove(i, j, rows, cols int) bool {
	return i >= 0 && j >= 0 && i < rows && j < cols
}

func findShortest(distances [][]int, visited [][]bool) (int, int) {
	shortI := -1
	shortJ := -1

	for i := 0; i < len(distances); i++ {
		for j := 0; j < len(distances[i]); j++ {
			isShortest := (shortI == -1 && shortJ == -1) || distances[i][j] < distances[shortI][shortJ]
			if isShortest && !visited[i][j] {
				shortI = i
				shortJ = j
			}
		}
	}
	return shortI, shortJ
}

func shortestPath(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	distances := make([][]int, rows)
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		distances[i] = make([]int, cols)
		visited[i] = make([]bool, cols)
		for j := 0; j < len(grid[i]); j++ {
			distances[i][j] = math.MaxInt
		}
	}

	distances[0][0] = 0
	i := 0
	j := 0

	for i != -1 && j != -1 {
		visited[i][j] = true

		if i == rows-1 && j == cols-1 {
			break
		}

		for _, delta := range directions {
			di := i + delta[0]
			dj := j + delta[1]
			if isValidMove(di, dj, rows, cols) && !visited[di][dj] {
				distances[di][dj] = minInt(distances[di][dj], distances[i][j]+grid[di][dj])
			}
		}
		i, j = findShortest(distances, visited)
	}

	return distances[rows-1][cols-1]
}
