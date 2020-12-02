package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("Day 2/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)

	fmt.Println("Calculating part 1 and 2")
	t := time.Now()
	validPart1 := 0
	validPart2 := 0
	for file.Scan() {
		num := file.Text()
		splits := strings.Split(num, " ")
		limits := strings.Split(splits[0], "-")
		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])

		letter := []rune(strings.TrimSuffix(splits[1], ":"))[0]
		password := []rune(splits[2])

		// Calculating part 1
		count := 0
		for i := 0; i < len(password); i++ {
			if letter == password[i] {
				count++
			}
		}
		if count >= min && count <= max {
			validPart1++
		}

		//Calculating part 2
		if password[min-1] != letter && password[max-1] == letter {
			validPart2++
		} else if password[min-1] == letter && password[max-1] != letter {
			validPart2++
		}
	}
	fmt.Println("Part 1 valid passwords:", validPart1)
	fmt.Println("Part 2 valid passwords:", validPart2)
	fmt.Println("Time taken:",time.Now().Sub(t))
}