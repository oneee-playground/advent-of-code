package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string
var example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestSolveExample(t *testing.T) {
	result := solve(example)
	assert.Equal(t, 50, result)
}

func TestSolveActual(t *testing.T) {
	result := solve(input)
	fmt.Println(result)
}

func TestSolve2Example(t *testing.T) {
	result := solve2(example)
	assert.Equal(t, 24, result)
}

func TestSolve2Actual(t *testing.T) {
	result := solve2(input)
	fmt.Println(result)
}
