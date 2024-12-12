package problem10

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"ssmmvv.github.io/aoc2024/util"
)

type coord struct {
	row, col int
}

type path struct {
	row, col, height int
}

func Problem10(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for idx, char := range line {
			row[idx] = util.MustParseInt(string(char))
		}
		grid = append(grid, row)
	}

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				newTrails := getTrails(grid, row, col)
				fmt.Printf("%d new trails at %d,%d\n", newTrails, row, col)
				sum += newTrails
			}
		}
	}
	fmt.Printf("total trails: %d\n", sum)
}

func getTrails(grid [][]int, row, col int) int {
	visited := map[coord]bool{}
	visited[coord{row, col}] = true
	numTrails := 0
	frontier := []path{{row, col, 0}}
	for len(frontier) > 0 {
		curPath := frontier[0]
		frontier = frontier[1:]
		if grid[curPath.row][curPath.col] == 9 {
			numTrails++
			continue
		}

		// up
		newRow := curPath.row - 1
		newCol := curPath.col
		if inBounds(grid, newRow, newCol) && grid[newRow][newCol] == curPath.height+1 {
			frontier = append(frontier, path{newRow, newCol, curPath.height + 1})
		}

		// down
		newRow = curPath.row + 1
		newCol = curPath.col
		if inBounds(grid, newRow, newCol) && grid[newRow][newCol] == curPath.height+1 {
			frontier = append(frontier, path{newRow, newCol, curPath.height + 1})
		}

		// left
		newRow = curPath.row
		newCol = curPath.col - 1
		if inBounds(grid, newRow, newCol) && grid[newRow][newCol] == curPath.height+1 {
			frontier = append(frontier, path{newRow, newCol, curPath.height + 1})
		}

		// right
		newRow = curPath.row
		newCol = curPath.col + 1
		if inBounds(grid, newRow, newCol) && grid[newRow][newCol] == curPath.height+1 {
			frontier = append(frontier, path{newRow, newCol, curPath.height + 1})
		}

	}
	return numTrails

}

func inBounds(grid [][]int, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}
