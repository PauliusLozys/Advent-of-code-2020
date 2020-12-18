package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("Day 18/input")
	file := bufio.NewScanner(f)

	total := 0
	t := time.Now()
	for file.Scan() {
		final := recursiveParse(file.Text())
		num, err := strconv.Atoi(recursiveParse(final))
		if err != nil {
			panic(err)
		}
		total += num
	}
	fmt.Println("Answer:", total)
	fmt.Println("Time taken:", time.Now().Sub(t))
}

func recursiveParse(line string) string {

	tokens := parenthesesSplit(line)

	if len(tokens) > 1 { // Contains "()"
		for idx, str := range tokens {
			if strings.Contains(str, "(") {
				tokens[idx] = recursiveParse(str)
			}
		}
	}
	
	for idx, str := range tokens {
		//t, f := tryEvaluate(str) // Part 1
		t, f := tryEvaluateOrder(str) // Part 2
		if f {
			tokens[idx] = t
		}
	}
	stringBuffer := bytes.Buffer{}
	for _, i := range tokens {
		if i != ""{
			stringBuffer.WriteString(i)
		}
	}
	return fmt.Sprint(stringBuffer.String())
}
func tryEvaluate(str string) (string, bool) {
	t := strings.Split(str, " ")
	for i := 0; i < len(t) - 2; i+=2 {
		n1, f := strconv.Atoi(t[i])
		n2, f1 := strconv.Atoi(t[i+2])
		if f != nil || f1 != nil {
			return "", false
		}
		switch t[i+1] {
		case "+":
			t[i+2] = fmt.Sprint(n1 + n2)
			break
		case "*":
			t[i+2] = fmt.Sprint(n1 * n2)
			break
		default:
			return "", false
		}
	}
	return t[len(t)-1], true
}

func tryEvaluateOrder(str string) (string, bool) {
	t := strings.Split(str, " ")
	buildEq := bytes.Buffer{}
	for i := 0; i < len(t) - 2; i+=2 {
		if 	t[i+1] == "+" {
			n, f := tryEvaluate(t[i] + " + " + t[i+2])
			if !f {
				return "", false
			}
			t[i+2] = n
			t = append(t[:i], t[i+2:]...)
			i = -2
		}
	}
	for _, i := range t {
		if i != ""{
			if i == "*" {
				buildEq.WriteString(" * ")
			} else {
				buildEq.WriteString(i)
			}
		}
	}

	rez := buildEq.String()
	return tryEvaluate(rez)
}

func parenthesesSplit(line string) []string {
	tmp := []string{}
	nesting := []int{}
	firstIndex := 0
	for idx, ch := range line {

		if ch == '(' {
			nesting = append(nesting, 0)
			if len(nesting) == 1 {
				tmp = append(tmp, line[firstIndex:idx])
				firstIndex = idx
			}
		}

		if ch == ')' {
			if len(nesting) == 1 {
				tmp = append(tmp, line[firstIndex+1: idx])
				firstIndex = idx+1
				nesting = nesting[:len(nesting)-1]
			} else {
				nesting = nesting[:len(nesting)-1]
			}
		}


	}
	if len(tmp) == 0 {
		tmp = append(tmp, line)
	} else if len(line) > firstIndex+1 {
		tmp = append(tmp, line[firstIndex:])
	}
	return tmp
}