package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("Day 14/input")
	file := bufio.NewScanner(f)
	hash := hashmap.New()
	t := time.Now()
	//part1(file, hash)
	part2(file, hash)
	fmt.Println("Time taken:", time.Now().Sub(t))

}

func part1 (file *bufio.Scanner, hash *hashmap.Map){
	mask := ""
	for file.Scan() {
		line := file.Text()
		tokens := strings.Split(line, " = ")
		if tokens[0] == "mask" {
			mask = tokens[1]
			continue
		}
		startIdx := strings.IndexByte(tokens[0], '[') + 1
		endIdx := strings.IndexByte(tokens[0], ']')

		arrIndex := tokens[0][startIdx:endIdx] // extract the value from [value]
		value, _ := strconv.Atoi(tokens[1])
		hash.Put(arrIndex, bitmask(value, mask))
	}

	var total int64 = 0
	for _, val := range hash.Values() {
		num := val.(int64)
		total += num
	}
	fmt.Println("Part 1 answer:", total)
}
func part2(file *bufio.Scanner, hash *hashmap.Map) {
	mask := ""
	for file.Scan() {
		line := file.Text()
		tokens := strings.Split(line, " = ")
		if tokens[0] == "mask" {
			mask = tokens[1]
			continue
		}
		startIdx := strings.IndexByte(tokens[0], '[') + 1
		endIdx := strings.IndexByte(tokens[0], ']')

		arrIndex, err := strconv.Atoi(tokens[0][startIdx:endIdx]) // extract the value from [value]
		if err != nil {
			panic(err)
		}
		value, _ := strconv.Atoi(tokens[1])
		indexMask := bitmaskPart2(arrIndex, mask)
		recursiveAdd(0, indexMask, hash, &value)
	}

	total := 0
	for _, val := range hash.Values() {
		num := val.(int)
		total += num
	}
	fmt.Println("Part 2 answer:", total)
}
func recursiveAdd(index int, indexmask string, hash *hashmap.Map, value *int) {
	if index >= len(indexmask) {
		// Need to convert and add to hashmap
		result, err := strconv.ParseInt(indexmask, 2, 64)
		if err != nil {
			panic(err)
		}
		hash.Put(result, *value)
		return
	}

	if indexmask[index] != 'X' {
		recursiveAdd(index+1, indexmask, hash, value)
	} else {

		tmp1 := []rune(indexmask)
		tmp1[index] = '0'
		indexmask = string(tmp1)
		recursiveAdd(index+1, indexmask, hash, value)

		tmp2 := []rune(indexmask)
		tmp2[index] = '1'
		indexmask = string(tmp2)
		recursiveAdd(index+1, indexmask, hash, value)
	}
}

func bitmaskPart2(value int, mask string) string {
	binary := strconv.FormatInt(int64(value), 2)
	length := len(binary)
	Buffer := bytes.Buffer{}
	workingMask := mask[len(mask) - length:]
	unusedMask := mask[:len(mask) - length]
	for i := 0; i < len(unusedMask); i++{
		Buffer.WriteByte(unusedMask[i])
	}
	for i := 0; i < length; i++  {
		if workingMask[i] == '0'{
			Buffer.WriteByte(binary[i])
		} else {
			Buffer.WriteByte(workingMask[i])
		}
	}
	return Buffer.String()
}

func bitmask (value int, mask string) int64 {
	 binary := strconv.FormatInt(int64(value), 2)
	 length := len(binary)
	 Buffer := bytes.Buffer{}
	 workingMask := mask[len(mask) - length:]
	 unusedMask := mask[:len(mask) - length]

	 for i := 0; i < len(unusedMask); i++{
		 if unusedMask[i] == '1'{
			 Buffer.WriteByte('1')
		 } else {
			 Buffer.WriteByte('0')
		 }
	 }

	 for i := 0; i < length; i++  {
		if workingMask[i] == 'X'{
			Buffer.WriteByte(binary[i])
		} else {
			Buffer.WriteByte(workingMask[i])
		}
	 }

	 result, err := strconv.ParseInt(Buffer.String(), 2, 64)
	 if err != nil {
	 	panic(err)
	 }
	 return result
 }

