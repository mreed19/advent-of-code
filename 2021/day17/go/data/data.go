package data

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

type Target struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func (t *Target) InTargetArea(x, y int) bool {
	return x >= t.MinX &&
		x <= t.MaxX &&
		y >= t.MinY &&
		y <= t.MaxY
}

func (t *Target) PastTargetArea(y int) bool {
	return y < t.MinY
}

func GetData(filename string) *Target {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)`)

	match := r.FindStringSubmatch(string(file))

	return &Target{
		MinX: parseInt(match[1]),
		MaxX: parseInt(match[2]),
		MinY: parseInt(match[3]),
		MaxY: parseInt(match[4]),
	}
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}
