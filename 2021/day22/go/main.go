package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day22/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	cubes := data.GetData(filename)
	for i, j := 0, len(cubes)-1; i < j; i, j = i+1, j-1 {
		cubes[i], cubes[j] = cubes[j], cubes[i]
	}

	sum := 0
	newCubes := make([]*data.Cube, 0)
	for _, cube := range cubes {
		cube.X1 = data.Max(-50, cube.X1)
		cube.X2 = data.Min(50, cube.X2)
		cube.Y1 = data.Max(-50, cube.Y1)
		cube.Y2 = data.Min(50, cube.Y2)
		cube.Z1 = data.Max(-50, cube.Z1)
		cube.Z2 = data.Min(50, cube.Z2)

		if cube.X1 <
			cube.X2 &&
			cube.Y1 <
				cube.Y2 &&
			cube.Z1 <
				cube.Z2 {

			if cube.On {
				sum += cube.Volume() - data.Overlap(cube, newCubes)
			}
			newCubes = append(newCubes, cube)
		}
	}

	return sum
}

func part2(filename string) int {
	defer duration(track("part2"))

	cubes := data.GetData(filename)
	for i, j := 0, len(cubes)-1; i < j; i, j = i+1, j-1 {
		cubes[i], cubes[j] = cubes[j], cubes[i]
	}

	sum := 0
	newCubes := make([]*data.Cube, 0)
	for _, cube := range cubes {
		if cube.On {
			sum += cube.Volume() - data.Overlap(cube, newCubes)
		}
		newCubes = append(newCubes, cube)
	}

	return sum
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
