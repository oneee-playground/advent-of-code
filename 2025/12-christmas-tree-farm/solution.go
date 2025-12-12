package main

import (
	"strconv"
	"strings"
)

func solve(input string) int {
	sections := strings.Split(input, "\n\n")

	presents := make([]int, len(sections)-1)
	for i, section := range sections[:len(sections)-1] {
		lines := strings.Split(section, "\n")

		cnt := 0
		for _, line := range lines[1:] {
			cnt += strings.Count(line, "#")
		}
		presents[i] = cnt
	}

	result := 0
	for _, line := range strings.Split(sections[len(sections)-1], "\n") {
		first, second, _ := strings.Cut(line, ": ")

		a, b, _ := strings.Cut(first, "x")
		w, _ := strconv.Atoi(a)
		h, _ := strconv.Atoi(b)
		size := w * h

		wants := make([]int, 0, len(presents))
		for digits := range strings.SplitSeq(second, " ") {
			n, _ := strconv.Atoi(digits)
			wants = append(wants, n)
		}

		cnt := 0
		for i, want := range wants {
			cnt += presents[i] * want
		}

		if size > cnt {
			result++
		}
	}

	return result
}

func solve2(input string) int {
	return 0
}
