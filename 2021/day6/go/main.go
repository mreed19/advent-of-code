package main

import (
	"fmt"
	"log"

	"github.com/mreed19/advent-of-code/2021/day6/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1())
	fmt.Println("Part 2 Solution:", part2())
}

func part1() int {
	return getLanterFishCount("data/input.txt", 80)
}

func part2() int {
	return getLanterFishCount("data/input.txt", 256)
}

func getLanterFishCount(filename string, days int) int {
	lanternFish, err := data.GetLanternFish(filename)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < days; i++ {
		newFish := lanternFish[0]
		lanternFish = lanternFish[1:]
		lanternFish[data.LanternFishReset] += newFish
		lanternFish = append(lanternFish, newFish)
	}

	sum := 0
	for _, val := range lanternFish {
		sum += val
	}

	return sum
}
