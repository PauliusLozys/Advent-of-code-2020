package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	mainRules []int
	alternativeRules []int
	isCharacter bool
	character string
}

func main() {
	f, _ := os.Open("Day 19/input")
	file := bufio.NewScanner(f)
	rules := map[int]Rule{}
	parseRuleInput(file, rules)

	matches := 0
	for file.Scan() {
		if matcher(file.Text(), rules) {
			matches++
		}
		//fmt.Println(matchRule(file.Text(), 0, 0, rules))
	}
	fmt.Println(matches)
}
func matcher(line string, rules map[int]Rule) bool {
	colorRed := "\033[31m"
	isValid, checkedLenght := matchRule(line, 0, 0, rules)
	if isValid && checkedLenght == len(line)-1 {
		fmt.Println(string(colorRed), line, checkedLenght)

		return true
	}
	return false
}

func matchRule(line string, index int, rule int, rules map[int]Rule) (bool, int) {
	currentRule := rules[rule]
	if currentRule.isCharacter {
		if index >= len(line) {
			return false, index
		}
		return string(line[index]) == currentRule.character, index
	}
	matchedMain := true
	deepestIndex := index
	for _, rule := range currentRule.mainRules {
		state, i := matchRule(line, deepestIndex, rule, rules)
		if !state {
			matchedMain = false

			break
		} else {
			deepestIndex = i +1
		}
	}
	if rule == 0 && !matchedMain{
		return false, deepestIndex
	}
	if !matchedMain && len(currentRule.alternativeRules) != 0 {
		for _, rule := range currentRule.alternativeRules {
			state, i := matchRule(line, deepestIndex, rule, rules)
			if !state {
				return false, index
			} else {
				deepestIndex = i +1
			}
		}
	} else if !matchedMain && len(currentRule.alternativeRules) == 0{
		return false, index
	}

	return true, deepestIndex-1
}

func parseRuleInput(file *bufio.Scanner, rules map[int]Rule) {
	for file.Scan() {
		if file.Text() == "" {
			break
		}
		tokens := strings.Split(file.Text(), ": ")
		ruleNumberTmp := tokens[0]
		ruleNumber, _ := strconv.Atoi(ruleNumberTmp)
		newRule := Rule{}
		if strings.Contains(tokens[1], "|") {
			rules := strings.Split(tokens[1], " | ")
			for _, i := range strings.Split(rules[0], " ") {
				num, _ := strconv.Atoi(i)
				newRule.mainRules = append(newRule.mainRules, num)
			}
			for _, i := range strings.Split(rules[1], " ") {
				num, _ := strconv.Atoi(i)
				newRule.alternativeRules = append(newRule.alternativeRules, num)
			}
		} else if strings.Contains(tokens[1], "\"") {
			newRule.isCharacter = true
			for _, i := range strings.Split(tokens[1], "\"") {
				if i != "" {
					newRule.character = i
				}
			}
		} else {
			for _, i := range strings.Split(tokens[1], " ") {
				num, _ := strconv.Atoi(i)
				newRule.mainRules = append(newRule.mainRules, num)
			}
		}
		rules[ruleNumber] = newRule
	}
}