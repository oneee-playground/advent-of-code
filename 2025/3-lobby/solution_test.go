package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string
var example = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestSolveExample(t *testing.T) {
	result := solve(example)
	assert.Equal(t, 357, result)
}

func TestSolveActual(t *testing.T) {
	result := solve(input)
	fmt.Println(result)
}

func TestSolve2Example(t *testing.T) {
	result := solve2(example)
	assert.Equal(t, 3121910778619, result)
}

func TestSolve2Actual(t *testing.T) {
	result := solve2(input)
	fmt.Println(result)
}
