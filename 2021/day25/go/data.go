package main

import (
	"bufio"
	"log"
	"os"
)

type Point struct {
	X int
	Y int
}

type Movers map[Point]struct{}

type Grid struct {
	Rows        int
	Cols        int
	EastMovers  Movers
	SouthMovers Movers
}

func (g *Grid) ToString() string {
	str := ""
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			p := Point{X: j, Y: i}
			if _, ok := g.EastMovers[p]; ok {
				str += ">"
			} else if _, ok := g.SouthMovers[p]; ok {
				str += "v"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func (g *Grid) IsSpotClear(p Point) bool {
	_, eastMoverOk := g.EastMovers[p]
	_, southMoverOk := g.SouthMovers[p]
	return !eastMoverOk && !southMoverOk
}

func (m *Movers) Move(p1, p2 Point) {
	(*m)[p2] = struct{}{}
	delete(*m, p1)
}

func (g *Grid) MoveEast() int {
	var newPoints [][]Point
	for p := range g.EastMovers {
		newPoint := p
		newPoint.X = (p.X + 1) % g.Cols
		if g.IsSpotClear(newPoint) {
			newPoints = append(newPoints, []Point{p, newPoint})
		}
	}
	for _, move := range newPoints {
		g.EastMovers.Move(move[0], move[1])
	}
	return len(newPoints)
}

func (g *Grid) MoveSouth() int {
	var newPoints [][]Point
	for p := range g.SouthMovers {
		newPoint := p
		newPoint.Y = (p.Y + 1) % g.Rows
		if g.IsSpotClear(newPoint) {
			newPoints = append(newPoints, []Point{p, newPoint})
		}
	}

	for _, move := range newPoints {
		g.SouthMovers.Move(move[0], move[1])
	}
	return len(newPoints)
}

func GetData(filename string) *Grid {

	grid := &Grid{
		EastMovers:  make(Movers),
		SouthMovers: make(Movers),
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rows := 0
	cols := 0
	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		if line != "" {
			rows++
			for col, char := range line {
				p := Point{X: col, Y: rows - 1}
				if char == '>' {
					grid.EastMovers[p] = struct{}{}
				} else if char == 'v' {
					grid.SouthMovers[p] = struct{}{}
				}
			}
		}
	}
	grid.Rows = rows
	grid.Cols = cols

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}
