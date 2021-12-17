package data

import (
	"bufio"
	"log"
	"os"
)

var hexMap = map[rune][]rune{
	'0': []rune("0000"),
	'1': []rune("0001"),
	'2': []rune("0010"),
	'3': []rune("0011"),
	'4': []rune("0100"),
	'5': []rune("0101"),
	'6': []rune("0110"),
	'7': []rune("0111"),
	'8': []rune("1000"),
	'9': []rune("1001"),
	'A': []rune("1010"),
	'B': []rune("1011"),
	'C': []rune("1100"),
	'D': []rune("1101"),
	'E': []rune("1110"),
	'F': []rune("1111"),
}

func GetData(filename string) []rune {
	var output []rune

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		for _, char := range line {
			output = append(output, hexMap[char]...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
