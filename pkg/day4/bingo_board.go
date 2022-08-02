package day4

type bingoBoard struct {
	board [][]int
	marks [][]int
}

func newBingoBoard(board [][]int) *bingoBoard {
	// Init with empty 5x5 matrix
	return &bingoBoard{board: board, marks: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}}
}

func (b *bingoBoard) MarkNumber(num int) bool {
	// Check board for the number and mark it on the "marks" matrix
	// Also change the found number to 0 on the board. To make the sum easier
	for i, row := range b.board {
		for j, value := range row {
			if value == num {
				b.marks[i][j] = value
				b.board[i][j] = 0
			}
		}
	}

	return b.winCondition()
}

func (b *bingoBoard) internalWinCondition(matrix [][]int) bool {
	for _, row := range matrix {
		rowCheck := true
		for _, val := range row {
			rowCheck = rowCheck && val != 0
		}

		if rowCheck {
			return true
		}
	}

	return false
}

func (b *bingoBoard) winCondition() bool {
	horizontal := b.internalWinCondition(b.marks)
	vertical := b.internalWinCondition(transpose(b.marks))
	return horizontal || vertical
}

func (b *bingoBoard) SumMarked() int {
	return sum(b.marks)
}

func (b *bingoBoard) SumUnmarked() int {
	return sum(b.board)
}

func sum(matrix [][]int) int {
	sum := 0
	for _, ints := range matrix {
		for _, val := range ints {
			sum += val
		}
	}

	return sum
}

func transpose(a [][]int) [][]int {
	newArr := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	return newArr
}
