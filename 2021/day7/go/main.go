package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/mreed19/advent-of-code/2021/day7/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 1 Medians Solution:", part1medians("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	d := data.GetPositions(filename)
	dataLength := len(d)

	sort.Ints(d)

	leftSums := make([]int, dataLength)
	rightSums := make([]int, dataLength)

	for i := 1; i < dataLength; i++ {
		leftSums[i] = leftSums[i-1] + (i * (d[i] - d[i-1]))
		rightSums[dataLength-1-i] = rightSums[dataLength-i] + (i * (d[dataLength-i] - d[dataLength-i-1]))
	}

	min := leftSums[0] + rightSums[0]
	for i := 1; i < dataLength; i++ {
		sum := leftSums[i] + rightSums[i]
		if sum < min {
			min = sum
		}
	}
	return min
}

func part1medians(filename string) int {
	d := data.GetPositions(filename)

	sort.Ints(d)

	median := getMedian(d)

	sum := 0
	for i := 0; i < len(d); i++ {
		sum += int(math.Abs(float64(median) - float64(d[i])))
	}
	return sum
}

func part2(filename string) int {
	d := data.GetPositions(filename)

	meanFloor, meanCeil := getMean(d)

	floorSum := 0
	ceilSum := 0
	for i := 0; i < len(d); i++ {
		floorDiff := math.Abs(meanFloor - float64(d[i]))
		ceilDiff := math.Abs(meanCeil - float64(d[i]))
		floorSum += int(floorDiff * (floorDiff + 1) / 2)
		ceilSum += int(ceilDiff * (ceilDiff + 1) / 2)
	}

	return int(math.Min(float64(floorSum), float64(ceilSum)))
}

func getMedian(arr []int) int {
	mid := len(arr) / 2
	if len(arr)%2 == 0 {
		return (arr[mid-1] + arr[mid]) / 2
	}
	return arr[mid]
}

func getMean(arr []int) (float64, float64) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	meanFloat := float64(sum) / float64(len(arr))
	return math.Floor(meanFloat), math.Ceil(meanFloat)
}
