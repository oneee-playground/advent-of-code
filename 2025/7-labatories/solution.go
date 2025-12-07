package main

import (
	"bytes"
)

func solve(input string) int {
	lines := bytes.Split([]byte(input), []byte{'\n'})

	beams := [][2]int{}
	for i := range lines[0] {
		if lines[0][i] == 'S' {
			beams = append(beams, [2]int{0, i})
			break
		}
	}

	cnt := 0

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		i, j := beam[0], beam[1]
		for {
			if lines[i][j] == '|' {
				break
			}

			lines[i][j] = '|'

			i++

			if i == len(lines) {
				break
			}

			if lines[i][j] == '^' {
				cnt++

				for _, d := range []int{-1, 1} {
					newj := j + d
					if newj < 0 || newj == len(lines[i]) {
						continue
					}

					beams = append(beams, [2]int{i, newj})
				}

				break
			}

		}

	}

	return cnt
}

func solve2(input string) int {
	lines := bytes.Split([]byte(input), []byte{'\n'})

	beam := [2]int{}
	for i := range lines[0] {
		if lines[0][i] == 'S' {
			beam = [2]int{0, i}
			break
		}
	}

	return shoot(lines, beam, make(map[[2]int]int))
}

func shoot(lines [][]byte, beam [2]int, memo map[[2]int]int) (timelines int) {
	i, j := beam[0], beam[1]

	for {
		i++
		if i == len(lines) {
			return 1
		}

		if lines[i][j] == '^' {
			if n := memo[[2]int{i, j}]; n != 0 {
				return n
			}

			sum := 0

			for _, d := range []int{-1, 1} {
				newj := j + d
				if newj < 0 || newj == len(lines[i]) {
					continue
				}
				sum += shoot(lines, [2]int{i, newj}, memo)
			}

			memo[[2]int{i, j}] = sum

			return sum
		}
	}
}
