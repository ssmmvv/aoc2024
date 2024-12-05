package problems

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"ssmmvv.github.io/aoc2024/util"
)

func Problem5(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Characters in the key must be before/after the values,
	// e.g. mustBeBefore[1] = {2, 3} means 1 must be before either 1 or 2
	mustBeAfter := map[int][]int{}
	mustBeBefore := map[int][]int{}
	orders := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			components := strings.Split(line, "|")
			left := util.MustParseInt(components[0])
			right := util.MustParseInt(components[1])

			currentAfter, found := mustBeAfter[right]
			if !found {
				mustBeAfter[right] = []int{left}
			} else {
				currentAfter = append(currentAfter, left)
				mustBeAfter[right] = currentAfter
			}

			currentBefore, found := mustBeBefore[left]
			if !found {
				mustBeBefore[left] = []int{right}
			} else {
				currentBefore = append(currentBefore, right)
				mustBeBefore[left] = currentBefore
			}
		}

		// all inputs are > 1 item, a comma is always present
		if strings.Contains(line, ",") {
			components := strings.Split(line, ",")
			order := make([]int, len(components))
			for idx, component := range components {
				order[idx] = util.MustParseInt(component)
			}
			orders = append(orders, order)
		}
	}

	// Iterate through each order, checking against the rules
	// Do we need to use transient rules, e.g. infer that 1|2 and 2|3 means 3,1 is invalid?
	// A: I guess not
	middle_sum := 0
	for _, order := range orders {
		if isCorrectOrder(order, mustBeAfter, mustBeBefore) {
			fmt.Println("Order is correct: " + fmt.Sprint(order))
			//	middle_sum += order[len(order)/2]
		} else {
			correct_order := getCorrectOrder([]int{}, order, mustBeAfter, mustBeBefore)
			if correct_order == nil {
				fmt.Println("Order could not be corrected: " + fmt.Sprint(order))
			} else {
				fmt.Printf("%s corrected to %s\n", fmt.Sprint(order), fmt.Sprint(correct_order))
				middle_sum += correct_order[len(correct_order)/2]
			}
		}
	}
	fmt.Printf("middle sum: %d\n", middle_sum)
}

func isCorrectOrder(order []int, mustBeAfter, mustBeBefore map[int][]int) bool {
	for idx := range len(order) {
		// None of the chars before the index must have idx in mustBeAfter
		for before_idx := idx; before_idx >= 0; before_idx-- {
			if slices.Contains(mustBeBefore[order[idx]], order[before_idx]) {
				fmt.Printf("%d at index %d cannot come after, %d at index %d\n",
					order[idx], idx, order[before_idx], before_idx)
				return false
			}
		}

		// None of the chars after the indexed char must have it in mustBeBefore
		for after_idx := idx; after_idx < len(order); after_idx++ {
			if slices.Contains(mustBeAfter[order[idx]], order[after_idx]) {
				fmt.Printf("%d at index %d cannot come before, %d at index %d\n",
					order[idx], idx, order[after_idx], after_idx)
				return false
			}
		}
	}
	return true
}

func getCorrectOrder(existing_order, remaining_order []int, mustBeAfter, mustBeBefore map[int][]int) []int {

	// Selct a page that:

SEARCH_PAGES:
	for _, possible_page := range remaining_order {
		// none of the existing char must be after it
		for _, existing_page := range existing_order {
			if slices.Contains(mustBeBefore[possible_page], existing_page) ||
				slices.Contains(mustBeAfter[existing_page], possible_page) {
				continue SEARCH_PAGES
			}
		}

		// None of the remaining chars must be before it
		for _, remaining_page := range remaining_order {
			if remaining_page == possible_page {
				continue
			}

			if slices.Contains(mustBeBefore[remaining_page], possible_page) ||
				slices.Contains(mustBeAfter[possible_page], remaining_page) {
				continue SEARCH_PAGES
			}
		}

		// See if a solution is found using this page at this location
		new_remaining_order := []int{}
		for _, page := range remaining_order {
			if page != possible_page {
				new_remaining_order = append(new_remaining_order, page)
			}
		}

		new_existing_order := slices.Clone(existing_order)
		new_existing_order = append(new_existing_order, possible_page)

		// if we just used the last page in the order, we just append and return
		if len(new_remaining_order) == 0 {
			return new_existing_order
		}

		final_order := getCorrectOrder(new_existing_order, new_remaining_order, mustBeAfter, mustBeBefore)
		if remaining_order != nil {
			return final_order
		}
		// if final_order is nil, we couldn't assemble a valid order, so try the next possible page
	}
	return nil // no order was found
}
