package problem7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"ssmmvv.github.io/aoc2024/util"
)

func Problem7(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		components := strings.Split(line, ":")
		target := util.MustParseInt(components[0])
		equation := components[1]
		equation = equation[1:] // strip out the leading space
		equationComponents := strings.Split(equation, " ")
		nums := make([]int, len(equationComponents))
		for idx, numStr := range equationComponents {
			nums[idx] = util.MustParseInt(numStr)
		}

		if isValidEquation(target, nums) {
			sum += target
		}
	}
	fmt.Printf("sum: %d\n", sum)
}

// we recursively operate by picking either + or *, and calling isValidEquation with a new target
func isValidEquation(target int, nums []int) bool {
	// base case. If nums is one element that matches the target we've succeeded
	if len(nums) == 1 {
		return nums[0] == target
	}

	// try addition
	newNums := []int{}
	newNums = append(newNums, nums[0]+nums[1])
	newNums = append(newNums, nums[2:]...)
	if isValidEquation(target, newNums) {
		return true
	}

	// try multiplication
	newNums = []int{}
	newNums = append(newNums, nums[0]*nums[1])
	newNums = append(newNums, nums[2:]...)
	if isValidEquation(target, newNums) {
		return true
	}

	// try concatenation
	newNums = []int{}
	newNums = append(newNums, util.MustParseInt(fmt.Sprintf("%d%d", nums[0], nums[1])))
	newNums = append(newNums, nums[2:]...)
	return isValidEquation(target, newNums)
}
