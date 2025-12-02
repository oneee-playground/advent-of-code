package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string
var example = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestSolveExample(t *testing.T) {
	result := solve(example)
	assert.Equal(t, 1227775554, result)
}

func TestSolveActual(t *testing.T) {
	result := solve(input)
	fmt.Println(result)
}

func TestSolve2Example(t *testing.T) {
	result := solve2(example)
	assert.Equal(t, 4174379265, result)
}

func TestSolve2Actual(t *testing.T) {
	result := solve2(input)
	fmt.Println(result)
}
