package util

import "strconv"

func MustParseInt(input string) int {
	result, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(result)
}
