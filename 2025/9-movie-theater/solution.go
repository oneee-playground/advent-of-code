package main

import (
	"strconv"
	"strings"
)

func solve(input string) int {
	lines := strings.Split(input, "\n")

	tiles := make([][2]int, len(lines))
	for idx, line := range lines {
		i, j, _ := strings.Cut(line, ",")
		x, _ := strconv.Atoi(i)
		y, _ := strconv.Atoi(j)

		tiles[idx] = [2]int{x, y}
	}

	space := 0
	for i, tile1 := range tiles {
		for _, tile2 := range tiles[i+1:] {
			width := tile1[0] - tile2[0]
			if width < 0 {
				width = -width
			}
			width++

			height := tile1[1] - tile2[1]
			if height < 0 {
				height = -height
			}
			height++

			space = max(space, width*height)
		}
	}

	return space
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")

	tiles := make([][2]int, len(lines))
	for idx, line := range lines {
		i, j, _ := strings.Cut(line, ",")
		x, _ := strconv.Atoi(i)
		y, _ := strconv.Atoi(j)

		tiles[idx] = [2]int{x, y}
	}

	connections := make([][2][2]int, len(tiles)-1)

	for i := 0; i < len(tiles)-1; i++ {
		tile1 := tiles[i]
		tile2 := tiles[i+1]

		connections = append(connections, [2][2]int{tile1, tile2})
	}
	connections = append(connections, [2][2]int{tiles[len(tiles)-1], tiles[0]})

	isInGreen := func(tile1, tile2 [2]int) bool {
		minx, miny := min(tile1[0], tile2[0]), min(tile1[1], tile2[1])
		maxx, maxy := max(tile1[0], tile2[0]), max(tile1[1], tile2[1])

		for _, conn := range connections {
			from, to := conn[0], conn[1]

			var idx, delta int

			for i := range 2 {
				switch {
				case from[i] < to[i]:
					idx = i
					delta = 1
				case from[i] > to[i]:
					idx = i
					delta = -1
				}
			}

			for tile := from; tile[idx] != to[idx]+delta; tile[idx] += delta {
				overlapX := minx < tile[0] && tile[0] < maxx
				overlapY := miny < tile[1] && tile[1] < maxy
				if overlapX && overlapY {
					return false
				}
			}

		}

		return true
	}

	space := 0
	for i, tile1 := range tiles {
		for _, tile2 := range tiles[i+1:] {
			if !isInGreen(tile1, tile2) {
				continue
			}

			width := tile1[0] - tile2[0]
			if width < 0 {
				width = -width
			}
			width++

			height := tile1[1] - tile2[1]
			if height < 0 {
				height = -height
			}
			height++

			space = max(space, width*height)
		}
	}
	return space
}
