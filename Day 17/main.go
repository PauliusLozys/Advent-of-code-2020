package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	input, _ := ioutil.ReadFile("Day 17/input")
	initialStateWithW := parseInput(string(input))
	initialState := parseInput(string(input))
	initialLength := len(strings.Split(string(input), "\n")[0])
	initialHeight := len(strings.Split(string(input), "\n")) - 1

	t := time.Now()
	for i := 1; i <= 6; i++ {
		nextStateWithW := map[string]bool{}
		nextStateNoW := map[string]bool{}
		// Part 2
		for w := 0 -i; w <= 0+i; w++{
			for z := 0 - i; z <= 0+i; z++ {
				for x := 0 - i; x <= initialLength+i; x++ {
					for y := 0 - i; y <= initialHeight+i; y++ {
						newState := newState(w, x, y, z, initialStateWithW)
						nextStateWithW = setPosition(w, x,y,z, nextStateWithW, newState)
					}
				}
			}
		}

		// Part 1
		for z := 0 - i; z <= 0+i; z++ {
			for x := 0 - i; x <= initialLength+i; x++ {
				for y := 0 - i; y <= initialHeight+i; y++ {
					newState := newState(0, x, y, z, initialState)
					nextStateNoW = setPosition(0, x,y,z, nextStateNoW, newState)
				}
			}
		}
		initialState = nextStateNoW
		initialStateWithW = nextStateWithW
	}
	fmt.Printf("Part 1: %d\n", len(initialState))
	fmt.Printf("Part 2: %d\n", len(initialStateWithW))
	fmt.Println("Time taken:", time.Now().Sub(t))
}

func newState(w, x, y, z int, currentState map[string]bool) bool {
	activeNumbers := 0
TOP:
	for dx := x-1; dx <= x+1; dx++{
		for dy := y-1; dy <= y+1; dy++{
			for dz := z-1; dz <= z+1; dz++{
				for dw := w-1; dw <= w+1; dw++{
					if dw == w && dx == x && dy == y && dz == z {
						continue
					}
					if activeAtPosition(dw, dx, dy, dz, currentState) {
						activeNumbers++
						if activeNumbers > 3 {
							break TOP
						}
					}
				}
			}
		}
	}

	currentPosIsActive := activeAtPosition(w, x, y, z, currentState)
	if currentPosIsActive && (activeNumbers == 2 || activeNumbers == 3) {
		return true
	}
	if currentPosIsActive {
		return false
	}
	if !currentPosIsActive && activeNumbers == 3 {
		return true
	}
	return false
}

func activeAtPosition(w, x, y, z int, currentState map[string]bool) bool {
	return currentState[fmt.Sprintf("%d,%d,%d,%d", w, x, y, z)]
}

func setPosition(w, x, y, z int, currentState map[string]bool, newVal bool) map[string]bool {
	if newVal {
		currentState[fmt.Sprintf("%d,%d,%d,%d", w, x, y, z)] = true
	} else {
		delete(currentState, fmt.Sprintf("%d,%d,%d,%d", w, x, y, z))
	}
	return currentState
}

func parseInput(input string) map[string]bool {
	positions := map[string]bool{}
	z := 0
	w := 0
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if string(char) == "#" {
				positions[fmt.Sprintf("%d,%d,%d,%d", w, x, y, z)] = true
			}
		}
	}
	return positions
}