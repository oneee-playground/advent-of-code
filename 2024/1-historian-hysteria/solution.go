package main

import (
	"slices"
	"strconv"
	"strings"
)

func solve(input string) int {
	pairs := strings.Split(input, "\n")

	var left, right []int
	for _, pair := range pairs {
		leftDigit, rightDigit, _ := strings.Cut(pair, "   ")
		leftN, _ := strconv.Atoi(leftDigit)
		rightN, _ := strconv.Atoi(rightDigit)

		left = append(left, leftN)
		right = append(right, rightN)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for idx, leftN := range left {
		rightN := right[idx]

		distance := leftN - rightN
		if distance < 0 {
			distance = -distance
		}

		sum += distance
	}

	return sum
}

func solve2(input string) int {
	pairs := strings.Split(input, "\n")

	var left []int
	right := make(map[int]int, 0)

	for _, pair := range pairs {
		leftDigit, rightDigit, _ := strings.Cut(pair, "   ")
		leftN, _ := strconv.Atoi(leftDigit)
		rightN, _ := strconv.Atoi(rightDigit)

		left = append(left, leftN)
		right[rightN]++
	}

	sum := 0
	for _, leftN := range left {
		score := leftN * right[leftN]

		sum += score
	}

	return sum
}
