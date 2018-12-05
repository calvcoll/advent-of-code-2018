package main

import (
	"github.com/calvcoll/advent-of-code/utils"
	"log"
	"strconv"
)

var frequencies = make(map[int]int)
var found = false

func getInput() []int {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	var nums []int
	for _, line := range lines {
		integer, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, integer)
	}
	return nums
}

func challenge1() {
	inputs := getInput()
	sum := 0
	for _, num := range inputs {
		sum += num
	}
	log.Printf("The total is %d", sum)
}

func logFrequency(frequency int) {
	count := frequencies[frequency] + 1
	frequencies[frequency] = count
	if count > 1 {
		found = true
	}
}

func challenge2() {
	inputs := getInput()
	sum := 0
	for ok := false; !ok; ok = found {
		for _, num := range inputs {
			sum += num
			logFrequency(sum)
			if found {
				break
			}
		}
	}
	log.Printf("The frequency first repeated is %d", sum)
}

func main() {
	challenge1()
	challenge2()
}
