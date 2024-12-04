package main

import (
	"fmt"
	"os"

	"ssmmvv.github.io/aoc2024/problems"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Enter problem number and input path, e.g. \"1 inputs/1-test\"")
		return
	}
	switch os.Args[1] {
	case "1":
		problems.Problem1(os.Args[2])
		break
	case "2":
		problems.Problem2(os.Args[2])
		break
	case "3":
		problems.Problem3(os.Args[2])
		break
	}
}
