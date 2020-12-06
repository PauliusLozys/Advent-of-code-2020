package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {

	f, err := os.Open("Day 6/input")
	if err != nil {
		panic(err)
	}
	answers_groups_part1 := make([]int, 26)
	answers_groups_part2 := make([]int, 26)

	file := bufio.NewScanner(f)
	stringBuffer := bytes.Buffer{}
	t := time.Now()
	for file.Scan() {
		line := file.Text()
		if line == ""{
			countAnswersPart1(stringBuffer, answers_groups_part1)
			countAnswersPart2(stringBuffer, answers_groups_part2)
			stringBuffer.Reset()
		} else {
			stringBuffer.WriteString(line)
			stringBuffer.WriteRune(' ')
		}
	}
	countAnswersPart1(stringBuffer, answers_groups_part1)
	countAnswersPart2(stringBuffer, answers_groups_part2)

	total := 0
	total2 := 0
	for i := 0; i < 26; i++ {
		total += answers_groups_part1[i]
		total2 += answers_groups_part2[i]
	}
	fmt.Println("Part 1 answer",total)
	fmt.Println("Part 2 answer", total2)
	fmt.Println("Time taken:", time.Now().Sub(t))
}

func countAnswersPart1(stringBuffer bytes.Buffer, answer_group []int) {
	str := stringBuffer.String()
	arr := make([]uint8, 26)
	for _, item := range str{
		if unicode.IsLetter(item){
			index := item - 97
			if arr[index] != 1{
				arr[index]++
			}
		}
	}
	for idx, item := range arr {
		answer_group[idx] += int(item)
	}
}

func countAnswersPart2(stringBuffer bytes.Buffer, answer_group []int) {
	str := stringBuffer.String()
	people := strings.Split(str, " ")
	arr := make([]uint8, 26)
	count := len(people)-1
	for i := 0; i < count; i++ {
		for _, item := range people[i]{
			if unicode.IsLetter(item){
				index := item - 97
				arr[index]++
			}
		}
		for idx, item := range arr {
			if item == uint8(count) {
				answer_group[idx] += 1
			}
		}
	}
}