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
	f, err := os.Open("Day 9/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	preamble := 25
	numbers := make([]int, preamble)
	tmpCopy := make([]int, preamble) // Temporary dummy array for shifting "numbers" array to the left
	allNumbers := []int{}
	for file.Scan() {
		num := file.Text()
		converted, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		allNumbers = append(allNumbers, converted)
	}
	invalidNumber := -1
	numbers = allNumbers[:25]
	t := time.Now()
	for i := preamble; i < len(allNumbers);i++ {
		num := allNumbers[i]
		if !checkIfValid(&num, numbers){
			fmt.Println("Part 1 answer:", num)
			invalidNumber = num
			break
		}
		for j:=0; j < preamble - 1; j++ { // Shifting the array to the left
			tmpCopy[j] = numbers[j+1]
		}
		numbers = tmpCopy
		numbers[preamble - 1] = num
	}
	i, j := 0, 1
	total := allNumbers[i] + allNumbers[j]
	for total != invalidNumber {
		if total < invalidNumber{
			j++
			total += allNumbers[j]
		} else {
			total -= allNumbers[i]
			i++
			if i == j { // If the first index catches up to the last, we move it
				j++
				total += allNumbers[j]
			}
		}
	}
	tmp := allNumbers[i:j]
	sort.Ints(tmp)
	sum := tmp[0] + tmp[len(tmp)-1]
	fmt.Println("Part 2 answer:", sum)
	fmt.Println("Time taken:", time.Now().Sub(t))
}
func checkIfValid(number *int, numbersToCheckWith []int) bool{
	length := len(numbersToCheckWith)
	for i := 0; i < length - 1; i++{
		for j := i + 1; j < length; j++ {
			if *number == numbersToCheckWith[i] + numbersToCheckWith[j] {
				return true
			}
		}
	}
	return false
}