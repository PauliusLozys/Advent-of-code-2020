package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {

	f, err := os.Open("Day 4/input")
	if err != nil {
		panic(err)
	}
	attributes_must := []string { "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validPassportsPart1 := 0
	validPassportsPart2 := 0
	file := bufio.NewScanner(f)
	stringBuffer := bytes.Buffer{}

	t := time.Now()
	for file.Scan() {
		line := file.Text()
		if line == ""{
			// Validate string
			validatePassword(stringBuffer, &validPassportsPart1, &validPassportsPart2, attributes_must)
			stringBuffer.Reset()
		} else {
			stringBuffer.WriteString(line)
			stringBuffer.WriteRune(' ')
		}
	}

	if stringBuffer.Len() != 0 { // Need to parse the last entry
		validatePassword(stringBuffer, &validPassportsPart1, &validPassportsPart2, attributes_must)
		stringBuffer.Reset()
	}

	fmt.Println("Part 1 valid password count:",validPassportsPart1)
	fmt.Println("Part 2 valid password count:",validPassportsPart2)
	fmt.Println("Time taken:", time.Now().Sub(t))
}

func validatePassword(stringBuffer bytes.Buffer, validPasswordCounter1, validPasswordCounter2 *int, attributes_must []string){
	entries := strings.Split(stringBuffer.String(), " ")
	requiredCountPart1 := 0
	requiredCountPart2 := 0
	for i := 0; i < len(entries) - 1; i++ { // Ignore the last empty entry
		pair := strings.Split(entries[i], ":") // [key, value]
		if contains(pair[0], attributes_must) {
			requiredCountPart1++
			if validateParameters(pair[0], pair[1]) {
				requiredCountPart2++
			}
		}
	}
	if requiredCountPart1 == len(attributes_must) {
		*validPasswordCounter1++
	}
	if requiredCountPart2 == len(attributes_must) {
		*validPasswordCounter2++
	}
}

func contains(item string, array []string) bool{
	for _, str := range array {
		if item == str {
			return true
		}
	}
	return false
}

func validateNumbers(str string, low, high int) bool{
	if len(str) != 4 {
		return false
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	if value < low || value > high {
		return false
	}
	return true
}

func validateParameters(key, value string) bool {
	switch key {
	case "byr":
		return validateNumbers(value, 1920, 2002)
	case "iyr":
		return validateNumbers(value, 2010, 2020)
	case "eyr":
		return validateNumbers(value, 2020, 2030)
	case "hgt":
		if strings.Contains(value, "cm"){
			value := strings.TrimRight(value,"cm")
			height, err := strconv.Atoi(value)
			if err != nil {
				return false
			}

			if height < 150 || height > 193 {
				return false
			}
			return true
		}
		if strings.Contains(value, "in"){
			value := strings.TrimRight(value,"in")
			height, err := strconv.Atoi(value)
			if err != nil {
				return false
			}

			if height < 59 || height > 76 {
				return false
			}
			return true
		}
		return false
	case "hcl":
		if len(value) != 7 || value[0] != '#' {
			return false
		}

		characters := strings.TrimPrefix(value, "#")
		for _, ch := range characters {
			if unicode.IsNumber(ch) {
				continue
			}
			if unicode.IsLetter(ch) {
				if ch < 'a' || ch > 'f' {
					return false
				}
			}
		}
		return true
	case "ecl":
		must := []string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		return contains(value, must)
	case "pid":
		if len(value) != 9 {
			return false
		}
		valid := true
		for _, ch := range value {
			if !unicode.IsNumber(ch) {
				valid = false
				break
			}
		}
		return valid
	default:
		break
	}
	return false
}
