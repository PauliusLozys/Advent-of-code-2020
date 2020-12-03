package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Open("Day 3/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	pattern := [][]rune{}

	for file.Scan() {
		line := file.Text()
		pattern = append(pattern, []rune(line))
	}
	row_length := len(pattern[0])
	array_length := len(pattern)

	// Move set
	right_moves := []int {3, 1, 5, 7, 1}
	down_moves :=  []int {1, 1, 1, 1, 2}


	total_moves := len(right_moves)
	row_indexes := make([]int, total_moves)
	column_indexes := make([]int, total_moves)
	val := make([]int, total_moves)

	fmt.Println("Calculating part 1 and 2")
	t := time.Now()
	for i := 0; i < array_length; i++ {
		for j := 0; j < total_moves; j++ {
			if row_indexes[j] <= array_length && pattern[row_indexes[j]][column_indexes[j] % row_length] == '#' {
				val[j]++
			}
			row_indexes[j] += down_moves[j]
			column_indexes[j]  += right_moves[j]
		}
	}

	total := 1
	for _, item := range val {
		total *= item
	}
	fmt.Println("Part 1 answer:", val[0])
	fmt.Println("Part 2 answer:", total)
	fmt.Println("Time taken:",time.Now().Sub(t))
}
