package data

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var polarities = [][3]int{
	{1, 1, 1},
	{1, -1, 1},
	{1, 1, -1},
	{1, -1, -1},
	{-1, 1, 1},
	{-1, -1, 1},
	{-1, 1, -1},
	{-1, -1, -1},
}

var permutations = [][3]int{
	{0, 1, 2},
	{0, 2, 1},
	{1, 0, 2},
	{1, 2, 0},
	{2, 0, 1},
	{2, 1, 0},
}

func getRotations() [][3][2]int {
	var rotations [][3][2]int
	for _, polarity := range polarities {
		for _, permutation := range permutations {
			rotations = append(rotations, [3][2]int{
				{polarity[0], permutation[0]},
				{polarity[1], permutation[1]},
				{polarity[2], permutation[2]},
			})
		}
	}
	return rotations
}

var Rotations = getRotations()

func Rotate(point [3]int, rotation [3][2]int) [3]int {
	return [3]int{
		point[rotation[0][1]] * rotation[0][0],
		point[rotation[1][1]] * rotation[1][0],
		point[rotation[2][1]] * rotation[2][0],
	}
}

func Translate(coords [3]int, translation [3]int) [3]int {
	return [3]int{
		coords[0] + translation[0],
		coords[1] + translation[1],
		coords[2] + translation[2],
	}
}

type Point struct {
	Coords [3]int
	Edges  [][3]int
}

type Scanner struct {
	Num      int
	Position [3]int
	Points   []*Point
}

func (s1 *Scanner) AddScanner(s2 *Scanner) {
	for _, p1 := range s1.Points {
		for _, p2 := range s2.Points {
			p1.Edges = append(p1.Edges, NewEdge(p1.Coords, p2.Coords))
			p2.Edges = append(p2.Edges, NewEdge(p2.Coords, p1.Coords))
		}
	}
	s1.Points = append(s1.Points, s2.Points...)
}

func Equal(p1, p2 [3]int) bool {
	equal := true
	for i := 0; i < 3; i++ {
		if p1[i] != p2[i] {
			equal = false
			break
		}
	}
	return equal
}

func SetEqual(e1, e2 [3]int) bool {
	equal := true
	for i := 0; i < 3; i++ {
		match := false
		for j := 0; j < 3; j++ {
			if Abs(e1[i]) == Abs(e2[j]) {
				match = true
				break
			}
		}
		if !match {
			equal = false
			break
		}
	}
	return equal
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func NewEdge(e1, e2 [3]int) [3]int {
	return [3]int{
		e2[0] - e1[0],
		e2[1] - e1[1],
		e2[2] - e1[2],
	}
}

func (s *Scanner) MapEdges() {
	for i := 0; i < len(s.Points)-1; i++ {
		for j := i + 1; j < len(s.Points); j++ {
			s.Points[i].Edges = append(s.Points[i].Edges, NewEdge(s.Points[i].Coords, s.Points[j].Coords))
			s.Points[j].Edges = append(s.Points[j].Edges, NewEdge(s.Points[j].Coords, s.Points[i].Coords))
		}
	}
}

func GetData(filename string) []*Scanner {
	var scanners []*Scanner

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bytes), "\n")

	var scanner *Scanner
	for _, line := range lines {
		if line == "" {
			scanner.Num = len(scanners)
			scanners = append(scanners, scanner)
		} else if strings.Contains(line, "scanner") {
			scanner = &Scanner{}
		} else {
			numStrings := strings.Split(line, ",")
			var nums []int
			for _, numString := range numStrings {
				num, err := strconv.Atoi(numString)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, num)
			}
			scanner.Points = append(scanner.Points, &Point{
				Coords: [3]int{
					nums[0],
					nums[1],
					nums[2],
				},
			})
		}
	}
	for _, scanner := range scanners {
		scanner.MapEdges()
	}

	return scanners
}
