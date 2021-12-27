package data

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Position int
	Score    int
}

type Dice interface {
	Roll() int
}

type DeterministicDice struct {
	Sides      int
	Count      int
	TotalCount int
}

func (d *DeterministicDice) Roll() int {
	d.TotalCount++
	d.Count = (d.Count % d.Sides) + 1
	return d.Count
}

func GetData(filename string) []*Player {
	var players []*Player

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line != "" {
			position, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				log.Fatal(err)
			}
			players = append(players, &Player{
				Position: position,
				Score:    0,
			})
		}
	}

	return players
}
