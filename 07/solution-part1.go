package main

import (
	"bufio"
	"fmt"
	"os"
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
			continue
		}

		parts := strings.Split(text, ": ")
		if len(parts) != 2 {
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

		// this is effective way of making power of 2 in go, 2 ^ equationLength is XOR instead of math power
		// 1 == 0000 0001
		// equationLength == 4
		// 1 << 4
		// 16 == 0001 0000
		// 2^4 == 16
		totalCombinations := 1 << (equationLength - 1)

		for combination := 0; combination < totalCombinations; combination++ {
			currentValue := equation[0]

			operation := combination
			for i := 0; i < equationLength-1; i++ {
				if operation%2 == 0 {
					currentValue *= equation[i+1]
				} else {
					currentValue += equation[i+1]
				}
				operation = operation / 2
			}

			if currentValue == correctSum {
				result += correctSum
				break
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println("True equasions:", result)
	fmt.Println("Took:", elapsed)
}
