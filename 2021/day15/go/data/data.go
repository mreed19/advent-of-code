package data

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetData(filename string) [][]int {
	var grid [][]int
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		vals := strings.Split(line, "")
		var lineVals []int
		for _, val := range vals {
			valNum, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			lineVals = append(lineVals, valNum)
		}
		grid = append(grid, lineVals)
	}

	return grid
}
