package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func solve(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0

	for _, line := range lines {
		fields := strings.Fields(line)

		lightDiagram := fields[0]
		lightDiagram = lightDiagram[1 : len(lightDiagram)-1]

		buttonsStr := fields[1 : len(fields)-1]

		lights := make([]bool, len(lightDiagram))
		expected := make([]bool, len(lightDiagram))
		for i := 0; i < len(lightDiagram); i++ {
			ch := lightDiagram[i]
			if ch == '#' {
				expected[i] = true
			}
		}

		buttons := make([][]int, len(buttonsStr))
		for i, str := range buttonsStr {
			digits := strings.Split(str[1:len(str)-1], ",")
			for _, digit := range digits {
				n, _ := strconv.Atoi(digit)
				buttons[i] = append(buttons[i], n)
			}
		}

		lightsOn := func() bool {
			return slices.Equal(expected, lights)
		}

		var permutate func(idx, pushed int) int
		permutate = func(idx, pushed int) int {
			if lightsOn() {
				// fmt.Println(pushed)
				// fmt.Println(memo)
				return pushed
			}
			if idx == len(buttons) {
				return math.MaxInt
			}

			minPushes := math.MaxInt

			doPush := func() {
				for _, i := range buttons[idx] {
					lights[i] = !lights[i]
				}
			}
			doPush()
			minPushes = min(minPushes, permutate(idx+1, pushed+1))
			doPush()
			minPushes = min(minPushes, permutate(idx+1, pushed))

			return minPushes
		}

		sum += permutate(0, 0)
	}

	return sum
}

func solve2(input string) int {
	// THIS IS SO SLOW. THEY SAY I SHOULD USE ILP RESOLVER.
	// BUT I DON'T KNOW DAMN ABOUT IT SO I'll JUST TAKE ADVANTAGE OF MULTICORE.
	// it's not finished even after hours. I need to come back to this later with some more knowledge.

	lines := strings.Split(input, "\n")

	var sum atomic.Uint64
	var wg sync.WaitGroup

	for _, line := range lines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fields := strings.Fields(line)

			joltageDiagram := fields[len(fields)-1]
			buttonsStr := fields[1 : len(fields)-1]

			expected := make([]int, 0)
			digits := strings.Split(joltageDiagram[1:len(joltageDiagram)-1], ",")
			for _, digit := range digits {
				n, _ := strconv.Atoi(digit)
				expected = append(expected, n)
			}
			actual := make([]int, len(expected))

			buttons := make([][]int, len(buttonsStr))
			for i, str := range buttonsStr {
				digits := strings.Split(str[1:len(str)-1], ",")
				for _, digit := range digits {
					n, _ := strconv.Atoi(digit)
					buttons[i] = append(buttons[i], n)
				}
			}

			inspect := func() (equal, over bool) {
				if slices.Equal(actual, expected) {
					return true, false
				}

				for i, n := range actual {
					if n > expected[i] {
						return false, true
					}
				}
				return false, false
			}

			var permutate func(idx, pushed int) int
			permutate = func(idx, pushed int) int {
				if idx == len(buttons) {
					return math.MaxInt
				}

				minPushes := math.MaxInt

				doPush := func(n int) {
					for _, i := range buttons[idx] {
						actual[i] += n
					}
				}

				for i := 0; ; i++ {
					equal, over := inspect()
					if equal {
						// fmt.Println(pushed + i)
						minPushes = min(minPushes, pushed+i)
					}
					if equal || over {
						doPush(-i)
						return minPushes
					}

					if minPushes <= pushed+i {
						return minPushes
					}

					minPushes = min(minPushes, permutate(idx+1, pushed+i))
					doPush(1)
				}
			}

			result := permutate(0, 0)

			sum.Add(uint64(result))
		}()
	}
	wg.Wait()

	return int(sum.Load())
}
