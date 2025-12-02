package main

import (
	"strconv"
	"strings"
)

func solve(input string) int {
	instructions := strings.Split(input, "\n")

	dial := 50
	zeros := 0 // How many times did the dial stop at 0?

	for _, ins := range instructions {
		direction := ins[0]
		tick, _ := strconv.Atoi(ins[1:])

		switch direction {
		case 'R': // no-op.
		case 'L':
			tick = -tick
		}

		dial += tick
		dial = (dial + 100) % 100

		if dial == 0 {
			zeros++
		}
	}

	return zeros
}

func solve2(input string) int {
	instructions := strings.Split(input, "\n")

	dial := 50
	zeros := 0 // How many times did the dial tick at 0?

	for _, ins := range instructions {
		direction := ins[0]
		tick, _ := strconv.Atoi(ins[1:])

		switch direction {
		case 'R':
			dial += tick

			zeros += dial / 100
			dial %= 100
		case 'L':
			rotated := dial - tick

			metZeros := 0
			if rotated <= 0 {
				metZeros = (-rotated / 100) + 1 // +1 for first 0.

				zeros += metZeros
				if dial == 0 {
					// if start was 0, +1 doesn't count.
					zeros--
				}
			}

			dial = (rotated + (100 * metZeros)) % 100
		}

	}

	return zeros
}
