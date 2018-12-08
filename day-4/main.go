package main

import (
	"fmt"
	"github.com/calvcoll/advent-of-code-2018/utils"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ByAlphabetically []string

func (a ByAlphabetically) Len() int {
	return len(a)
}

func (a ByAlphabetically) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAlphabetically) Less(i, j int) bool {
	return a[i] < a[j]
}

func getInput() []string {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	sort.Sort(ByAlphabetically(lines))
	return lines
}

type Guard struct {
	Id                 int         `json:"id"`
	TotalTimeAsleep    int         `json:"total_time_asleep"`
	MinutesAsleepCount map[int]int `json:"minutes_asleep_count"`
}

func InitGuard(id int) (guard Guard) {
	return Guard{
		Id:                 id,
		TotalTimeAsleep:    0,
		MinutesAsleepCount: map[int]int{},
	}
}

var regexId = `#(?P<id>\d+)`
var regexTime = `\[\d{4}-\d{2}-\d{2} \d{2}:(?P<mins>\d{2})\]`
var guards = map[int]Guard{}

func getMostSleptMinute(guard Guard) (minuteSleptMost int) {
	highestFrequency := 0
	for minute, freq := range guard.MinutesAsleepCount {
		if freq > highestFrequency {
			minuteSleptMost = minute
			highestFrequency = freq
		}
	}
	return minuteSleptMost
}

func challenge1() {
	wokeGuard := Guard{
		TotalTimeAsleep: 0,
	}
	for _, guard := range guards {
		if guard.TotalTimeAsleep > wokeGuard.TotalTimeAsleep {
			wokeGuard = guard
		}
	}
	freqMin := getMostSleptMinute(wokeGuard)
	fmt.Printf("Most woke guard is: %s\n", utils.ToJsonString(wokeGuard))
	fmt.Printf("Guard %d slept %d minutes, the most frequent was %d. For an answer of %d\n",
		wokeGuard.Id, wokeGuard.TotalTimeAsleep, freqMin, wokeGuard.Id*freqMin)
}

func challenge2() {
	highestFreq := 0
	mostSlept := 0
	var slepGuard Guard
	for _, guard := range guards {
		min := getMostSleptMinute(guard)
		if guard.MinutesAsleepCount[min] > highestFreq {
			mostSlept = min
			slepGuard = guard
			highestFreq = guard.MinutesAsleepCount[min]
		}
	}
	fmt.Printf("The guard was %d, asleep on %d for %d times, for an answer of: %d\n", slepGuard.Id, mostSlept, highestFreq, slepGuard.Id*mostSlept)
}

func main() {
	setup()
	challenge1()
	challenge2()
}

func setup() {
	var currentGuard Guard
	counter := -1
	timeFellAsleep := 0

	for _, line := range getInput() {
		lowerLine := strings.ToLower(line)

		if strings.Contains(lowerLine, "guard") {
			regex := regexp.MustCompile(regexId)
			id, _ := strconv.Atoi(regex.FindStringSubmatch(line)[1])

			if counter == -1 {
				currentGuard = InitGuard(id)
			}
			if counter > 0 {
				counter = -1
				guards[currentGuard.Id] = currentGuard
				if guards[id].Id == id {
					currentGuard = guards[id]
				} else {
					currentGuard = InitGuard(id)
				}
			}
		}
		if strings.Contains(lowerLine, "falls asleep") {
			regex := regexp.MustCompile(regexTime)
			mins, _ := strconv.Atoi(regex.FindStringSubmatch(line)[1])
			timeFellAsleep = mins
		}
		if strings.Contains(lowerLine, "wakes up") {
			regex := regexp.MustCompile(regexTime)
			mins, _ := strconv.Atoi(regex.FindStringSubmatch(line)[1])
			currentGuard.TotalTimeAsleep += mins - timeFellAsleep
			for i := timeFellAsleep; i <= mins; i++ {
				currentGuard.MinutesAsleepCount[i] += 1
			}
			timeFellAsleep = 0
		}
		counter += 1
	}
}
