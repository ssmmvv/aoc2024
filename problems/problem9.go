package problems

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)

	freeSectors := []*Sector{}
	valSectors := []*Sector{}
	fileSlice := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		for idx, char := range line {
			numChars := util.MustParseInt(string(char))
			if idx%2 == 0 { // even idx is file
				fileId := idx / 2
				valSectors = append(valSectors, &Sector{len(fileSlice), fileId, numChars})
				for i := 0; i < numChars; i++ {
					fileSlice = append(fileSlice, fileId)
				}
			} else {
				freeSectors = append(freeSectors, &Sector{len(fileSlice), -1, numChars})
				for i := 0; i < numChars; i++ {
					fileSlice = append(fileSlice, -1)
				}
			}
		}
	}

	for _, val := range fileSlice {
		if val >= 0 {
			fmt.Printf("%d", val)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("\n")

	// non-empty sectors, right to left
COALESCE:
	for idx := len(valSectors) - 1; idx >= 0; idx-- {
		valSector := valSectors[idx]

		// find the leftmost free sector that fits
		for _, freeSector := range freeSectors {
			if freeSector.length >= valSector.length && freeSector.idx < valSector.idx {
				for offset := 0; offset < valSector.length; offset++ {
					fileSlice[freeSector.idx+offset] = valSector.val
					fileSlice[valSector.idx+offset] = -1
				}
				freeSector.length = freeSector.length - valSector.length
				freeSector.idx = freeSector.idx + valSector.length
				continue COALESCE
			}
		}
	}

	// Part 1 code
	// // right cursor points to swap src idx
	// right_cursor := len(fileSlice) - 1
	// for fileSlice[right_cursor] == -1 {
	// 	right_cursor--
	// }

	// // left cursor points to swap dest idx
	// left_cursor := 0
	// for fileSlice[left_cursor] != -1 {
	// 	left_cursor++
	// }

	// // we swap and advance cursors until they meet
	// for right_cursor > left_cursor {
	// 	fileSlice[left_cursor] = fileSlice[right_cursor]
	// 	fileSlice[right_cursor] = -1

	// 	for fileSlice[right_cursor] == -1 {
	// 		right_cursor--
	// 	}

	// 	for fileSlice[left_cursor] != -1 {
	// 		left_cursor++
	// 	}
	// }
	for _, val := range fileSlice {
		if val >= 0 {
			fmt.Printf("%d", val)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("\n")
	sum := checksum(fileSlice)
	fmt.Printf("checksum: %d\n", sum)
}

func checksum(in []int) int {
	sum := 0
	for idx, id := range in {
		if id > 0 {
			sum += idx * id
		}
	}
	return sum
}
