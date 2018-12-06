package main

import (
	"fmt"
	"github.com/calvcoll/advent-of-code-2018/utils"
	"log"
	"strings"
)

var locations = make(map[Position]Location)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Location struct {
	Claims []Claim `json:"claims"`
}

func (location Location) getClaimCountOnLocation() (claimCount int) {
	return len(location.Claims)
}

type Claim struct {
	Id     int `json:"id"`
	PosX   int `json:"pos_x"`
	PosY   int `json:"pos_y"`
	XClaim int `json:"x_claim"`
	YClaim int `json:"y_claim"`
}

func ClaimFromArray(array []string) Claim {
	return Claim{
		Id:     utils.LogErrorAtoI(array[0]),
		PosX:   utils.LogErrorAtoI(array[1]),
		PosY:   utils.LogErrorAtoI(array[2]),
		XClaim: utils.LogErrorAtoI(array[3]),
		YClaim: utils.LogErrorAtoI(array[4]),
	}
}

func cleanupLine(str string) (cleanString string) {
	// #1333 @ 69,240: 11x28 -> 1333 69 240 11 28
	replace := strings.Replace(str, "#", "", -1)
	replace = strings.Replace(replace, "@ ", "", -1)
	replace = strings.Replace(replace, ":", "", -1)
	replace = strings.Replace(replace, "x", " ", -1)
	replace = strings.Replace(replace, ",", " ", -1)
	return replace
}

func (claim Claim) logClaim() {
	for x := claim.PosX; x < claim.PosX+claim.XClaim; x++ {
		for y := claim.PosY; y < claim.PosY+claim.YClaim; y++ {
			position := Position{X: x, Y: y}
			location := locations[position]

			location.Claims = append(location.Claims, claim)
			locations[position] = location
		}
	}
}

func setup() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	for _, line := range lines {
		line = cleanupLine(line)
		claim := ClaimFromArray(strings.Split(line, " "))
		claim.logClaim()
	}
}

func challenge1() {
	conflicts := 0
	for _, loc := range locations {
		if len(loc.Claims) > 1 {
			conflicts++
		}
	}
	fmt.Printf("There are %d conflicts. \n", conflicts)
}

func challenge2() {
	var noConflictClaim Claim
	for _, loc := range locations {
		if len(loc.Claims) == 1 {
			conflict := false
			for _, claim := range loc.Claims {
				for x := claim.PosX; (x < claim.PosX+claim.XClaim) && !conflict; x++ {
					for y := claim.PosY; y < claim.PosY+claim.YClaim && !conflict; y++ {
						position := Position{X: x, Y: y}
						location := locations[position]

						conflict = len(location.Claims) > 1
					}
				}
				if !conflict {
					noConflictClaim = claim
					break
				}
			}
		}
	}
	fmt.Printf("The claim without conflicts is %s.\n", utils.ToJsonString(noConflictClaim))
}

func main() {
	setup()
	challenge1()
	challenge2()
}
