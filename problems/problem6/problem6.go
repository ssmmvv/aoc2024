package problem6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	NORTH   = '^'
	WEST    = '<'
	EAST    = '>'
	SOUTH   = 'v'
	BARRIER = '#'
	VISITED = 'X'
)

func Problem6(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	guard_chars := "^><v"

	// Characters in the key must be before/after the values,
	// e.g. mustBeBefore[1] = {2, 3} means 1 must be before either 1 or 2
	orig_board := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		newRow := make([]rune, len(line))
		for idx, new_rune := range line {
			newRow[idx] = new_rune
		}
		orig_board = append(orig_board, newRow)
	}

	// for every potential obstruction coordinate
	num_obstacle_possible := 0
	for ob_row := 0; ob_row < len(orig_board); ob_row++ {
		for ob_col := 0; ob_col < len(orig_board[0]); ob_col++ {
			if strings.Contains(guard_chars, string(orig_board[ob_row][ob_col])) {
				continue // can't place the obstacle on the goard
			}

			// create a new board
			test_board := cloneBoard(orig_board)

			// add an obstable
			test_board[ob_row][ob_col] = BARRIER

			// test if there's a cycle
			if hasCycle(test_board) {
				num_obstacle_possible++
			}
		}
	}

	fmt.Printf("%d obstacle locations\n", num_obstacle_possible)

	// row, col := findGuard(orig_board)
	// for inBounds(orig_board, row, col) {
	// 	row, col = advance(orig_board, row, col)
	// }

	// printBoard(orig_board)
	// fmt.Printf("visited %d locations\n", countVisited(orig_board))
}

/* Detects a cycle in the guard's traversal. We do this by keeping a set of distinct locations *AND* facings
 * because we're allowed to revisit the same location as long as we're facing a different direction. If the
 * guard is in the same location *AND* facing as it was prior, then it's a cycle.
 */
func hasCycle(board [][]rune) bool {
	visited := map[string]bool{}

	row, col := findGuard(board)
	guard := board[row][col]
	for inBounds(board, row, col) {
		guard = board[row][col]
		key := buildKey(guard, row, col)
		_, found := visited[key]
		if found {
			return true
		}
		visited[key] = true
		row, col = advance(board, row, col)
	}
	return false
}

func buildKey(guard rune, row, col int) string {
	return fmt.Sprintf("%c:%d,%d", guard, row, col)
}

func findGuard(board [][]rune) (int, int) {
	for rowIdx, rowChars := range board {
		for col, char := range rowChars {
			if char == NORTH || char == WEST || char == SOUTH || char == EAST {
				return rowIdx, col
			}
		}
	}
	return -1, -1 // should never be reached on valid unput
}

func inBounds(board [][]rune, row, col int) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}

// advances the guard and returns the new coordinates of the guard. Write 'X' to the board
// as the guard moves
func advance(board [][]rune, row, col int) (int, int) {
	dx := 0
	dy := 0
	existing_direction := board[row][col]
	switch existing_direction {
	case NORTH:
		dy = -1
	case SOUTH:
		dy = 1
	case WEST:
		dx = -1
	case EAST:
		dx = 1
	}

	new_col := col + dx
	new_row := row + dy

	// did we go off the end?
	if !inBounds(board, new_row, new_col) {
		board[row][col] = VISITED // mark our last location as visited
		return new_row, new_col   // TODO remove this du\plicate check in the main function
	}

	// Are we obstructed
	if board[new_row][new_col] == BARRIER {
		// if we are, turn right
		board[row][col] = turnRight(existing_direction)
		return row, col // return the same coordinates (but facing changed)
	} else {
		// Advance
		// first, the existing location becomes an 'X'
		board[row][col] = VISITED
		// write the guard to the new location
		board[new_row][new_col] = existing_direction
		// return new coordinates
		return new_row, new_col
	}
}

func turnRight(current_dir rune) rune {
	switch current_dir {
	case NORTH:
		return EAST
	case EAST:
		return SOUTH
	case SOUTH:
		return WEST
	case WEST:
		return NORTH
	}
	panic(1)
}

func countVisited(board [][]rune) int {
	count := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == VISITED {
				count++
			}
		}
	}
	return count
}

func printBoard(board [][]rune) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}

func cloneBoard(board [][]rune) [][]rune {
	new_board := make([][]rune, len(board))
	for row, rowChars := range board {
		new_board[row] = make([]rune, len(rowChars))
		for idx, character := range rowChars {
			new_board[row][idx] = character
		}
	}
	return new_board
}
