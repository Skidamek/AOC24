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

	orderingRules := make(map[int][]int)
	printOrders := make([][]int, 0)

	firstPart := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			firstPart = false
			continue
		}

		if firstPart {
			split := strings.Split(line, "|")
			key, _ := strconv.Atoi(split[0])
			value, _ := strconv.Atoi(split[1])
			orderingRules[key] = append(orderingRules[key], value)
		} else {
			split := strings.Split(line, ",")
			intList, _ := strSliceToIntSlice(split)
			printOrders = append(printOrders, intList)
		}
	}

	mids := make([]int, 0)

	for _, orderList := range printOrders {

		if isSorted(orderList, orderingRules) {
			continue
		}

		sortedOrder := append([]int{}, orderList...)

		// sort it till its doneee
		for !isSorted(sortedOrder, orderingRules) {
			// sort in reverse
			for numId := len(sortedOrder) - 1; numId >= 0; numId-- {
				num := sortedOrder[numId]

				// check if num exists as a key in ordering rules
				// if so get the values and check if any of are before the key number
				// swap and repeat

				if values, exists := orderingRules[num]; exists {
					for _, value := range values {
						valueID := -1
						for i, v := range sortedOrder {
							if v == value {
								valueID = i
								break
							}
						}

						if valueID == -1 {
							continue
						}

						if numId > valueID {
							// swap them!
							sortedOrder[numId], sortedOrder[valueID] = sortedOrder[valueID], sortedOrder[numId]
						}
					}
				}
			}
		}

		mids = append(mids, sortedOrder[len(sortedOrder)/2])
	}

	result := 0
	for _, num := range mids {
		result += num
	}

	elapsed := time.Since(start)

	fmt.Println("Result:", result)
	fmt.Println("Took:", elapsed)
}

func isSorted(orderList []int, orderingRules map[int][]int) bool {
	valid := true
	for numId, num := range orderList {
		// check if num exists as a key in ordering rules
		// if so get the values and check if any of are before the key number
		// if so continue next list because this one is invalid

		if values, exists := orderingRules[num]; exists {
			for _, value := range values {
				// Find the index of `value` in `orderList`
				valueID := -1
				for i, v := range orderList {
					if v == value {
						valueID = i
						break
					}
				}

				if valueID == -1 {
					continue
				}

				if numId > valueID {
					valid = false
					break
				}
			}
		}
	}

	return valid
}

func strSliceToIntSlice(strs []string) ([]int, error) {
	intList := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err // Return error if conversion fails
		}
		intList[i] = num
	}
	return intList, nil
}
