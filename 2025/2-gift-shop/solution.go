package main

import (
	"strconv"
	"strings"
)

func solve(input string) int {
	ranges := strings.Split(input, ",")

	sum := 0

	for _, ran := range ranges {
		first, second, _ := strings.Cut(ran, "-")
		start, _ := strconv.Atoi(first)
		end, _ := strconv.Atoi(second)

		for cur := start; cur <= end; cur++ {
			if !isValid(cur) {
				sum += cur
			}
		}
	}

	return sum
}

func isValid(id int) bool {
	str := strconv.Itoa(id)

	if len(str)%2 == 1 {
		return true
	}

	return str[:len(str)/2] != str[len(str)/2:]
}

func solve2(input string) int {
	ranges := strings.Split(input, ",")

	sum := 0

	for _, ran := range ranges {
		first, second, _ := strings.Cut(ran, "-")
		start, _ := strconv.Atoi(first)
		end, _ := strconv.Atoi(second)

		for cur := start; cur <= end; cur++ {
			if !isValid2(cur) {
				sum += cur
			}
		}
	}

	return sum
}

func isValid2(id int) bool {
	str := strconv.Itoa(id)

	for window := 1; window <= len(str)/2; window++ {
		if len(str)%window != 0 {
			continue
		}

		base := str[:window]
		for i := 1; ; i++ {
			start := window * i
			if start == len(str) {
				return false
			}
			compare := str[start : start+window]
			if base != compare {
				break
			}
		}
	}

	return true
}
