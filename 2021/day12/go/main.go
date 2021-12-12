package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day12/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	d := data.GetData(filename)

	pathCount := 0
	visited := make(map[string]bool)
	var dfs func(c *data.Cave)
	dfs = func(c *data.Cave) {
		if c.Name == "end" {
			pathCount++
			return
		}
		visited[c.Name] = c.IsSmallCave

		for _, cave := range d.CaveMap[c.Name] {
			if visited[cave.Name] {
				continue
			}
			dfs(cave)
		}
		visited[c.Name] = false
	}

	dfs(d.Caves["start"])

	return pathCount
}

func part2(filename string) int {
	defer duration(track("part2"))

	d := data.GetData(filename)

	pathCount := 0
	visited := make(map[string]int)
	var dfs func(c *data.Cave, allowDouble bool)
	dfs = func(c *data.Cave, allowDouble bool) {
		currAllowDouble := allowDouble

		if val, ok := visited[c.Name]; ok && c.IsSmallCave && val >= 1 {
			if !currAllowDouble {
				return
			}
			currAllowDouble = false
		}

		if c.Name == "start" {
			if _, ok := visited["start"]; ok {
				return
			}
		}

		if c.Name == "end" {
			pathCount++
			return
		}

		visits := 1
		if val, ok := visited[c.Name]; ok {
			visits += val
		}

		visited[c.Name] = visits

		for _, cave := range d.CaveMap[c.Name] {
			dfs(cave, currAllowDouble)
		}

		visited[c.Name]--
	}

	dfs(d.Caves["start"], true)

	return pathCount
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
