package problem8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord struct {
	row, col int
}

func Problem8(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennaLocations := map[rune][]Coord{}
	numRows := 0
	numCols := 0
	for scanner.Scan() {
		line := scanner.Text()
		numCols = len(line) // should just do for first line
		for idx, curRune := range line {
			if curRune != '.' {
				coordList, present := antennaLocations[curRune]
				if !present {
					antennaLocations[curRune] = []Coord{{numRows, idx}}
				} else {
					coordList = append(coordList, Coord{numRows, idx})
					antennaLocations[curRune] = coordList
				}
			}
		}
		numRows++ // since we don't incr until here, numRows == rowIdx in above code
	}

	// iterate
	numNullPoints := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if isNullPoint(antennaLocations, row, col) {
				fmt.Printf("%d, %d is a null point\n", row, col)
				numNullPoints++
			}
		}
	}

	fmt.Printf("num null points: %d\n", numNullPoints)
}

// to determine if there's a null point:
// iterate over all frequencies
//
//			read all antenna coordinates
//				get the displacement from input coord
//	         is there an antenna at displacement * 2?
//					return true (doesn't matter which frequency it is)
//
// return false
func isNullPoint(antennaLocations map[rune][]Coord, row, col int) bool {
	for _, coords := range antennaLocations {
		// create set of coordinates
		antennaMap := map[Coord]bool{}
		if row == 1 && col == 7 {
			fmt.Println("Debiging!")
		}
		for _, antennaCoord := range coords {
			antennaMap[Coord{antennaCoord.row, antennaCoord.col}] = true
		}

		for testAntenna := range antennaMap {
			d_row := testAntenna.row - row
			d_col := testAntenna.col - col

			if d_col == 0 && d_row == 0 {
				// if there's at least 2 antennas of the same freq, both antennas are a node
				return len(antennaMap) > 2
			}

			normAntenna1 := normalize(d_row, d_col)

			// for every other antenna of the same frequency
			for testAntenna2 := range antennaMap {
				if testAntenna2 == testAntenna {
					continue // skip same antenna
				}
				d_row_2 := testAntenna2.row - row
				d_col_2 := testAntenna2.col - col
				normAntenna2 := normalize(d_row_2, d_col_2)

				// are the antennas in line?
				if normAntenna1.row == normAntenna2.row && normAntenna1.col == normAntenna2.col {
					return true
				}
			}

		}
	}
	return false
}

func euclidAlgo(left, right int) int {
	if left < 0 {
		left = -left
	}
	if right < 0 {
		right = -right
	}
	if left == 0 {
		return right
	}
	if right == 0 {
		return left
	}
	if right == left {
		return left
	}

	if left > right {
		return euclidAlgo(left-right, right)
	} else {
		return euclidAlgo(left, right-left)
	}
}

func normalize(in_row, in_col int) Coord {
	row := in_row
	col := in_col

	if row < 0 {
		row = -row
		col = -col
	}
	if row == 0 {
		return Coord{0, 1}
	}
	if col == 0 {
		return Coord{1, 0}
	}

	denominator := euclidAlgo(row, col)
	return Coord{row / denominator, col / denominator}
}
