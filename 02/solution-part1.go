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

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	safeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if isSequenceSafe(parts) {
			//fmt.Println("Result is safe:", parts)
			safeCount++
		} else {
			//fmt.Println("Result is not safe:", parts)
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Total safe sequences:", safeCount)
	fmt.Println("Took:", elapsed)
}

func isSequenceSafe(parts []string) bool {
	if len(parts) < 2 {
		return false
	}

	numbers := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Invalid number:", part)
			return false
		}
		numbers[i] = num
	}

	isIncreasing := numbers[1] > numbers[0]
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		absDiff := abs(diff)

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}

	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
