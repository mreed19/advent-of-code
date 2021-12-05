package data

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

type Vent struct {
	Point1 *Point `json:"Point1"`
	Point2 *Point `json:"Point2"`
}

func NewPoint(pointText string) (*Point, error) {
	parts := strings.Split(pointText, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("NewPoint: x: %v", err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("NewPoint: y: %v", err)
	}

	return &Point{X: x, Y: y}, nil
}

func NewHZVent(line string) (*Vent, error) {
	parts := strings.Split(line, " -> ")
	point1, err := NewPoint(parts[0])
	if err != nil {
		return nil, fmt.Errorf("NewVent: point1: %v", err)
	}
	point2, err := NewPoint(parts[1])
	if err != nil {
		return nil, fmt.Errorf("NewVent: point2: %v", err)
	}

	if point1.X == point2.X || point1.Y == point2.Y {
		return &Vent{Point1: point1, Point2: point2}, nil
	}
	return nil, nil
}

func NewVent(line string) (*Vent, error) {
	parts := strings.Split(line, " -> ")
	point1, err := NewPoint(parts[0])
	if err != nil {
		return nil, fmt.Errorf("NewVent: point1: %v", err)
	}
	point2, err := NewPoint(parts[1])
	if err != nil {
		return nil, fmt.Errorf("NewVent: point2: %v", err)
	}

	return &Vent{Point1: point1, Point2: point2}, nil
}

func ReadInput(filename string, ventConstructor func(string) (*Vent, error)) []*Vent {
	var vents []*Vent

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		vent, err := ventConstructor(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if vent != nil {
			vents = append(vents, vent)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return vents
}
