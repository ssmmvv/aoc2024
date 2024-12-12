package main

import (
	"fmt"
	"os"

	"ssmmvv.github.io/aoc2024/problems/problem1"
	"ssmmvv.github.io/aoc2024/problems/problem10"
	"ssmmvv.github.io/aoc2024/problems/problem2"
	"ssmmvv.github.io/aoc2024/problems/problem3"
	"ssmmvv.github.io/aoc2024/problems/problem4"
	"ssmmvv.github.io/aoc2024/problems/problem5"
	"ssmmvv.github.io/aoc2024/problems/problem6"
	"ssmmvv.github.io/aoc2024/problems/problem7"
	"ssmmvv.github.io/aoc2024/problems/problem8"
	"ssmmvv.github.io/aoc2024/problems/problem9"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Enter problem number and input path, e.g. \"1 inputs/1-test\"")
		return
	}
	switch os.Args[1] {
	case "1":
		problem1.Problem1(os.Args[2])
	case "2":
		problem2.Problem2(os.Args[2])
	case "3":
		problem3.Problem3(os.Args[2])
	case "4":
		problem4.Problem4(os.Args[2])
	case "5":
		problem5.Problem5(os.Args[2])
	case "6":
		problem6.Problem6(os.Args[2])
	case "7":
		problem7.Problem7(os.Args[2])
	case "8":
		problem8.Problem8(os.Args[2])
	case "9":
		problem9.Problem9(os.Args[2])
	case "10":
		problem10.Problem10(os.Args[2])
	default:
		fmt.Printf("Problem %s not recognized\n", os.Args[1])
	}
}
