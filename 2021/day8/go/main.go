package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/mreed19/advent-of-code/2021/day8/go/data"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("data/input.txt"))
	fmt.Println("Part 2 Solution:", part2("data/input.txt"))
}

func part1(filename string) int {
	_, outputs := data.GetData(filename)

	count := 0
	for _, output := range outputs {
		count += get1478Count(output)
	}

	return count
}

func part2(filename string) int {
	nums, outputs := data.GetData(filename)

	numbers := &Numbers{}

	sum := 0
	for index := range nums {
		numbers.get1478(nums[index])
		numbers.get023569(nums[index])
		numbersMap := numbers.toMap()

		numString := ""
		for _, val := range outputs[index] {
			valSorted := strings.Split(val, "")
			sort.Strings(valSorted)
			numString += numbersMap[strings.Join(valSorted, "")]
		}
		numVal, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal(err)
		}
		sum += numVal
	}

	return sum
}

func get1478Count(arr []string) int {
	count := 0
	for _, val := range arr {
		length := len(val)
		if length == 2 ||
			length == 4 ||
			length == 3 ||
			length == 7 {
			count++
		}
	}
	return count
}

type Numbers struct {
	zero  []string
	one   []string
	two   []string
	three []string
	four  []string
	five  []string
	six   []string
	seven []string
	eight []string
	nine  []string
}

func (n *Numbers) toMap() map[string]string {
	numbersMap := make(map[string]string)
	numbersMap[strings.Join(n.zero, "")] = "0"
	numbersMap[strings.Join(n.one, "")] = "1"
	numbersMap[strings.Join(n.two, "")] = "2"
	numbersMap[strings.Join(n.three, "")] = "3"
	numbersMap[strings.Join(n.four, "")] = "4"
	numbersMap[strings.Join(n.five, "")] = "5"
	numbersMap[strings.Join(n.six, "")] = "6"
	numbersMap[strings.Join(n.seven, "")] = "7"
	numbersMap[strings.Join(n.eight, "")] = "8"
	numbersMap[strings.Join(n.nine, "")] = "9"
	return numbersMap
}

func (n *Numbers) get1478(arr []string) {
	for _, val := range arr {
		valArr := strings.Split(val, "")
		sort.Strings(valArr)
		switch len(val) {
		case 2:
			n.one = valArr
		case 4:
			n.four = valArr
		case 3:
			n.seven = valArr
		case 7:
			n.eight = valArr
		}
	}
}

func (n *Numbers) get023569(arr []string) {
	for _, val := range arr {
		valArr := strings.Split(val, "")
		sort.Strings(valArr)
		switch len(val) {
		case 6:
			if match, _ := containsSubstrings(val, n.seven...); !match {
				n.six = valArr
			} else if match, _ := containsSubstrings(val, n.four...); match {
				n.nine = valArr
			} else {
				n.zero = valArr
			}
		case 5:
			if match, _ := containsSubstrings(val, n.one...); match {
				n.three = valArr
			} else if _, matchCount := containsSubstrings(val, n.four...); matchCount == 3 {
				n.five = valArr
			} else {
				n.two = valArr
			}
		}
	}
}

func containsSubstrings(str string, subs ...string) (bool, int) {
	isCompleteMatch := true
	matchCount := 0
	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matchCount++
		} else {
			isCompleteMatch = false
		}
	}
	return isCompleteMatch, matchCount
}
