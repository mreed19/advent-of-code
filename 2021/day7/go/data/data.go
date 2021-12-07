package data

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetPositions(filename string) []int {
	var positions []int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		positionStrings := strings.Split(scanner.Text(), ",")

		for _, positionString := range positionStrings {
			position, err := strconv.Atoi(positionString)
			if err != nil {
				log.Fatal(err)
			}

			positions = append(positions, position)
		}
	}
	return positions
}
