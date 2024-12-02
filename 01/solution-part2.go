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
		fmt.Println(err)
		return
	}
	defer file.Close()

	// lists of ints
	leftList, rightList := []int{}, []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		result := strings.Split(line, "   ")
		// fmt.Println(result)
		var left string = result[0]
		var right string = result[1]
		leftInt, err := strconv.Atoi(left)

		if err != nil {
			fmt.Println(err)
			continue
		}

		rightInt, err := strconv.Atoi(right)

		if err != nil {
			fmt.Println(err)
			continue
		}

		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//fmt.Println(leftList)
	//fmt.Println(rightList)

	var totalScore int = 0

	for len(leftList) > 0 {
		score, newLeftList := simillarityScore(leftList, rightList)

		leftList = newLeftList

		totalScore += score

		//fmt.Println("newLeftList Score:", newLeftList, score)
	}

	elapsed := time.Since(start)

	fmt.Println("totalScore:", totalScore)
	fmt.Println("Took:", elapsed)
}

func simillarityScore(left []int, right []int) (int, []int) {

	leftNum := left[0]

	var times int = 0

	for _, num := range right {
		if leftNum == num {
			times += 1
		}
	}

	var newLeftList []int = removeElement(left, 0)

	return times * leftNum, newLeftList
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}
