package day4

import "fmt"

type bingoBoard struct {
	board      [][]int
	marks      []int
	BoardIndex int
}

func newBingoBoard(board [][]int, boardIndex int) *bingoBoard {
	// Init with empty 5x5 matrix
	return &bingoBoard{board: board, marks: []int{}, BoardIndex: boardIndex}
}

func (b *bingoBoard) MarkNumber(num int) bool {
	// Mark the board by setting the position to -1
	// Add the marked values to an array for later if needed
	for i, row := range b.board {
		for j, value := range row {
			if value == num {
				b.marks = append(b.marks, value)
				b.board[i][j] = -1
			}
		}
	}

	return b.winCondition()
}

func (b *bingoBoard) internalWinCondition(matrix [][]int) bool {
	for _, row := range matrix {
		rowCheck := true
		for _, val := range row {
			rowCheck = rowCheck && val < 0
		}

		if rowCheck {
			fmt.Println("Board", b.BoardIndex, "is a winner")
			return true
		}
	}

	return false
}

func (b *bingoBoard) winCondition() bool {
	return b.internalWinCondition(b.board) || b.internalWinCondition(transpose(b.board))
}

func (b *bingoBoard) SumUnmarked() int {
	sum := 0
	for _, ints := range b.board {
		for _, val := range ints {
			// Do not sum negatives used to mark the boards
			if val > 0 {
				sum += val
			}
		}
	}

	return sum
}

func (b *bingoBoard) LastMarkedNumber() int {
	return b.marks[len(b.marks)-1]
}

func (b *bingoBoard) PrintBoard() {
	fmt.Println("Board")
	printBoard(b.board)

	fmt.Println("Marks")
	fmt.Println(b.marks)
}

func printBoard(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[row]); column++ {
			fmt.Print(matrix[row][column], " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
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
