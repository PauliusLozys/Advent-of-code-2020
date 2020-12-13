package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("Day 13/input")
	file := bufio.NewScanner(f)
	file.Scan()
	timestamp, _ := strconv.Atoi(file.Text())
	file.Scan()
	buses := strings.Split(file.Text(), ",")

	t := time.Now()
	part1(timestamp, buses)
	part2(buses)
	fmt.Println("Time taken:", time.Now().Sub(t))
}

func inverceOfN(N, mod int) int {
	m := N % mod
	multiplyer := 1
	if m != 1 {
		multiplyer++
		for {
			if m * multiplyer % mod == 1{
				break
			}
			multiplyer++
		}
	}
	return multiplyer
}
func part1(time int, buses []string) {
	minTime := math.MaxInt32
	busID := -1
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		busTime, _ := strconv.Atoi(bus)
		dif := ((time / busTime) * busTime + busTime) - time
		if dif < minTime {
			busID = busTime
			minTime = dif
		}

	}
	fmt.Println("Part 1 answer:", busID * minTime)
}
func part2(buses []string){
	freeTime := make([]int, len(buses))
	busList := []int{}
	busIndex := 0
	for _, bus := range buses {
		if bus == "x" {
			freeTime[busIndex]++
			continue
		}
		num, _ := strconv.Atoi(bus)
		busList = append(busList, num)
		busIndex++
	}
	freeTime = freeTime[:len(busList)] // Cut the unnecessary array part
	for idx, item := range freeTime{ // Create the free time offset
		if idx + 1 < len(freeTime){
			freeTime[idx+1] += item+1
		}
	}
	// Solve using chinese remainder theorem
	b := make([]int, len(freeTime))
	for idx, item := range busList{
		b[idx] = item - freeTime[idx]
	}
	b[0] = 0 // The first one will always be 0, because it will be the modulus of itself
	N := 1
	for _, item := range busList{
		N *= item
	}
	n := make([]int, len(b))
	for idx, item := range busList {
		n[idx] = N / item
	}

	x := make([]int, len(b))
	for idx, item := range n {
		x[idx] = inverceOfN(item, busList[idx])
	}
	total := 0

	for idx, item := range b {
		total += item * n[idx] * x[idx]
	}
	fmt.Println("Part 2 answer:",total % N)
}