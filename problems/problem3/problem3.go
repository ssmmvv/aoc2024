package problem3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"ssmmvv.github.io/aoc2024/util"
)

func Problem3(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	all_input := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		all_input += line
	}

	reg_exp := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	do_exp := regexp.MustCompile(`do\(\)`)
	dont_exp := regexp.MustCompile(`don't\(\)`)

	total := 0

	var do_idxs []int
	do_idxs = append(do_idxs, 0) // artificially add a 'do' at zero to make sure ops are enabled at the start
	var dont_idxs []int

	for _, match := range do_exp.FindAllStringIndex(all_input, -1) {
		do_idxs = append(do_idxs, match[0])
	}

	for _, match := range dont_exp.FindAllStringIndex(all_input, -1) {
		dont_idxs = append(dont_idxs, match[0])
	}

	fmt.Printf("do_idxs: ")
	fmt.Println(do_idxs)
	fmt.Printf("dont_idxs: ")
	fmt.Println(dont_idxs)
	matches := reg_exp.FindAllStringIndex(all_input, -1)
	for _, match := range matches {
		// we should have a bunch of digits and one comma between match[0] + 4 (inclusive) and match[1] - 1 (exclusive)
		if !is_enabled(do_idxs, dont_idxs, match[0]+1) {
			continue
		}
		args := all_input[match[0]+4 : match[1]-1]
		// fmt.Println(args)
		arg_components := strings.Split(args, ",")
		left := util.MustParseInt(arg_components[0])
		right := util.MustParseInt(arg_components[1])
		total += left * right
		fmt.Printf("%d x %d = %d, total is now %d\n", left, right, left*right, total)
	}

	fmt.Printf("total: %d\n", total)
}

// Returns true if operations are enabled at the given index.
func is_enabled(do_idxs, dont_idxs []int, op_idx int) bool {
	closest_do := 0
	cursor := 0
	for cursor < len(do_idxs) && do_idxs[cursor] < op_idx {
		cursor++
	}
	// in either termination case, we've actually advanced cursor just *past* the last "do"
	cursor--
	closest_do = do_idxs[cursor]

	closest_dont := 0
	cursor = 0
	for cursor < len(dont_idxs) && dont_idxs[cursor] < op_idx {
		cursor++
	}
	cursor--
	if cursor < 0 {
		return true
	}
	closest_dont = dont_idxs[cursor]
	return (closest_do > closest_dont)

}
