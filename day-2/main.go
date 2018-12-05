package main

import (
	"fmt"
	"github.com/calvcoll/advent-of-code-2018/utils"
	"log"
	"strings"
)

var twoCharWords []string
var threeCharWords []string

type Char rune

func getInput() []string {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return lines
}

func checkCounts(counts map[Char]int, id string) {
	foundTwo := false
	foundThree := false
	for _, count := range counts {
		if !foundTwo && count == 2 {
			foundTwo = true
			twoCharWords = append(twoCharWords, id)
		}
		if !foundThree && count == 3 {
			foundThree = true
			threeCharWords = append(threeCharWords, id)
		}
	}
}

func setup() {
	inputs := getInput()
	for _, id := range inputs {
		counts := make(map[Char]int)
		for _, char := range id {
			counts[Char(char)] += 1
		}
		checkCounts(counts, id)
	}
}

func challenge1() {
	fmt.Printf("The total two count is: %d\nThe total three count is: %d\nThe answer is: %d\n", len(twoCharWords), len(threeCharWords), len(twoCharWords)*len(threeCharWords))
}

func challenge2() {
	found := false
	matchWord1 := ""
	matchWord2 := ""
	for _, word := range twoCharWords {
		for _, word2 := range twoCharWords {
			if word != word2 {
				// dont want the same word of course
				found = isMatch(word, word2)
			}
			if found {
				matchWord1 = word
				matchWord2 = word2
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		for _, word := range threeCharWords {
			for _, word2 := range threeCharWords {
				if word != word2 {
					// dont want the same word of course
					found = isMatch(word, word2)
				}
				if found {
					matchWord1 = word
					matchWord2 = word2
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			for _, word := range twoCharWords {
				for _, word2 := range threeCharWords {
					if word != word2 {
						// dont want the same word of course
						found = isMatch(word, word2)
					}
					if found {
						matchWord1 = word
						matchWord2 = word2
						break
					}
				}
				if found {
					break
				}
			}
		}
	}

	if found {
		fmt.Printf("Found words: %s %s\nMatching letters are: %s\n", matchWord1, matchWord2, findOneLetterDifference(matchWord1, matchWord2))
	} else {
		fmt.Println("Couldn't match words")
	}
}

func findOneLetterDifference(s string, s2 string) string {
	var index int
	for index = 0; index < len(s); index++ {
		if s[index] != s2[index] {
			break
		}
	}
	return strings.Replace(s, string(s[index]), "", -1)
}

func isMatch(s string, s2 string) bool {
	count := 0
	if len(s) != len(s2) {
		return false
	}
	for index := 0; index < len(s); index++ {
		if s[index] == s2[index] {
			count += 1
		}
	}
	if count == len(s)-1 {
		return true
	}
	return false
}

func main() {
	setup()
	challenge1()
	challenge2()
}
