package utils

import "strconv"

func StringToInt(f string) int {
	n, err := strconv.Atoi(f)
	if err != nil {
		panic(err)
	}
	return n
}
