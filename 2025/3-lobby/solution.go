package main

import (
	"strconv"
	"strings"
)

func solve(input string) int {
	banks := strings.Split(input, "\n")

	sum := 0

	for _, bank := range banks {
		head := bank[len(bank)-2]
		tail := bank[len(bank)-1]

		for i := len(bank) - 3; i >= 0; i-- {
			digit := bank[i]
			if digit >= head {
				if tail < head {
					tail = head
				}
				head = digit
			}
		}

		num, _ := strconv.Atoi(string([]byte{head, tail}))
		sum += num
	}

	return sum
}

func solve2(input string) int {
	banks := strings.Split(input, "\n")
	sum := 0

	for _, bank := range banks {
		train := make([]byte, 12)
		copy(train, []byte(bank)[len(bank)-12:])

		for i := len(bank) - 13; i >= 0; i-- {
			digit := bank[i]
			push(train, digit)
		}

		num, _ := strconv.Atoi(string(train))
		sum += num
	}

	return sum
}

func push(train []byte, value byte) {
	if len(train) == 0 {
		return
	}

	if value >= train[0] {
		push(train[1:], train[0])
		train[0] = value
	}
}
