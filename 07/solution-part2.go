package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()

	equations := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()

		if len(text) == 0 {
			fmt.Println("input is incorrect")
			continue
		}

		parts := strings.Split(text, ": ")
		if len(parts) != 2 {
			fmt.Println("input is incorrect")
			continue
		}

		equations = append(equations, parts)
	}

	result := 0

	for _, equationParts := range equations {
		equation := make([]int, 0)
		for _, part := range strings.Split(equationParts[1], " ") {
			number := 0
			fmt.Sscanf(part, "%d", &number)
			equation = append(equation, number)
		}

		correctSum := 0
		fmt.Sscanf(equationParts[0], "%d", &correctSum)

		equationLength := len(equation)
		if equationLength < 2 || correctSum == 0 {
			fmt.Println("input is incorrect")
			continue
		}

		totalCombinations := intPow(3, equationLength-1)

		for combination := 0; combination < totalCombinations; combination++ {
			currentValue := equation[0]
			var operation string

			for i := 0; i < equationLength-1; i++ {
				operationIndex := (combination / intPow(3, i)) % 3

				switch operationIndex {
				case 0:
					operation = "+"
				case 1:
					operation = "*"
				case 2:
					operation = "||"
				}

				switch operation {
				case "+":
					currentValue += equation[i+1]
				case "*":
					currentValue *= equation[i+1]
				case "||":
					currentValue, _ = strconv.Atoi(fmt.Sprintf("%d%d", currentValue, equation[i+1]))
				}
			}

			if currentValue == correctSum {
				fmt.Println("Equation:", equationParts[1], "=", correctSum)
				result += correctSum
				break
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println("True equasions:", result)
	fmt.Println("Took:", elapsed)
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
