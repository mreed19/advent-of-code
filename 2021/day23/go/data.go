package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var MoveCosts = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

var Rooms = map[rune]int{
	'A': 2,
	'B': 4,
	'C': 6,
	'D': 8,
}

type Room struct {
	Letter    rune
	Spots     int
	Settled   []rune
	Unsettled []rune
}

func (r *Room) ToString() string {
	return fmt.Sprintf("%s:%d:%s:%s", string(r.Letter), r.Spots, string(r.Settled), string(r.Unsettled))
}

type State struct {
	Energy  int
	Hallway [11]rune
	Rooms   [4]Room
}

type HallwayMove struct {
	Index int
	Dist  int
}

func (s *State) GetHallwayMoves(roomLetter rune) []HallwayMove {
	var moves []HallwayMove
	startIndex := Rooms[roomLetter]

	// left
	dist := 1
	for i := startIndex - 1; i >= 0; i-- {
		dist++
		if s.Hallway[i] != '.' {
			break
		}
		if i != 2 && i != 4 && i != 6 {
			moves = append(moves, HallwayMove{
				Index: i,
				Dist:  dist,
			})
		}
	}

	//right
	dist = 1
	for i := startIndex + 1; i < 11; i++ {
		dist++
		if s.Hallway[i] != '.' {
			break
		}
		if i != 4 && i != 6 && i != 8 {
			moves = append(moves, HallwayMove{
				Index: i,
				Dist:  dist,
			})
		}
	}

	return moves
}

func (s *State) Complete() bool {
	for _, room := range s.Rooms {
		if len(room.Settled) < room.Spots {
			return false
		}
	}
	return true
}

func (s *State) ToString() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", string(s.Hallway[0:11]), s.Rooms[0].ToString(), s.Rooms[1].ToString(), s.Rooms[2].ToString(), s.Rooms[3].ToString())
}

func GetData(filename string) State {
	state := State{}

	for i := 0; i < 11; i++ {
		state.Hallway[i] = '.'
	}
	var rooms [4]Room
	rooms[0].Letter = 'A'
	rooms[1].Letter = 'B'
	rooms[2].Letter = 'C'
	rooms[3].Letter = 'D'

	roomRegex := regexp.MustCompile("#(A|B|C|D)#(A|B|C|D)#(A|B|C|D)#(A|B|C|D)#")

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		matches := roomRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			for i := 0; i < 4; i++ {
				rooms[i].Unsettled = append([]rune(matches[i+1]), rooms[i].Unsettled...)
				rooms[i].Spots++
			}
		}
	}

	for i := range rooms {
		var settled []rune
		for _, unsettled := range rooms[i].Unsettled {
			if unsettled != rooms[i].Letter {
				break
			}
			settled = append(settled, unsettled)
		}
		rooms[i].Settled = settled
		rooms[i].Unsettled = rooms[i].Unsettled[len(settled):]
	}
	state.Rooms = rooms

	return state
}
