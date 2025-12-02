package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string
var example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestSolveExample(t *testing.T) {
	result := solve(example)
	assert.Equal(t, 3, result)
}

func TestSolveActual(t *testing.T) {
	result := solve(input)
	fmt.Println(result)
}

func TestSolve2Example(t *testing.T) {
	result := solve2(example)
	assert.Equal(t, 6, result)
}

func TestSolve2Actual(t *testing.T) {
	result := solve2(input)
	fmt.Println(result)
}
