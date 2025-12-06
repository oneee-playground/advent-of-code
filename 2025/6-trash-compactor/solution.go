package main

import (
	"strconv"
	"strings"
)

func solve(input string) int {
	lines := strings.Split(input, "\n")

	var sheets [][]int
	for i := 0; ; i++ {
		numStrs := strings.Fields(lines[i])

		if i == 0 {
			sheets = make([][]int, len(numStrs))
		}

		if _, err := strconv.Atoi(numStrs[0]); err != nil {
			break
		}

		for idx, numStr := range numStrs {
			n, _ := strconv.Atoi(numStr)

			sheets[idx] = append(sheets[idx], n)
		}
	}

	sum := 0

	operators := strings.Fields(lines[len(lines)-1])
	for idx, op := range operators {
		base := 0
		sheet := sheets[idx]

		switch op {
		case "+":
			for _, num := range sheet {
				base += num
			}
		case "*":
			base = 1
			for _, num := range sheet {
				base *= num
			}
		}

		sum += base
	}

	return sum
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	idx := 0
	numCount := len(lines) - 1

	for idx < len(lines[0]) {
		base := 0
		operator := lines[numCount][idx]

		if operator == '*' {
			base = 1
		}

		for idx < len(lines[0]) {
			nextSheet := true

			num := 0

			for l := 0; l < numCount; l++ {
				line := lines[l]
				digit := line[idx]

				if digit == ' ' {
					continue
				}

				if num != 0 {
					num *= 10
				}

				nextSheet = false

				num += int(digit) - '0'
			}

			idx++

			if nextSheet {
				break
			}

			switch operator {
			case '*':
				base *= num
			case '+':
				base += num
			}
		}

		sum += base
	}

	return sum
}
