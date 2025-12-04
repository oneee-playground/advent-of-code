package main

import (
	"strings"
)

func solve(input string) int {
	diagram := strings.Split(input, "\n")

	cnt := 0

	for i, line := range diagram {
		for j, ch := range line {
			if ch != '@' {
				continue
			}

			if canAccess(diagram, i, j) {
				cnt++
			}
		}
	}

	return cnt
}

func canAccess(diagram []string, i, j int) bool {
	delta := [][2]int{
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	papers := 0
	for _, d := range delta {
		x, y := i+d[0], j+d[1]
		if x < 0 || x >= len(diagram[0]) {
			continue
		}
		if y < 0 || y >= len(diagram) {
			continue
		}

		if diagram[x][y] == '@' {
			papers++
		}
	}

	return papers < 4
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")
	diagram := make([][]byte, len(lines))
	for idx, line := range lines {
		diagram[idx] = []byte(line)
	}

	cnt := 0

	for {
		temp := 0

		for i, line := range diagram {
			for j, ch := range line {
				if ch != '@' {
					continue
				}

				if canAccess2(diagram, i, j) {
					diagram[i][j] = '.'
					temp++
				}
			}
		}
		cnt += temp

		if temp == 0 {
			break
		}
	}

	return cnt
}

func canAccess2(diagram [][]byte, i, j int) bool {
	delta := [][2]int{
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	papers := 0
	for _, d := range delta {
		x, y := i+d[0], j+d[1]
		if x < 0 || x >= len(diagram[0]) {
			continue
		}
		if y < 0 || y >= len(diagram) {
			continue
		}

		if diagram[x][y] == '@' {
			papers++
		}
	}

	return papers < 4
}
