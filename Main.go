//GoCalc, app only works with 2 positive integers at this time.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a int
var b int
var answer int
var operator string

func main() {
	fmt.Println("Welcome to GoCalc")
	x := 0
	for x != -1 {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter -1 to quit")
		fmt.Print("Start by entering a problem: ")

		//scanner open
		scanner.Scan()
		equation := scanner.Text()
		if equation == "-1" {
			break
		}

		//strip whitespace
		fmtequ := stripWhiteSpace(equation)

		//send values to array
		var eqArray []string
		eqArray = toArray(fmtequ)

		//get values and operator from array
		getOperator(eqArray)
		getVariables(eqArray)

		//dictate the equation type
		solveEquation(operator, a, b)

		//print readable answer to equation
		fmt.Println(fmtequ, "= ", answer)
	}

}
func stripWhiteSpace(input string) string {
	input1 := strings.ReplaceAll(input, " ", "")
	return input1
}
func solveEquation(operator string, a int, b int) {

	switch operator {
	case "+":
		answer = add(a, b)
	case "-":
		answer = sub(a, b)
	case "*":
		answer = mult(a, b)
	case "/":
		answer = divi(a, b)
	}

}

// read in equation array and find the operator
func getOperator(eqArray []string) {
	for _, value := range eqArray {
		if string(value) == "+" || string(value) == "-" || string(value) == "/" || string(value) == "*" {
			operator = string(value)
		}
	}
}

// read in Equation array and assign values to variables A & B
func getVariables(eqArray []string) {
	var value1 string
	var value2 string
	i := 0
	for _, value := range eqArray {
		if _, err := strconv.Atoi(value); err == nil {
			value1 += string(value)
			eqArray[i] = ""
			i++
		} else {
			eqArray[i] = ""
			for _, value := range eqArray {
				if _, err := strconv.Atoi(value); err == nil {
					value2 += string(value)
				}
			}
			break
		}

	}

	a = toInt(value1)
	b = toInt(value2)
}
func toInt(value string) int {
	intvalue, _ := strconv.Atoi(value)
	return intvalue
}

// Takes string in and returns array of strings
func toArray(eq string) []string {
	var eqArray []string
	for _, value := range eq {
		eqArray = append(eqArray, string(value))
	}

	return eqArray
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func sub(a int, b int) int {
	dif := a - b
	return dif
}

func mult(a int, b int) int {
	prod := a * b
	return prod
}

func divi(a int, b int) int {
	quot := a / b
	return quot
}
