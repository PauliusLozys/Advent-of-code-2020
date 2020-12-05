package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

		code = strings.ReplaceAll(code, "F", "0")
		code = strings.ReplaceAll(code, "L", "0")
		code = strings.ReplaceAll(code, "B", "1")
		code = strings.ReplaceAll(code, "R", "1")
		ID, _ := strconv.ParseInt(code, 2, 32)
		id_array = append(id_array, int(ID))
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
