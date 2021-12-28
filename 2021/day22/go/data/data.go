package data

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Cube struct {
	On bool
	X1 int
	X2 int
	Y1 int
	Y2 int
	Z1 int
	Z2 int
}

func (c *Cube) Volume() int {
	return (c.X2 - c.X1 + 1) * (c.Y2 - c.Y1 + 1) * (c.Z2 - c.Z1 + 1)
}

func LineOverlap(min0, max0, min1, max1 int) (int, int) {
	return Max(min0, min1), Min(max0, max1)
}

func Overlap(curr *Cube, cubes []*Cube) int {
	var sums []int
	for i, c := range cubes {
		overlapMinX, overlapMaxX := LineOverlap(curr.X1, curr.X2, c.X1, c.X2)
		overlapMinY, overlapMaxY := LineOverlap(curr.Y1, curr.Y2, c.Y1, c.Y2)
		overlapMinZ, overlapMaxZ := LineOverlap(curr.Z1, curr.Z2, c.Z1, c.Z2)

		if overlapMaxX-overlapMinX >= 0 &&
			overlapMaxY-overlapMinY >= 0 &&
			overlapMaxZ-overlapMinZ >= 0 {
			newCube := &Cube{
				X1: overlapMinX,
				X2: overlapMaxX,
				Y1: overlapMinY,
				Y2: overlapMaxY,
				Z1: overlapMinZ,
				Z2: overlapMaxZ,
			}
			sums = append(sums, newCube.Volume()-Overlap(newCube, cubes[i+1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	sum := 0
	for _, val := range sums {
		sum += val
	}
	return sum
}

func GetData(filename string) []*Cube {
	var cubes []*Cube

	r := regexp.MustCompile(`^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		matches := r.FindStringSubmatch(line)

		cubes = append(cubes, &Cube{
			On: onOffBool(matches[1]),
			X1: parseInt(matches[2]),
			X2: parseInt(matches[3]),
			Y1: parseInt(matches[4]),
			Y2: parseInt(matches[5]),
			Z1: parseInt(matches[6]),
			Z2: parseInt(matches[7]),
		})
	}

	return cubes
}

func onOffBool(s string) bool {
	if s == "on" {
		return true
	} else if s == "off" {
		return false
	}
	panic("invalid on/off value")
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
