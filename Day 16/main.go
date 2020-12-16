package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Requirement struct {
	text string
	limit1 int
	limit12 int
	limit21 int
	limit22 int
	possibleIndex []int
}

func Create(text string, range1, range2, range12, range22 int) *Requirement{
	tmp := Requirement{text: text, limit1: range1, limit12: range2, limit21: range12, limit22: range22}
	t := make([]int, 20)
	for i := 0; i < 20; i++ {
		t[i] = i
	}
	tmp.possibleIndex = t
	return &tmp
}

func main() {
	f, _ := os.Open("Day 16/input")
	file := bufio.NewScanner(f)

	require := []*Requirement{}
	validTickets := [][]int{}
	for file.Scan() {
		if file.Text() == ""{
			break
		}
		line := strings.Split(file.Text(), ": ")
		limits := strings.Split(line[1], " or ")
		lim1 := strings.Split(limits[0], "-")
		lim2 := strings.Split(limits[1], "-")
		num1, _ := strconv.Atoi(lim1[0])
		num2, _ := strconv.Atoi(lim1[1])
		num3, _ := strconv.Atoi(lim2[0])
		num4, _ := strconv.Atoi(lim2[1])
		require = append(require, Create(line[0], num1, num2, num3, num4))
	}

	file.Scan()
	file.Scan() // Move the scan line to my ticket
	myTicketNumbers := strings.Split(file.Text(), ",")
	myTicket := []int{}
	for _, number := range myTicketNumbers {
		num, _ := strconv.Atoi(number)
		myTicket = append(myTicket, num)
	}

	file.Scan()
	file.Scan() // Move the scan line to nearby tickets

	total := 0
	t := time.Now()
	for file.Scan() {
		stringNumbers := strings.Split(file.Text(), ",")
		numbers := []int{}
		for _, n := range stringNumbers {
			num, _ := strconv.Atoi(n)
			numbers = append(numbers, num)
		}
		if isValid(numbers, require, &total) {
			validTickets = append(validTickets, numbers)
		}
	}
	fmt.Println("Part 1 answer:", total)

	// Filter required fields indexes by checking each ticket and each of its numbers
	for _, req := range require {
		for i :=0 ; i< len(validTickets[0]); i++ {
			for row := 0; row < len(validTickets); row++ {
				if !meetsRequirements(validTickets[row][i], *req) {
					req.possibleIndex = remove(i, req.possibleIndex)
					break
				}
			}
		}
	}

	fieldCount := len(require)
	for {
		count := 0
		for _, req := range require {
			if len(req.possibleIndex) == 1 {
				for _, req2 := range require {
					if req.text != req2.text {
						req2.possibleIndex = remove(req.possibleIndex[0], req2.possibleIndex)
					}
				}
			}
			count += len(req.possibleIndex)
		}
		if count == fieldCount{
			break
		}
	}

	multiply := 1
	for _, req := range require {
		if strings.HasPrefix(req.text, "departure") {
			multiply *= myTicket[req.possibleIndex[0]]
		}
	}
	fmt.Println("Part 2 answer:",multiply)
	fmt.Println("Time taken:", time.Now().Sub(t))
}
func remove(number int, array []int) []int {
	tmp := []int{}
	for _, val := range array {
		if val != number {
			tmp = append(tmp, val)
		}
	}
	return tmp
}

func isValid(numbers []int, requirements []*Requirement, total *int) bool {
	for _, number := range numbers {
		isValid := false
		for _, req := range requirements {
			if number >= req.limit1 && number <= req.limit12 || number >= req.limit21 && number <= req.limit22 {
				isValid = true
				break
			}
		}
		if !isValid {
			*total += number
			return false
		}
	}
	return true
}

func meetsRequirements(number int, req Requirement) bool {
	if number >= req.limit1 && number <= req.limit12 || number >= req.limit21 && number <= req.limit22 {
		return true
	}
	return false
}