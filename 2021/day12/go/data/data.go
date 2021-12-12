package data

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type Cave struct {
	Name        string
	IsSmallCave bool
}

type CaveGraph struct {
	Caves   map[string]*Cave
	CaveMap map[string][]*Cave
}

func GetData(filename string) *CaveGraph {
	caveGraph := &CaveGraph{
		Caves:   make(map[string]*Cave),
		CaveMap: make(map[string][]*Cave),
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		caves := strings.Split(line, "-")
		for _, caveName := range caves {
			if _, ok := caveGraph.Caves[caveName]; !ok {
				caveGraph.Caves[caveName] = &Cave{
					Name:        caveName,
					IsSmallCave: IsLower(caveName),
				}
			}
		}

		caveGraph.CaveMap[caves[0]] = append(caveGraph.CaveMap[caves[0]], caveGraph.Caves[caves[1]])
		caveGraph.CaveMap[caves[1]] = append(caveGraph.CaveMap[caves[1]], caveGraph.Caves[caves[0]])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return caveGraph
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
