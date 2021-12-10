package data

import (
	"bufio"
	"log"
	"os"
)

func GetData(filename string) [][]rune {
	var lines [][]rune

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
