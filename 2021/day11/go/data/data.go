package data

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetData(filename string) [][]int {
	var levels [][]int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var lineLevels []int
		lineLevelsString := strings.Split(scanner.Text(), "")
		for _, levelString := range lineLevelsString {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				log.Fatal(err)
			}
			lineLevels = append(lineLevels, level)
		}
		levels = append(levels, lineLevels)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return levels
}
