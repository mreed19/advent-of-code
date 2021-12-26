package data

import (
	"log"
	"os"
	"strings"
)

type Bit uint8
type Algorithm map[uint16]Bit
type Image [][]Bit

func (i Image) ToString() string {
	str := ""
	for _, row := range i {
		for _, col := range row {
			if col == 1 {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}

	return str
}

func (im Image) GetBit(i, j int, baseBit Bit) Bit {
	rows := len(im)
	cols := len(im[0])

	if i >= 0 && i < rows && j >= 0 && j < cols {
		return im[i][j]
	}

	return baseBit
}

func GetData(filename string) (Algorithm, Image) {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	algorithmLine := lines[0]
	algorithm := make(map[uint16]Bit)
	for i, val := range algorithmLine {
		if val == '#' {
			algorithm[uint16(i)] = 1
		} else {
			algorithm[uint16(i)] = 0
		}
	}

	lines = lines[1:]

	var image [][]Bit
	for _, line := range lines {
		if line == "" {
			continue
		}
		var imageLine []Bit
		for _, val := range line {
			if val == '#' {
				imageLine = append(imageLine, Bit(1))
			} else {
				imageLine = append(imageLine, Bit(0))
			}
		}
		image = append(image, imageLine)
	}

	return algorithm, image
}
