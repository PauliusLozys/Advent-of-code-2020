package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	f, err := os.Open("Day 10/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	allNumbers := []int{}
	allNumbers = append(allNumbers, 0)

	for file.Scan() {
		num := file.Text()
		converted, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		allNumbers = append(allNumbers, converted)
	}

	sort.Ints(allNumbers)

	max_jolts := allNumbers[len(allNumbers)-1] + 3
	allNumbers = append(allNumbers, max_jolts)
	t := time.Now()
	tmp := make([]int, len(allNumbers)-1)
	for i := 0;i < len(allNumbers)-1;i++{
		tmp[i] = allNumbers[i] - allNumbers[i+1]
	}
	count1 := 0
	count3 := 0

	for _, i := range tmp {
		if i == -1 {
			count1++
		} else if i == -3 {
			count3++
		}
	}
	//fmt.Println("Differences of 1 jolt:", count1)
	//fmt.Println("Differences of 3 jolt:", count3)
	fmt.Println("Part 1 answer:", count1 * count3)

	accum := map[int]int{0: 1}
	allNumbers = allNumbers[1:len(allNumbers)-1]
	for _, i := range allNumbers {
		accum[i] = accum[i-1] + accum[i-2] + accum[i-3]
	}
	fmt.Println("Part 2 answer:", accum[allNumbers[len(allNumbers)-1]])
	fmt.Println("Time taken:", time.Now().Sub(t))
}