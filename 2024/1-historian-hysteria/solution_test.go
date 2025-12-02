package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string
var example = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestSolveExample(t *testing.T) {
	result := solve(example)
	assert.Equal(t, 11, result)
}

func TestSolveActual(t *testing.T) {
	result := solve(input)
	fmt.Println(result)
}

func TestSolve2Example(t *testing.T) {
	result := solve2(example)
	assert.Equal(t, 31, result)
}
func TestSolve2Actual(t *testing.T) {
	result := solve2(input)
	fmt.Println(result)
}