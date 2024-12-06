package main

import (
	"advent-of-code-2024/lib"
	"fmt"
	"strconv"
	"strings"
)

const TestString string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

/*
	Part 1 Notes

	Parse the input and break it into rules and updates
	The rules specify that the first number has to appear before the second

	We need to find out which updates are valid according to the rules
	For the valid updates, we find the middle page number and add it to a total

	There's a handful of ways to go about this:
	1. Write a function that can determine if a rule is valid for
	a given update (does x appear before y if both appear)
	Iterate over updates, and for each update iterate over rules.

	2. Assuming there are no contradicting rules, we can make a list and sort, iterating over
	all of the rules until we get a sorted list.
	Then we should be able to see if the update is a subset of the sorted list.

	1 seems easier.
*/

type Rule struct {
	Lesser  int
	Greater int
}

type Update []int

func solvePart1(input string) int {
	splitString := strings.Split(input, "\n\n")
	rules := parseRules(splitString[0])
	updates := parseUpdates(splitString[1])

	total := 0

	for _, update := range updates {
		if update.satisfiesRules(rules) {
			total += update.middlePageNumber()
		}
	}

	return total
}

func parseRules(rulesString string) []Rule {
	lines := strings.Split(rulesString, "\n")
	rules := []Rule{}
	for _, line := range lines {
		split := strings.Split(line, "|")
		lesser, _ := strconv.Atoi(split[0])
		greater, _ := strconv.Atoi(split[1])
		rules = append(rules, Rule{Lesser: lesser, Greater: greater})
	}
	return rules
}

func parseUpdates(updatesString string) []Update {
	lines := strings.Split(updatesString, "\n")
	updates := []Update{}
	for _, line := range lines {
		split := strings.Split(line, ",")
		update := []int{}
		for _, num := range split {
			num, _ := strconv.Atoi(num)
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	return updates
}

func (u Update) satisfiesRule(rule Rule) bool {
	sawGreater := false
	for _, num := range u {
		if num == rule.Greater {
			sawGreater = true
		}
		if num == rule.Lesser {
			if sawGreater {
				return false
			} else {
				return true
			}
		}
	}
	return true
}

func (u Update) satisfiesRules(rules []Rule) bool {
	for _, rule := range rules {
		if !u.satisfiesRule(rule) {
			return false
		}
	}
	return true
}

func (u Update) middlePageNumber() int {
	return u[len(u)/2]
}

/*
	Part 2 Notes

	We need to find the updates that are invalid
	Then we write a method that takes an rule that the update doesn't satisfy, and changes
	the update to satisfy the rule.
	We iterate over all of the rules for that update until it is valid.
	Then we add the middle number of that update to the total.
*/

func solvePart2(input string) int {
	splitString := strings.Split(input, "\n\n")
	rules := parseRules(splitString[0])
	updates := parseUpdates(splitString[1])

	total := 0

	for _, update := range updates {
		if !update.satisfiesRules(rules) {
			for !update.satisfiesRules(rules) {
				for _, rule := range rules {
					if !update.satisfiesRule(rule) {
						update = update.fixFromRule(rule)
					}
				}
			}
			total += update.middlePageNumber()
		}
	}

	return total
}

func (u Update) fixFromRule(rule Rule) Update {
	newUpdate := []int{}
	for _, num := range u {
		if num == rule.Lesser {
			newUpdate = append(newUpdate, rule.Lesser)
			newUpdate = append(newUpdate, rule.Greater)
		} else if num != rule.Greater {
			newUpdate = append(newUpdate, num)
		}
	}
	return newUpdate
}

func main() {
	lib.AssertEqual(143, solvePart1(TestString))
	lib.AssertEqual(123, solvePart2(TestString))

	dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
