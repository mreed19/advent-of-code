package data

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

type Fold struct {
	Direction string `json:"Direction"`
	FoldLine  int    `json:"FoldLine"`
}

func GetData(filename string) (map[int]map[int]bool, []*Fold) {

	points := make(map[int]map[int]bool)
	var folds []*Fold

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	startFolds := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			startFolds = true
			continue
		}

		if startFolds {
			line = strings.ReplaceAll(line, "fold along ", "")
			parts := strings.Split(line, "=")
			foldLine, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			folds = append(folds, &Fold{Direction: parts[0], FoldLine: foldLine})
		} else {
			parts := strings.Split(line, ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			if points[x] == nil {
				points[x] = make(map[int]bool)
			}
			points[x][y] = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return points, folds
}
