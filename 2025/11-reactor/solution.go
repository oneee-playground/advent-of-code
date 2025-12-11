package main

import (
	"strings"
)

func solve(input string) int {
	lines := strings.Split(input, "\n")

	paths := make(map[string][]string, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)

		key := fields[0]
		key = key[:len(key)-1]

		paths[key] = fields[1:]
	}

	visited := make(map[string]int)

	var find func(label string) int
	find = func(label string) int {
		if n, ok := visited[label]; ok {
			return n
		}

		if label == "out" {
			return 1
		}

		sum := 0
		for _, next := range paths[label] {
			sum += find(next)
		}

		visited[label] = sum

		return sum
	}

	return find("you")
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")

	paths := make(map[string][]string, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)

		key := fields[0]
		key = key[:len(key)-1]

		paths[key] = fields[1:]
	}

	visited := make(map[string]int)

	tostr := func(b bool) string {
		if b {
			return "O"
		}
		return "X"
	}

	var find func(label string, dac, fft bool) int
	find = func(label string, dac, fft bool) int {
		hash := label + tostr(dac) + tostr(fft)
		if n, ok := visited[hash]; ok {
			return n
		}

		if label == "out" {
			if dac && fft {
				return 1
			}
			return 0
		}

		sum := 0
		for _, next := range paths[label] {
			sum += find(next, dac || label == "dac", fft || label == "fft")
		}

		visited[hash] = sum

		return sum
	}

	return find("svr", false, false)
}
