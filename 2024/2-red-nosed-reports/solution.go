package main

import (
	"strconv"
	"strings"
)

func solve(input string) int {
	reports := strings.Split(input, "\n")

	cnt := 0
	for _, report := range reports {
		levelsStr := strings.Split(report, " ")
		levels := make([]int, len(levelsStr))
		for idx, level := range levelsStr {
			n, _ := strconv.Atoi(level)
			levels[idx] = n
		}

		if isSafe(levels) {
			cnt++
		}
	}

	return cnt
}

func isSafe(levels []int) bool {
	initial := levels[1]-levels[0] > 0

	for idx, level := range levels[1:] {
		diff := level - levels[idx]

		if diff > 0 != initial {
			return false
		}

		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func solve2(input string) int {
	reports := strings.Split(input, "\n")

	cnt := 0
	for _, report := range reports {
		levelsStr := strings.Split(report, " ")
		levels := make([]int, len(levelsStr))
		for idx, level := range levelsStr {
			n, _ := strconv.Atoi(level)
			levels[idx] = n
		}

		if isSafe(levels) {
			cnt++
			continue
		}

		temp := make([]int, len(levels)-1)
		for toRemove := range len(levels) {
			copy(temp, levels[:toRemove])
			copy(temp[toRemove:], levels[toRemove+1:])
			if isSafe(temp) {
				cnt++
				break
			}
		}

	}

	return cnt
}
