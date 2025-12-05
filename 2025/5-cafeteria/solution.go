package main

import (
	"slices"
	"strconv"
	"strings"
)

func solve(input string) int {
	inputFresh, inputAvail, _ := strings.Cut(input, "\n\n")

	ranges := make([][2]int, 0)
	for _, line := range strings.Split(inputFresh, "\n") {
		start, end, _ := strings.Cut(line, "-")

		startN, _ := strconv.Atoi(start)
		endN, _ := strconv.Atoi(end)

		ranges = append(ranges, [2]int{startN, endN})
	}

	cnt := 0

	for _, digits := range strings.Split(inputAvail, "\n") {
		num, _ := strconv.Atoi(digits)

		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				cnt++
				break
			}
		}

	}

	return cnt
}

func solve2(input string) int {
	inputFresh, _, _ := strings.Cut(input, "\n\n")

	ranges := make([][2]int, 0)
	for _, line := range strings.Split(inputFresh, "\n") {
		start, end, _ := strings.Cut(line, "-")

		startN, _ := strconv.Atoi(start)
		endN, _ := strconv.Atoi(end)

		ranges = append(ranges, [2]int{startN, endN})
	}

	// Sort by ascending start, ascending end.
	slices.SortFunc(ranges, func(a, b [2]int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	read := 0
	cnt := 0

	for _, r := range ranges {
		read = max(read, r[0])

		if read > r[1] {
			continue
		}

		diff := r[1] - read + 1
		cnt += diff

		read = r[1] + 1
	}

	return cnt
}
