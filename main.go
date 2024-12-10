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
	case "2":
		problems.Problem2(os.Args[2])
	case "3":
		problems.Problem3(os.Args[2])
	case "4":
		problems.Problem4(os.Args[2])
	case "5":
		problems.Problem5(os.Args[2])
	case "6":
		problems.Problem6(os.Args[2])
	case "7":
		problems.Problem7(os.Args[2])
	case "8":
		problems.Problem8(os.Args[2])
	default:
		fmt.Printf("Problem %s not recognized\n", os.Args[1])
	}
}
