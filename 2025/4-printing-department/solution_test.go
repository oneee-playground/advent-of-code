package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string
var example = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestSolveExample(t *testing.T) {
	result := solve(example)
	assert.Equal(t, 13, result)
}

func TestSolveActual(t *testing.T) {
	result := solve(input)
	fmt.Println(result)
}

func TestSolve2Example(t *testing.T) {
	result := solve2(example)
	assert.Equal(t, 43, result)
}

func TestSolve2Actual(t *testing.T) {
	result := solve2(input)
	fmt.Println(result)
}
