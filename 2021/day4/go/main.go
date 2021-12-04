package main

import (
	"fmt"

	"github.com/mreed19/advent-of-code/2021/day4/go/bingo"
	"github.com/mreed19/advent-of-code/2021/day4/go/data"
)

func main() {
	fmt.Printf("Part 1 Solution: %d\n", part1())
	fmt.Printf("Part 2 Solution: %d\n", part2())
}

func part1() int {
	vals, boardDatasets := data.GetInput()

	boards := make([]*bingo.BingoBoard, len(boardDatasets))
	for i, boardDataset := range boardDatasets {
		boards[i] = bingo.NewBingoBoard(boardDataset)
	}

	max := 0
	finished := false

	for _, val := range vals {
		for _, board := range boards {
			if !board.Done {
				board.MarkValue(val)
				if board.Done {
					finished = true
					currVal := board.Sum * val
					if currVal > max {
						max = currVal
					}
				}
			}
		}
		if finished {
			return max
		}
	}

	return -1
}

func part2() int {
	vals, boardDatasets := data.GetInput()

	boards := make([]*bingo.BingoBoard, len(boardDatasets))
	for i, boardDataset := range boardDatasets {
		boards[i] = bingo.NewBingoBoard(boardDataset)
	}

	boardCount := len(boards)
	for _, val := range vals {
		for _, board := range boards {
			if !board.Done {
				board.MarkValue(val)
				if board.Done {
					boardCount--
					if boardCount <= 0 {
						return board.Sum * val
					}
				}
			}
		}
	}

	return -1
}
