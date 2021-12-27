package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mreed19/advent-of-code/2021/day21/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	defer duration(track("part1"))

	players := data.GetData(filename)

	die := data.DeterministicDice{
		Sides:      100,
		Count:      0,
		TotalCount: 0,
	}

	winner := -1

	for winner == -1 {
		for playerNum, player := range players {
			for i := 0; i < 3; i++ {
				player.Position = (player.Position+die.Roll()-1)%10 + 1
			}
			player.Score += player.Position

			if player.Score >= 1000 {
				winner = playerNum
				break
			}
		}
	}

	var results int
	for i, player := range players {
		if i != winner {
			results = player.Score * die.TotalCount
		}
	}

	return results
}

type Game struct {
	turn      int
	positions []int
	scores    []int64
}

func (g Game) ToString() string {
	return fmt.Sprintf("%d:%d:%d:%d:%d", g.positions[0], g.scores[0], g.positions[1], g.scores[1], g.turn)
}

var quantumDie = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func part2(filename string) int64 {
	defer duration(track("part2"))

	players := data.GetData(filename)

	wins := make(map[string]int64)

	game := Game{
		turn:      0,
		positions: []int{players[0].Position, players[1].Position},
		scores:    []int64{0, 0},
	}

	var play func(Game) int64
	play = func(game Game) int64 {
		if val, ok := wins[game.ToString()]; ok {
			return val
		}

		sum := int64(0)

		for roll, count := range quantumDie {
			newGame := Game{
				turn:      game.turn,
				positions: []int{game.positions[0], game.positions[1]},
				scores:    []int64{game.scores[0], game.scores[1]},
			}

			newGame.positions[newGame.turn] = (newGame.positions[newGame.turn]+roll-1)%10 + 1
			newGame.scores[newGame.turn] += int64(newGame.positions[newGame.turn])

			if newGame.scores[newGame.turn] >= 21 {
				if newGame.turn == 0 {
					sum += count
				}
			} else {
				newGame.turn = newGame.turn ^ 1
				sum += count * play(newGame)
			}
		}
		wins[game.ToString()] = sum
		return sum
	}

	return play(game)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
