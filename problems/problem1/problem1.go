package problem1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"ssmmvv.github.io/aoc2024/util"
)

func Problem1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left_counts := map[int]int{}
	right_counts := map[int]int{}

	scanner := bufio.NewScanner(file)
	var left, right []int
	for scanner.Scan() {
		line := scanner.Text()
		split_lines := strings.Split(line, "   ")
		left_int := util.MustParseInt(split_lines[0])
		right_int := util.MustParseInt(split_lines[1])
		left = append(left, int(left_int))
		right = append(right, int(right_int))

		left_count, found := left_counts[left_int]
		if !found {
			left_counts[left_int] = 1
		} else {
			left_counts[left_int] = left_count + 1
		}

		right_count, found := right_counts[right_int]
		if !found {
			right_counts[right_int] = 1
		} else {
			right_counts[right_int] = right_count + 1
		}
	}
	sort.Ints(left)
	sort.Ints(right)

	total_distance := 0
	for idx := range len(left) {
		distance := left[idx] - right[idx]
		if distance < 0 {
			distance = -distance
		}
		total_distance += distance
	}
	fmt.Printf("distance: %d\n", total_distance)

	similarity := 0
	for left_num, left_count := range left_counts {
		right_count, found := right_counts[left_num]
		if found {
			new_similarity := left_count * right_count * left_num
			similarity += new_similarity
		}
	}
	fmt.Printf("simialrity: %d\n", similarity)
}
