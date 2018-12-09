package main

import (
	"fmt"
	"github.com/calvcoll/advent-of-code-2018/utils"
	"log"
	"strings"
)

func main() {
	elements := setup()
	challenge1(elements[0])
	challenge2(elements[0])
}

func setup() []string {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return lines
}

func reactPolymer(elements string) (polymer string) {
	for {
		found := false
		for index, _ := range elements {
			if index+1 < len(elements) {
				char1 := string(elements[index])
				char2 := string(elements[index+1])
				if char1 != char2 && (char1 == strings.ToUpper(char2) || char1 == strings.ToLower(char2)) {
					found = true
					elements = strings.Replace(elements, char1+char2, "", 1)
				}
			}
		}
		if !found {
			break
		}
	}
	return elements
}

func challenge1(elements string) {
	elements = reactPolymer(elements)
	fmt.Println("The full length of the polymer is:", len(elements))
}

func challenge2(elements string) {
	biggestUnit := "Aa"
	shortestPolymer := elements
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, unit := range alphabet {
		tempPolymer := elements
		tempPolymer = strings.Replace(tempPolymer, strings.ToLower(unit), "", -1)
		tempPolymer = strings.Replace(tempPolymer, strings.ToUpper(unit), "", -1)
		tempPolymer = reactPolymer(tempPolymer)
		if len(tempPolymer) < len(shortestPolymer) {
			biggestUnit = unit
			shortestPolymer = tempPolymer
		}
	}
	fmt.Printf("The most troublesome polymer is: %s and the full polymer is: %d\n", biggestUnit, len(shortestPolymer))
}
