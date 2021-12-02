package main

import "fmt"

func main() {
	fmt.Printf("Part 1 Solution: %d\n", part1())
	fmt.Printf("Part 2 Solution: %d\n", part2())
}

func part1() int {
	c := commands()

	horizontal := 0
	depth := 0

	for _, command := range c {
		switch command.direction {
		case "forward":
			horizontal += command.count
		case "up":
			depth -= command.count
		case "down":
			depth += command.count
		}
	}

	return horizontal * depth
}

func part2() int {
	c := commands()

	horizontal := 0
	depth := 0
	aim := 0

	for _, command := range c {
		switch command.direction {
		case "forward":
			horizontal += command.count
			depth += aim * command.count
		case "up":
			aim -= command.count
		case "down":
			aim += command.count
		}
	}

	return horizontal * depth
}
