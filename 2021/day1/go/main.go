package main

import (
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	m := measurements()
	inc := 0
	for i := 0; i+1 < len(m); i++ {
		if m[i] < m[i+1] {
			inc++
		}
	}
	fmt.Println(inc)
}

func part2() {
	m := measurements()
	inc := 0
	for i := 0; i+3 < len(m); i++ {
		prev := m[i] + m[i+1] + m[i+2]
		curr := m[i+1] + m[i+2] + m[i+3]
		if curr > prev {
			inc++
		}
	}
	fmt.Println(inc)
}
