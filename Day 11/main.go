package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	xLength = 0
	yLength = 0
)

func main() {
	f, err := os.Open("Day 11/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	grid := [][]int{}

	for file.Scan() {
		num := file.Text()
		line := []int{}
		for _, i := range num {
			line = append(line, int(i))
		}
		grid = append(grid, line)
	}
	xLength = len(grid)
	yLength = len(grid[0])
	t := time.Now()
	for {
		newGrid := [][]int{}
		for x, line := range grid{
			newLine := []int{}
			for y, symb := range line {
				symb := evaluateSeatPart1(x,y, symb, grid)
				//symb := evaluateSeatPart2(x,y, symb, allNumbers)
				newLine = append(newLine, symb)
			}
			newGrid = append(newGrid, newLine)
		}
		if checkIfSame(grid, newGrid) {
			fmt.Println("Same grid detected")
			break
		}
		grid = newGrid
	}
	fmt.Println("Answer:", checkForOccupiedSeats(grid))
	fmt.Println("Time taken:", time.Now().Sub(t))
}
func checkForOccupiedSeats(grid [][]int) int {
	count := 0
	for _, arr := range grid{
		for _, item := range arr {
			if item == '#' {
				count++
			}
		}
	}
	return count
}
func valid(x, y int) bool {
	return x < xLength && x >= 0 && y < yLength && y >= 0
}
func checkIfSame(oldGrid, newGrid [][]int) bool {
	for x, arr := range oldGrid{
		for y, item := range arr {
			if newGrid[x][y] != item {
				return false
			}
		}
	}
	return true
}
func evaluateSeatPart1(x, y, symb int, grid [][]int) int{
	if symb == '.' {
		return symb
	}
	occupations := 0

	if valid(x-1,y-1) && grid[x-1][y-1] == '#' {
		occupations++
	}
	if valid(x-1,y) && grid[x-1][y] == '#' {
		occupations++
	}
	if valid(x,y-1) && grid[x][y-1] == '#' {
		occupations++
	}
	if valid(x-1,y+1) &&  grid[x-1][y+1] == '#' {
		occupations++
	}
	if valid(x,y+1) && grid[x][y+1] == '#' {
		occupations++
	}
	if valid(x+1,y-1) && grid[x+1][y-1] == '#' {
		occupations++
	}
	if valid(x+1,y) && grid[x+1][y] == '#' {
		occupations++
	}
	if valid(x+1,y+1) && grid[x+1][y+1] == '#' {
		occupations++
	}

	if symb == 'L' && occupations == 0 {
		return '#'
	}
	if symb == '#' && occupations >= 4 {
		return 'L'
	}
	return symb
}
func evaluateSeatPart2(x, y, symb int, grid [][]int) int{
	if symb == '.' {
		return symb
	}
	occupations := 0

	// Yeah, this is as good as it gets :(
	for i := 1; x-i >=0 && y-i >= 0; i++{
		if grid[x-i][y-i] == '#'{
			occupations++
			break
		}
		if grid[x-i][y-i] == 'L'{
			break
		}
	}
	for i := 1; x-i >=0; i++{
		if grid[x-i][y] == '#' {
			occupations++
			break
		}
		if grid[x-i][y] == 'L' {
			break
		}
	}
	for i := 1; y-i >= 0; i++{
		if grid[x][y-i] == '#' {
			occupations++
			break
		}
		if grid[x][y-i] == 'L' {
			break
		}
	}
	for i := 1; valid(x-i,y+i); i++{
		if grid[x-i][y+i] == '#'{
			occupations++
			break
		}
		if grid[x-i][y+i] == 'L'{
			break
		}
	}
	for i := 1; valid(x,y+i); i++{
		if grid[x][y+i] == '#' {
			occupations++
			break
		}
		if grid[x][y+i] == 'L' {
			break
		}
	}
	for i := 1; valid(x+i,y-i); i++{
		if grid[x+i][y-i] == '#'{
			occupations++
			break
		}
		if grid[x+i][y-i] == 'L'{
			break
		}
	}
	for i := 1; valid(x+i,y); i++{
		if grid[x+i][y] == '#'{
			occupations++
			break
		}
		if grid[x+i][y] == 'L'{
			break
		}
	}
	for i := 1; valid(x+i,y+i); i++{
		if grid[x+i][y+i] == '#'{
			occupations++
			break
		}
		if grid[x+i][y+i] == 'L'{
			break
		}
	}

	if symb == 'L' && occupations == 0 {
		return '#'
	}
	if symb == '#' && occupations >= 5 {
		return 'L'
	}
	return symb
}