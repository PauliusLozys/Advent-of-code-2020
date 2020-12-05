package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {

	f, err := os.Open("Day 5/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	id_array := []int {}
	t := time.Now()
	for file.Scan() {
		code := file.Text()
		start_row := 0
		end_row := 127
		for i := 0; i < 7; i++ {
			if code[i] == 'F' { // Lower half
				end_row = (end_row - start_row) / 2 + start_row
			} else { // Upper half
				start_row = end_row - (end_row - start_row) / 2
			}
		}
		start_column := 0
		end_column := 7
		for i := 7; i < 10; i++ {
			if code[i] == 'L' { // Lower half
				end_column = (end_column - start_column) / 2 + start_column
			} else { // Upper half
				start_column = end_column - (end_column - start_column) / 2
			}
		}

		ID := start_row * 8 + start_column
		id_array = append(id_array, ID)
	}

	sort.Ints(id_array)
	fmt.Println("Part 1 answer:", id_array[len(id_array)-1])
	for i := 0; i < len(id_array) - 1; i++ {
		if (id_array[i] + 1) != id_array[i+1] {
			fmt.Println("Part 2 answer:", id_array[i] + 1)
			break
		}
	}
	fmt.Println("Time taken:", time.Now().Sub(t))
}
