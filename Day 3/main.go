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
	length := len(pattern[0])
	val1 := 0
	val2 := 0
	val3 := 0
	val4 := 0
	val5 := 0

	row1 := 0
	row2 := 0
	row3 := 0
	row4 := 0

	t := time.Now()
	for i := 0; i < len(pattern); i++ { // 1 Down
		if pattern[i][i % length] == '#' {
			val1++
		}
		if pattern[i][row1 % length] == '#' {
			val2++
		}
		if pattern[i][row2 % length] == '#' {
			val3++
		}
		if pattern[i][row3 % length] == '#' {
			val4++
		}
		row1 += 3
		row2 += 5
		row3 += 7
	}


	for i := 0; i < len(pattern); i+=2 { // 2 Down
		if pattern[i][row4 % length] == '#' {
			val5++
		}
		row4++
	}

	fmt.Println("Time taken:",time.Now().Sub(t))
	fmt.Println("Part 1 answer:", val1)
	fmt.Println("Part 2 answer:", val1*val2*val3*val4*val5)
}
