package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day20/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	algo, image := data.GetData(filename)

	return getPixelCountAfterIterations(image, algo, 2)
}

func part2(filename string) int {
	defer duration(track("part2"))

	algo, image := data.GetData(filename)

	return getPixelCountAfterIterations(image, algo, 50)
}

func getPixelCountAfterIterations(image data.Image, algo data.Algorithm, iterations int) int {
	var lit int
	for count := 0; count < iterations; count++ {
		lit = 0

		rows := len(image)
		cols := len(image[0])

		var baseBit data.Bit
		if algo[0] == 1 && count%2 == 1 {
			baseBit = data.Bit(1)
		} else {
			baseBit = data.Bit(0)
		}

		for row := range image {
			image[row] = append(append([]data.Bit{baseBit}, image[row]...), baseBit)
		}

		topFiller := make([]data.Bit, cols+2)
		bottomFiller := make([]data.Bit, cols+2)
		for i := 0; i < cols+2; i++ {
			topFiller[i] = baseBit
			bottomFiller[i] = baseBit
		}
		image = append([][]data.Bit{topFiller}, image...)
		image = append(image, bottomFiller)

		newImage := make(data.Image, rows+2)
		for i := 0; i < len(newImage); i++ {
			newImage[i] = make([]data.Bit, cols+2)
			for j := 0; j < len(newImage[i]); j++ {
				newBit := uint16(0)
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						newBit = newBit << 1
						newBit += uint16(image.GetBit(i+k, j+l, baseBit))

					}
				}
				val := algo[newBit]
				lit += int(val)
				newImage[i][j] = val
			}
		}

		image = newImage
	}
	return lit
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
