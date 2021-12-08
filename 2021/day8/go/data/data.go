package data

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetData(filename string) ([][]string, [][]string) {
	var nums [][]string
	var outputs [][]string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		nums = append(nums, strings.Split(parts[0], " "))
		outputs = append(outputs, strings.Split(parts[1], " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums, outputs
}
