package problems

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	target_str = "XMAS"
)

func Problem4(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the input to a matrix of runes, each row is one line of file input text
	var char_matrix [][]rune
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		char_matrix = append(char_matrix, make([]rune, len(line)))
		for idx, char := range line {
			char_matrix[row][idx] = char
		}
		row++
	}

	// Iterate over the entire matrix
	num_xmas := 0
	for row := range len(char_matrix) {
		for col := range len(char_matrix[row]) {
			//num_xmas += get_num_xmas(char_matrix, row, col)
			if is_x_match(char_matrix, row, col) {
				num_xmas++
			}
		}
	}
	fmt.Println("Num XMAS: %d\n", num_xmas)
}

func get_num_xmas(char_matrix [][]rune, row, col int) int {
	num_xmas := 0
	// try all possible combinations of directions (except when both deltas zero)
	for drow := -1; drow <= 1; drow++ {
		for dcol := -1; dcol <= 1; dcol++ {
			if (dcol != 0 || drow != 0) && is_match(char_matrix, row, col, drow, dcol) {
				num_xmas++
			}
		}
	}
	return num_xmas
}

func is_match(char_matrix [][]rune, row, col, drow, dcol int) bool {
	target_idx := 0
	for target_idx < len(target_str) {
		// If we fell off the matrix, or don't match target str, return false
		if !in_bounds(char_matrix, row, col) {
			return false
		}
		if char_matrix[row][col] != rune(target_str[target_idx]) {
			return false
		}

		// advance both our cursor in the target string and the matrix row & col
		target_idx++
		row += drow
		col += dcol
	}
	// if the above loop terminated without returning, we matched the whole target_str
	return true
}

func is_x_match(char_matrix [][]rune, row, col int) bool {
	// check bounds (need to be one off from verical and horizontal borders)
	if row == 0 || row == len(char_matrix)-1 || col == 0 || col == len(char_matrix[0])-1 {
		return false
	}

	// if center isn't an 'A' fail
	if char_matrix[row][col] != 'A' {
		return false
	}
	// top row M M
	if char_matrix[row-1][col-1] == 'M' &&
		char_matrix[row-1][col+1] == 'M' &&
		char_matrix[row+1][col+1] == 'S' &&
		char_matrix[row+1][col-1] == 'S' {
		return true
	}

	if char_matrix[row-1][col-1] == 'S' &&
		char_matrix[row-1][col+1] == 'M' &&
		char_matrix[row+1][col+1] == 'M' &&
		char_matrix[row+1][col-1] == 'S' {
		return true
	}

	if char_matrix[row-1][col-1] == 'S' &&
		char_matrix[row-1][col+1] == 'S' &&
		char_matrix[row+1][col+1] == 'M' &&
		char_matrix[row+1][col-1] == 'M' {
		return true
	}

	if char_matrix[row-1][col-1] == 'M' &&
		char_matrix[row-1][col+1] == 'S' &&
		char_matrix[row+1][col+1] == 'S' &&
		char_matrix[row+1][col-1] == 'M' {
		return true
	}

	return false
}

func in_bounds(char_matrix [][]rune, row, col int) bool {
	// assumes fixed width rows, but input meets this condition
	if row < 0 || col < 0 || row >= len(char_matrix) || col >= len(char_matrix[0]) {
		return false
	}
	return true
}
