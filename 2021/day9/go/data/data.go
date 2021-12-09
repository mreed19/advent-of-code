package data

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetData(filename string) [][]int {

	var smokeMap [][]int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var lineNums []int
		for _, lineNumString := range strings.Split(line, "") {
			lineNum, err := strconv.Atoi(lineNumString)
			if err != nil {
				log.Fatal(err)
			}
			lineNums = append(lineNums, lineNum)
		}
		smokeMap = append(smokeMap, lineNums)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return smokeMap
}
