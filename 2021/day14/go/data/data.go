package data

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Polymer struct {
	Next      []string `json:"Next"`
	Count     int      `json:"Count"`
	NextCount int      `json:"NextCount"`
}

func GetData(filename string) map[string]*Polymer {
	polymers := make(map[string]*Polymer)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	template := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		// do something
		line := scanner.Text()
		parts := strings.Split(line, " -> ")

		polymers[parts[0]] = &Polymer{
			Next: []string{parts[0][:1] + parts[1][:1], parts[1][:1] + parts[0][1:2]},
		}
	}

	for i := 0; i < len(template)-1; i++ {
		polymers[template[i:i+2]].Count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return polymers
}
