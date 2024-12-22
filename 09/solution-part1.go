package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()

	diskMap := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			diskMap += text
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	output := make([]int, 0)
	for id, char := range diskMap {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Printf("Invalid character '%c' in input\n", char)
			continue
		}

		emptySpace := id%2 == 0
		output = append(output, repeatChar(id, num, emptySpace)...)
	}

	moveValuesToSpaces(&output)

	newOutput := make([]int, 0)
	for _, num := range output {
		if num != -1 {
			newOutput = append(newOutput, num)
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Checksum:", checksum(newOutput))
	fmt.Println("Took:", elapsed)
}

func checksum(output []int) int64 {
	var checksum int64 = 0
	for id, num := range output {
		checksum += int64(num) * int64(id)
	}
	return checksum
}

func repeatChar(id int, count int, emptySpace bool) []int {
	if count <= 0 {
		return []int{}
	}

	result := make([]int, count)
	for i := 0; i < count; i++ {
		if emptySpace {
			result[i] = id / 2
		} else {
			result[i] = -1
		}
	}

	return result
}

func moveValuesToSpaces(output *[]int) {
	for i := 0; i < len(*output); i++ {
		// Skip actual data
		if (*output)[i] != -1 {
			continue
		}

		// Find the first unfilled value that can be moved to this empty space
		for j := len(*output) - 1; j >= 0; j-- {
			if (*output)[j] != -1 {
				(*output)[i] = (*output)[j]
				(*output)[j] = -1
				break
			}
		}
	}
}
