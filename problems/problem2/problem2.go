package problem2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"ssmmvv.github.io/aoc2024/util"
)

func Problem2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	num_safe := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_safe := false
		line := scanner.Text()
		split_lines := strings.Split(line, " ")
		if is_safe(split_lines) {
			line_safe = true
		} else { // try skipping one element
			for cut := range len(split_lines) {
				var new_split []string
				new_split = append(new_split, split_lines[:cut]...)
				new_split = append(new_split, split_lines[cut+1:]...)
				if is_safe(new_split) {
					line_safe = true
					break
				}
			}
		}
		if line_safe {
			num_safe++
			fmt.Printf("%s is safe\n", line)
		} else {
			fmt.Printf("%s is unsafe\n", line)
		}
	}
	fmt.Printf("total safe: %d\n", num_safe)
}

func is_safe(split_lines []string) bool {
	is_safe := true
	last_val := util.MustParseInt(split_lines[0])
	new_val := util.MustParseInt(split_lines[1])
	last_diff := new_val - last_val
	if last_diff > 3 || last_diff < -3 || last_diff == 0 {
		return false
	}
	idx := 2
	for idx < len(split_lines) {
		last_val = new_val
		new_val = util.MustParseInt(split_lines[idx])
		new_diff := new_val - last_val

		// if direction changed, unsafe
		if (last_diff < 0 && new_diff >= 0) || (last_diff > 0 && new_diff <= 0) {
			is_safe = false
			break
		}

		// if magintude is greater than 2, unsafe
		if new_diff > 3 || new_diff < -3 {
			is_safe = false
			break
		}

		last_diff = new_diff
		idx++
	}
	return is_safe
}
