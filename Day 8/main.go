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

	f, err := os.Open("Day 8/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	instructions := []string{}
	for file.Scan() {
		line := file.Text()
		instructions = append(instructions, line)
	}
	accumulator := 0
	was_executed := make([]bool,len(instructions))
	instructions_to_test := []int{}
	execution_index := 0

	t := time.Now()
	for {
		if execution_index >= len(instructions) {
			fmt.Println("Reached the end of the program... quiting")
			fmt.Println("Part 2 accumulator value:", accumulator)
			break
		} else if was_executed[execution_index] {
			fmt.Println("Infinite loop detected... quiting")
			break
		}
		was_executed[execution_index] = true // Mark instruction as executed

		current_instruction := instructions[execution_index]
		command, argument := SplitInstruction(current_instruction)

		switch command {
		case "acc":
			accumulator += argument
			execution_index++
			break
		case "jmp":
			instructions_to_test = append(instructions_to_test, execution_index)
			execution_index += argument
			break
		case "nop":
			instructions_to_test = append(instructions_to_test, execution_index)
			execution_index++
			break
		}
	}
	fmt.Println("Part 1 accumulator value:", accumulator)

	to_test_index := 0
	succeesed := false
	for !succeesed {
		execution_index := 0
		accumulator := 0
		was_executed = make([]bool,len(instructions))
		for {
			if execution_index >= len(instructions) {
				succeesed = true
				fmt.Println("Reached the end of the program... quiting")
				fmt.Println("Part 2 accumulator value:", accumulator)
				//fmt.Println("Corrupted command index", instructions_to_test[to_test_index]+1)
				break
			} else if was_executed[execution_index] {
				break // Command swap failed
			}
			was_executed[execution_index] = true // Mark instruction as executed

			current_instruction := instructions[execution_index]
			command, argument := SplitInstruction(current_instruction)

			if execution_index == instructions_to_test[to_test_index]{
				if command == "jmp" {
					command = "nop"
				} else if command == "nop" {
					command = "jmp"
				}
			}

			switch command {
			case "acc":
				accumulator += argument
				execution_index++
				break
			case "jmp":
				execution_index += argument
				break
			case "nop":
				execution_index++
				break
			}
		}
		to_test_index++
	}
	fmt.Println("Time taken: ", time.Now().Sub(t))

}

func SplitInstruction(instruction string) (string, int){
	inst := strings.Split(instruction, " ")
	number, err := strconv.Atoi(inst[1])
	if err != nil {
		panic(err)
	}
	return inst[0], number
}