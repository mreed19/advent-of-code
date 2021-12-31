package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	// fmt.Println("Part 1 Solution:", part1("input-sample.txt"))
	fmt.Println("Part 1 Solution:", part1("input.txt"))
	// fmt.Println("Part 2 Solution:", part2("input-sample-2.txt"))
	fmt.Println("Part 2 Solution:", part2("input-2.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	initialState := GetData(filename)

	return GetMinEnergy(initialState)
}

func part2(filename string) int {
	defer duration(track("part2"))

	initialState := GetData(filename)

	return GetMinEnergy(initialState)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func GetMinEnergy(initialState State) int {
	visited := make(map[string]int)

	minEnergy := math.MaxInt

	var play func(State)
	play = func(s State) {
		if s.Energy > minEnergy {
			return
		}

		energy := s.Energy
		hallway := s.Hallway
		rooms := s.Rooms

		if s.Complete() {
			if energy < minEnergy {
				minEnergy = energy
			}
			return
		}

		stateKey := s.ToString()

		if val, ok := visited[stateKey]; ok && val <= energy {
			return
		}

		visited[stateKey] = energy

		// Move Hallway to Room
		for hallwayIndex, hallwaySpot := range hallway {
			if hallwaySpot == '.' || hallwaySpot == 'X' {
				continue
			}

			// Search for matching Room that is settled
			for roomIndex, room := range rooms {
				if hallwaySpot == room.Letter {
					if len(room.Unsettled) == 0 {

						canGetToRoom := true
						if hallwayIndex < Rooms[hallwaySpot] {
							for i := hallwayIndex + 1; i < Rooms[hallwaySpot]; i++ {
								if hallway[i] != '.' {
									canGetToRoom = false
									break
								}
							}
						} else {
							for i := hallwayIndex - 1; i > Rooms[hallwaySpot]; i-- {
								if hallway[i] != '.' {
									canGetToRoom = false
									break
								}
							}
						}

						if canGetToRoom {
							newEnergy := energy
							newHallway := hallway
							newRooms := rooms

							newHallway[hallwayIndex] = '.'
							newRooms[roomIndex].Settled = append(newRooms[roomIndex].Settled, hallwaySpot)
							newEnergy += (abs(Rooms[room.Letter]-hallwayIndex) + room.Spots - len(room.Settled)) * MoveCosts[room.Letter]

							play(State{
								Energy:  newEnergy,
								Rooms:   newRooms,
								Hallway: newHallway,
							})
						}
					}
				}
			}
		}
		// Done Moving Hallway to Room

		// Move Room to Hallway
		for roomIndex, room := range rooms {
			if len(room.Unsettled) > 0 {
				unsettled := room.Unsettled[len(room.Unsettled)-1]

				for _, hallwayMove := range s.GetHallwayMoves(room.Letter) {
					newEnergy := energy
					newHallway := hallway
					newRooms := rooms

					newHallway[hallwayMove.Index] = unsettled

					newRooms[roomIndex].Unsettled = newRooms[roomIndex].Unsettled[:len(newRooms[roomIndex].Unsettled)-1]

					newEnergy += (room.Spots - len(room.Settled) - len(room.Unsettled) + hallwayMove.Dist) * MoveCosts[unsettled]

					play(State{
						Energy:  newEnergy,
						Hallway: newHallway,
						Rooms:   newRooms,
					})
				}
			}
		}
		// Done moving Room to Hallway
	}
	// End Play function

	play(initialState)

	return minEnergy
}
