package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	Part1()
	Part2()
}
func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
func Part1 (){
	f, err := os.Open("Day 12/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	facingDirection := 2 // west - 0, north - 1, east - 2, south - 3
	north := 0
	west := 0
	south := 0
	east := 0

	for file.Scan() {
		line := file.Text()
		direction := line[:1]
		coord, _ := strconv.Atoi(line[1:])

		switch direction {
		case "N":
			north += coord
			break
		case "S":
			south += coord
			break
		case "E":
			east += coord
			break
		case "W":
			west += coord
			break
		case "L":
			facingDirection = (facingDirection - (coord / 90) +4) % 4
			break
		case "R":
			facingDirection = (facingDirection + (coord / 90) +4) % 4
			break
		case "F":
			switch facingDirection {
			case 0:
				west += coord
				break
			case 1:
				north += coord
				break
			case 2:
				east += coord
				break
			case 3:
				south += coord
				break
			}
		}
	}
	f.Close()
	absoluteX := Abs(west - east)
	absoluteY := Abs(north - south)
	fmt.Println("Part 1 ansewer: ", absoluteX + absoluteY)
}

func Part2 (){
	f, err := os.Open("Day 12/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	x := 0
	y := 0
	wx := 10
	wy := 1
	for file.Scan() {
		line := file.Text()
		direction := line[:1]
		coord, _ := strconv.Atoi(line[1:])

		switch direction {
		case "N":
			wy += coord
			break
		case "S":
			wy -= coord
			break
		case "E":
			wx += coord
			break
		case "W":
			wx -= coord
			break
		case "R":
			nx := float64(wx) * math.Cos(math.Pi * float64(coord) / 180) + float64(wy) * math.Sin(math.Pi * float64(coord) / 180)
			ny := -float64(wx) * math.Sin(math.Pi * float64(coord) / 180) + float64(wy) * math.Cos(math.Pi * float64(coord) / 180)
			wx = int(math.Round(nx))
			wy = int(math.Round(ny))

			break
		case "L":
			nx := float64(wx) * math.Cos(-math.Pi * float64(coord) / 180) + float64(wy) * math.Sin(-math.Pi * float64(coord) / 180)
			ny := -float64(wx) * math.Sin(-math.Pi * float64(coord) / 180) + float64(wy) * math.Cos(-math.Pi * float64(coord) / 180)
			wx = int(math.Round(nx))
			wy = int(math.Round(ny))
			break
		case "F":
			x += coord * wx
			y += coord * wy
		}
	}
	f.Close()
	absoluteX := math.Abs(float64(x))
	absoluteY := math.Abs(float64(y))
	fmt.Println("Part 2 ansewer: ", absoluteX + absoluteY )
}