package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	val := make([]int, 200)

	//val := []int { 1721,979, 366, 299,	675, 1456}

	f, err := os.Open("inputTask1")
	if err != nil {
		panic(err)
	}
	file := bufio.NewScanner(f)

	index := 0
	for file.Scan() {
		num, err := strconv.Atoi(file.Text())
		if err != nil {
			panic(err)
		}
		val[index] = num
		index++
	}


	fmt.Println("Calculating part 1")
	t := time.Now()
	for i := 0; i < len(val) - 1; i++ {
		for j := i; j < len(val); j++ {
			if val[i] + val[j] == 2020 {
				fmt.Println("Answer:", val[i] * val[j], "numbers are:", val[i], val[j])
				i = len(val)
				j = len(val)
			}
		}
	}
	fmt.Println("Time Taken:", time.Now().Sub(t))


	fmt.Println("Calculating part 2")
	// 3SUM method
	t1 := time.Now()
	sort.Ints(val)
	for i := 0; i < len(val) -1; i++ {
		l := i + 1
		r := len(val) - 1
		x := val[i]

		for l < r {
			if x + val[l] + val[r] == 2020 {
				fmt.Println("Answer:", val[i] * val[l] * val[r], "numbers are:", val[i], val[l], val[r])
				i = len(val)
				break
			} else if x + val[l] + val[r] < 2020 {
				l++
			} else {
				r--
			}

		}
	}

	fmt.Println("Time Taken:", time.Now().Sub(t1))
}