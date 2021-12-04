package bingo

type position struct {
	x int
	y int
}

type BingoBoard struct {
	Sum       int
	Done      bool
	rowChecks []int
	colChecks []int
	board     map[int]position
}

func NewBingoBoard(board [][]int) *BingoBoard {
	bingoBoard := &BingoBoard{
		Done:      false,
		colChecks: make([]int, len(board)),
		rowChecks: make([]int, len(board[0])),
		board:     make(map[int]position),
	}

	for i := 0; i < len(board); i++ {
		bingoBoard.rowChecks[i] = 0
		for j := 0; j < len(board[i]); j++ {
			bingoBoard.rowChecks[j] = 0
			val := board[i][j]
			bingoBoard.Sum += val
			bingoBoard.board[val] = position{x: j, y: i}
		}
	}

	return bingoBoard
}

func (b *BingoBoard) MarkValue(val int) {
	if pos, ok := b.board[val]; ok {
		b.Sum -= val
		b.rowChecks[pos.y]++
		b.colChecks[pos.x]++

		if b.rowChecks[pos.y] >= len(b.colChecks) || b.colChecks[pos.x] >= len(b.rowChecks) {
			b.Done = true
		}
	}
}
