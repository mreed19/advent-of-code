package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/mreed19/advent-of-code/2021/day9/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	smokeMap := data.GetData(filename)
	lowPoints := findLowPoints(smokeMap)

	dangerRating := 0
	for _, lowPoint := range lowPoints {
		dangerRating += smokeMap[lowPoint.X][lowPoint.Y] + 1
	}

	return dangerRating
}

func part2(filename string) int {
	defer duration(track("part2"))

	smokeMap := data.GetData(filename)
	lowPoints := findLowPoints(smokeMap)

	sizeChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(lowPoints))

	go func() {
		wg.Wait()
		close(sizeChannel)
	}()

	for index := range lowPoints {
		go func(point *Point) {
			defer wg.Done()

			var dfs func(x int, y int)
			size := 0
			visited := make([][]bool, len(smokeMap))
			for index := range smokeMap {
				visited[index] = make([]bool, len(smokeMap[index]))
			}
			dfs = func(x int, y int) {
				if x < 0 ||
					y < 0 ||
					x >= len(smokeMap) ||
					y >= len(smokeMap[x]) ||
					visited[x][y] ||
					smokeMap[x][y] == 9 {
					return
				}

				visited[x][y] = true

				size++

				dfs(x-1, y)
				dfs(x+1, y)
				dfs(x, y-1)
				dfs(x, y+1)
			}
			dfs(point.X, point.Y)
			sizeChannel <- size
		}(lowPoints[index])
	}

	var sizes []int
	for size := range sizeChannel {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)

	result := 1
	for i := len(sizes) - 3; i < len(sizes); i++ {
		result *= sizes[i]
	}

	return result
}

type Point struct {
	X int
	Y int
}

func findLowPoints(smokeMap [][]int) []*Point {
	var lowPoints []*Point
	for i := 0; i < len(smokeMap); i++ {
		for j := 0; j < len(smokeMap[i]); j++ {
			if (i-1 >= 0 && smokeMap[i-1][j] <= smokeMap[i][j]) ||
				(i+1 < len(smokeMap) && smokeMap[i+1][j] <= smokeMap[i][j]) ||
				(j-1 >= 0 && smokeMap[i][j-1] <= smokeMap[i][j]) ||
				(j+1 < len(smokeMap[i]) && smokeMap[i][j+1] <= smokeMap[i][j]) {
				continue
			}
			lowPoints = append(lowPoints, &Point{X: i, Y: j})
		}
	}
	return lowPoints
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
