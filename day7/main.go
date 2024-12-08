package main

import (
	"advent-of-code-2024/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const TestString string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

/*
	Part 1 Notes

	Okay we need to figure out which of the equations can be solved by adding * and + between
	the numbers (order of operations not a factor here, everything is left to right).

	I think there might be a fancy way of figuring this out using factors and stuff, but that's
	not the simple way.

	Simple way seems to be a recursive solution, that passes the remaining array with
	the current value, spwans more functions, one to multiply and one to add the next number

	Set a total to 0
	Parse each of the lines for an equation with result and nums.
	For each equation:
		Pop the first number, set that to the result
		Pass the currentResult and remaining array to the recursive function
		If recursive function returns true, then add the result to the total

	The recursive function
		Returns false currentResult is larger than the intendedResult
		Returns false if the array is empty
		Pops the next num off the array
	  	Return the recursive function with (currentResult + nextNum, remainingArray) ||
				(currrentResult * num, remainingArray)

*/

type Equation struct {
	Result int
	Nums   []int
}

func readEquations(input string) []Equation {
	equations := []Equation{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result, _ := strconv.Atoi(parts[0])
		nums := strings.Split(parts[1], " ")
		numsInts := make([]int, len(nums))
		for i, num := range nums {
			numsInts[i], _ = strconv.Atoi(num)
		}
		equations = append(equations, Equation{Result: result, Nums: numsInts})
	}
	return equations
}

func solvableEquation(equation Equation) bool {
	newSlice := equation.Nums[1:]
	currentResult := equation.Nums[0]
	return solveEquation(equation.Result, currentResult, newSlice)
}

func solveEquation(desiredResult, currentResult int, nums []int) bool {
	if len(nums) == 0 {
		return desiredResult == currentResult
	}

	if desiredResult < currentResult {
		return false
	}

	newSlice := nums[1:]
	return solveEquation(desiredResult, currentResult+nums[0], newSlice) ||
		solveEquation(desiredResult, currentResult*nums[0], newSlice)
}

func solvePart1(input string) int {
	equations := readEquations(input)

	total := 0
	for _, equation := range equations {
		if solvableEquation(equation) {
			total += equation.Result
		}
	}

	return total
}

/*
	Part 2 Notes

	Okay now there is a new operation || that combines the numbers.
	I might need to optimize seeing as the base N value got a bit bigger here.
	But lets try it without optimization first.

	I need to write a function that takes two numbers and combines them.
	Then I use the same code, except instead of just adding and multiplying,
	I also combine.
*/

func combineNumbers(num1, num2 int) int {
	multipliedFirstNum := num1 * (int(math.Pow(10, float64(digitCount(num2)))))
	return multipliedFirstNum + num2
}

func digitCount(num int) int {
	s := fmt.Sprintf("%d", num)
	return len(s)
}

func solveEquationWithCombination(desiredResult, currentResult int, nums []int) bool {
	if len(nums) == 0 {
		return desiredResult == currentResult
	}

	if desiredResult < currentResult {
		return false
	}

	newSlice := nums[1:]
	return solveEquationWithCombination(desiredResult, currentResult+nums[0], newSlice) ||
		solveEquationWithCombination(desiredResult, currentResult*nums[0], newSlice) ||
		solveEquationWithCombination(desiredResult, combineNumbers(currentResult, nums[0]), newSlice)
}

func solvableEquationWithCombination(equation Equation) bool {
	newSlice := equation.Nums[1:]
	currentResult := equation.Nums[0]
	return solveEquationWithCombination(equation.Result, currentResult, newSlice)
}

func solvePart2(input string) int {
	equations := readEquations(input)

	total := 0
	for _, equation := range equations {
		if solvableEquationWithCombination(equation) {
			total += equation.Result
		}
	}

	return total
}

func main() {
	lib.AssertEqual(3749, solvePart1(TestString))
	lib.AssertEqual(11387, solvePart2(TestString))

	// dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
