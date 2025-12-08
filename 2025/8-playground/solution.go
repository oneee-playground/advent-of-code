package main

import (
	"maps"
	"slices"
	"strconv"
	"strings"
)

func solve(input string, connections int) int {
	lines := strings.Split(input, "\n")

	nodes := make([][3]int, len(lines))
	for idx, line := range lines {
		digits := strings.Split(line, ",")
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		z, _ := strconv.Atoi(digits[2])

		nodes[idx] = [3]int{x, y, z}
	}

	edges := make([][3]int, 0)
	for i, node1 := range nodes {
		for j, node2 := range nodes[i+1:] {
			// It's not sqrt'ed but we only use it to compare.
			distance := 0
			for n := range 3 {
				distance += (node1[n] - node2[n]) * (node1[n] - node2[n])
			}

			edges = append(edges, [3]int{distance, i, i + 1 + j})
		}
	}

	slices.SortFunc(edges, func(a, b [3]int) int {
		return a[0] - b[0]
	})

	// kruskal algorithm.
	root := make([]int, len(nodes))
	for i := range root {
		root[i] = i
	}

	// union-find
	find := func(n int) int {
		var do func(int) int
		do = func(n int) int {
			if root[n] != n {
				root[n] = do(root[n])
			}

			return root[n]
		}
		return do(n)
	}
	union := func(a, b int) {
		a, b = find(a), find(b)
		root[b] = a
	}

	connectionsLeft := connections
	for _, edge := range edges {
		if connectionsLeft == 0 {
			break
		}

		union(edge[1], edge[2])

		connectionsLeft--
	}

	circuits := make(map[int]int)
	for _, n := range root {
		circuits[find(n)]++
	}

	list := slices.SortedFunc(maps.Values(circuits), func(a, b int) int {
		return b - a
	})

	sum := 1
	for i := range 3 {
		sum *= list[i]
	}

	return sum
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")

	nodes := make([][3]int, len(lines))
	for idx, line := range lines {
		digits := strings.Split(line, ",")
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		z, _ := strconv.Atoi(digits[2])

		nodes[idx] = [3]int{x, y, z}
	}

	edges := make([][3]int, 0)
	for i, node1 := range nodes {
		for j, node2 := range nodes[i+1:] {
			// It's not sqrt'ed but we only use it to compare.
			distance := 0
			for n := range 3 {
				distance += (node1[n] - node2[n]) * (node1[n] - node2[n])
			}

			edges = append(edges, [3]int{distance, i, i + 1 + j})
		}
	}

	slices.SortFunc(edges, func(a, b [3]int) int {
		return a[0] - b[0]
	})

	// kruskal algorithm.
	root := make([]int, len(nodes))
	for i := range root {
		root[i] = i
	}

	// union-find
	find := func(n int) int {
		var do func(int) int
		do = func(n int) int {
			if root[n] != n {
				root[n] = do(root[n])
			}

			return root[n]
		}
		return do(n)
	}
	union := func(a, b int) {
		a, b = find(a), find(b)
		root[b] = a
	}

	done := func() bool {
		base := find(root[0])
		for _, n := range root[1:] {
			if find(n) != base {
				return false
			}
		}
		return true
	}

	var lastEdge [3]int
	for _, edge := range edges {
		union(edge[1], edge[2])

		if done() {
			lastEdge = edge
			break
		}
	}

	node1 := nodes[lastEdge[1]]
	node2 := nodes[lastEdge[2]]

	return node1[0] * node2[0]
}
