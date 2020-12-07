package main

import (
	"bufio"
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	hash_count := hashmap.New()
	hash_map := hashmap.New()
	f, err := os.Open("Day 7/input")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(f)

	for file.Scan() {
		line := file.Text()
		values := strings.Split(line, " contain ")
		values[1] = strings.TrimRight(values[1], ".")
		values[1] = strings.ReplaceAll(values[1], "bags", "bag")
		values[0] = strings.ReplaceAll(values[0], "bags", "bag")

		// values [Bag, BagContent]
		hash_map.Put(values[0], values[1])
		hash_count.Put(values[0], false)

	}

	t := time.Now()
	part1(hash_map, hash_count)
	part2(hash_map)
	fmt.Println("Time taken:", time.Now().Sub(t))
}

func part2 (hash_map *hashmap.Map){
	total := 0
	tmpVal, _ := hash_map.Get("shiny gold bag")
	val := tmpVal.(string)
	content := strings.Split(val, ", ")

	for _, i := range content {
		pair := getValuePair(i)
		i, _ := strconv.Atoi(pair[0])
		total += i * (recursiveFind2(pair[1], hash_map) + 1)
	}

	fmt.Println("Total bags contained in shiny gold bag:",total)
}
func recursiveFind2(key string, hash_map *hashmap.Map) int{
	tmpVal, _ := hash_map.Get(key)
	val := tmpVal.(string)
	content := strings.Split(val, ", ")
	total := 0
	for _, i := range content {

		pair := getValuePair(i)
		if pair[0] != "no" {
			i, _ := strconv.Atoi(pair[0])
			total += i * (recursiveFind2(pair[1], hash_map) + 1)
		} else {
			return 0
		}
	}
	return total
}
func part1 (hash_map, hash_count *hashmap.Map) {
	for _, item := range hash_map.Keys() {
		tmp, _ := hash_count.Get(item)
		tmpVal, _ := hash_map.Get(item)
		isSet := tmp.(bool)
		val := tmpVal.(string)
		if !isSet && strings.Contains(val, "shiny gold bag"){
			hash_count.Put(item, true)
			newSearch, _ := item.(string)
			findRecursive(hash_map, hash_count, newSearch)
		}
	}

	total := 0
	for _, i := range hash_count.Keys() {
		found, _ := hash_count.Get(i)
		f, _ := found.(bool)
		if f {
			total++
		}
	}
	fmt.Println("Total bags containing shiny gold bag:",total)

}
func findRecursive(hash_map, hash_count *hashmap.Map, newKey string){

	for _, item := range hash_map.Keys() {
		tmp, _ := hash_count.Get(item)
		tmpVal, _ := hash_map.Get(item)
		isSet := tmp.(bool)
		val := tmpVal.(string)
		if !isSet && strings.Contains(val, newKey){
			hash_count.Put(item, true)
			newSearch, _ := item.(string)
			findRecursive(hash_map, hash_count, newSearch)
		}
	}

}

func getValuePair(contains string) []string{
	return strings.SplitN(contains, " ", 2)
}