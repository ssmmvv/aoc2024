package problems

import (
	"bufio"
	"log"
	"os"

	"ssmmvv.github.io/aoc2024/util"
)

type Sector struct {
	idx, val, length int
}

func Problem9(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for idx, char := range line {
			row[idx] = util.MustParseInt(string(char))
		}
	}
}
