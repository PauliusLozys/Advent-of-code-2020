package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"strconv"
	"strings"
	"time"
)

type Storage struct {
	turn int
	firstTimeSpoken bool
}

func main() {

	input := "12,1,16,3,11,0"
	startingNum := strings.Split(input, ",")
	hash := hashmap.New()
	previousNumber := 0
	for turn, i := range startingNum {
		num, _ := strconv.Atoi(i)
		previousNumber = num
		hash.Put(num, Storage{turn: turn+1, firstTimeSpoken: true})
	}

	turn := len(startingNum)
	t := time.Now()
	for turn != 30000000 {
		nextNumber := -1
		tmp, _ := hash.Get(previousNumber)
		val := tmp.(Storage)

		if val.firstTimeSpoken{
			nextNumber = 0
			told , found := hash.Get(0)
			if !found { // 0 is a new value
				hash.Put(0, Storage{turn: turn+1, firstTimeSpoken: true})
			} else { // 0 already exists in the set
				old := told.(Storage)
				hash.Put(0, Storage{turn: old.turn, firstTimeSpoken: false})
			}
		} else {
			nextNumber = turn - val.turn
			told , found := hash.Get(nextNumber)
			if !found { // nextNumber is a new value
				hash.Put(nextNumber, Storage{turn: turn+1, firstTimeSpoken: true})
			} else { // nextNumber already exist in the set
				old := told.(Storage)
				hash.Put(nextNumber, Storage{turn: old.turn, firstTimeSpoken: false})
			}
			// Update the current number to his current turn as last turn
			hash.Put(previousNumber, Storage{turn: turn, firstTimeSpoken: false})
		}
		turn++
		previousNumber = nextNumber
		if turn == 2020 {
			fmt.Println("Part 1 answer:", previousNumber)
		}
	}
	fmt.Println("Part 2 answer:", previousNumber)
	fmt.Println("Time taken:", time.Now().Sub(t))
}