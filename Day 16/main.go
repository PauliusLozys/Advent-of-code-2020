package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Requirement struct {
	text string
	limit1 int
	limit12 int
	limit21 int
	limit22 int
}

func main() {
	f, _ := os.Open("Day 16/input")
	file := bufio.NewScanner(f)

	require := []Requirement{}
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
		require = append(require, Requirement{text: line[0], limit1: num1, limit12: num2, limit21: num3, limit22: num4})
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
	fmt.Println(total)

	discoveredField := map[string]int{}
	for index := 0; index < len(validTickets[0]); index++ {
		cope := map[string]int{}
		for key, val := range discoveredField{
			cope[key] = val
		}
		for _, req := range require {
			isValid := true
			for _, line := range validTickets {
				if !meetsRequirment(line[index], req) {
					isValid = false
				}
			}
			if isValid {
				discoveredField[req.text] += 1
			}
			for key, _ := range cope {
				delete(discoveredField, key)
			}
		}

		fmt.Println(index, discoveredField)
	}

	fmt.Println(discoveredField)
	multiply := 1

	in := []int{8, 4, 18, 9, 17, 7, 2, 14, 0, 13, 1, 16, 5, 12, 6, 3, 15, 11, 19, 10}
	for idx, index := range in {
		fmt.Println(myTicket[index], "-", require[idx].text, "=", index)
		if strings.HasPrefix(require[idx].text, "departure") {
			multiply *= myTicket[index]
		}
	}
	//fmt.Println(arr2)
	fmt.Println(multiply)

}
func contains(str string, array []string) (bool, int) {
	for idx, i := range array {
		if i == str {
			return true , idx
		}
	}
	return false, -1
}

func isValid(numbers []int, requirements []Requirement, total *int) bool {
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

func meetsRequirment(number int, req Requirement) bool {
	if number >= req.limit1 && number <= req.limit12 || number >= req.limit21 && number <= req.limit22 {
		return true
	}
	return false
}