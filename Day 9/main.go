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
		if !checkIfValid(num, numbers){
			fmt.Println("Part 1 answer:", num)
			invalidNumber = num
			break
		}
		tmp := make([]int, len(numbers))
		for j:=0; j < len(numbers) - 1; j++ {
			tmp[j] = numbers[j+1]
		}
		numbers = tmp
		numbers[len(numbers) - 1] = num
	}
	for i := 0; i < len(allNumbers) - 1;i++ {
		total := allNumbers[i]
		for j := i + 1; j < len(allNumbers);j++ {
			total += allNumbers[j]
			if total == invalidNumber{
				tmp := allNumbers[i:j]
				sort.Ints(tmp)
				sum := tmp[0] + tmp[len(tmp)-1]
				fmt.Println("Part 2 answer:", sum)
				i = len(allNumbers)
				j = len(allNumbers)
			} else if total > invalidNumber { // Invalid set found, check next
				break
			}
		}
	}
	fmt.Println("Time taken:", time.Now().Sub(t))
}

func checkIfValid(number int, numbersToCheckWith []int) bool{
	length := len(numbersToCheckWith)
	for i := 0; i < length - 1; i++{
		for j := i + 1; j < length; j++ {
			if number == numbersToCheckWith[i] + numbersToCheckWith[j] {
				return true
			}
		}
	}
	return false
}