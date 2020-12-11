package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Open("Day 11/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)
	allNumbers := [][]int{}

	for file.Scan() {
		num := file.Text()
		line := []int{}
		for _, i := range num {
			line = append(line, int(i))
		}
		allNumbers = append(allNumbers, line)
	}
	t := time.Now()
	for {
		newGrid := [][]int{}
		for x, arr := range allNumbers{
			newLine := []int{}
			for y, symb := range arr {
				//symb := evaluateSeatPart1(x,y, symb, allNumbers)
				symb := evaluateSeatPart2(x,y, symb, allNumbers)
				newLine = append(newLine, symb)
			}
			newGrid = append(newGrid, newLine)
		}
		if checkIfSame(allNumbers, newGrid) {
			fmt.Println("Same grid detected")
			break
		}
		allNumbers = newGrid
	}

	fmt.Println("Answer:", checkForOccupiedSeats(allNumbers))
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
	xLength := len(grid)
	yLength := len(grid[0])

	if x-1 >= 0 && y-1 >= 0 && grid[x-1][y-1] == '#' {
		occupations++
	}
	if x-1 >= 0 && grid[x-1][y] == '#' {
		occupations++
	}
	if y-1 >= 0 && grid[x][y-1] == '#' {
		occupations++
	}
	if x-1 >= 0 && y+1 < yLength &&  grid[x-1][y+1] == '#' {
		occupations++
	}
	if y+1 < yLength && grid[x][y+1] == '#' {
		occupations++
	}
	if x+1 < xLength && y-1 >= 0 && grid[x+1][y-1] == '#' {
		occupations++
	}
	if x+1 < xLength && grid[x+1][y] == '#' {
		occupations++
	}
	if x+1 < xLength && y+1 < yLength && grid[x+1][y+1] == '#' {
		occupations++
	}

	if symb == 'L' && occupations == 0{
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
	xLength := len(grid)
	yLength := len(grid[0])

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
	for i := 1; x-i >=0 && y+i < yLength; i++{
		if grid[x-i][y+i] == '#'{
			occupations++
			break
		}
		if grid[x-i][y+i] == 'L'{
			break
		}
	}
	for i := 1; y+i < yLength; i++{
		if grid[x][y+i] == '#' {
			occupations++
			break
		}
		if grid[x][y+i] == 'L' {
			break
		}
	}
	for i := 1; x+i < xLength && y-i >= 0; i++{
		if grid[x+i][y-i] == '#'{
			occupations++
			break
		}
		if grid[x+i][y-i] == 'L'{
			break
		}
	}
	for i := 1; x+i < xLength; i++{
		if grid[x+i][y] == '#'{
			occupations++
			break
		}
		if grid[x+i][y] == 'L'{
			break
		}
	}
	for i := 1;x+i < xLength && y+i < yLength; i++{
		if grid[x+i][y+i] == '#'{
			occupations++
			break
		}
		if grid[x+i][y+i] == 'L'{
			break
		}
	}

	if symb == 'L' && occupations == 0{
		return '#'
	}
	if symb == '#' && occupations >= 5 {
		return 'L'
	}
	return symb
}