package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()

	lines := make([][]rune, 0) // y x

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		line := make([]rune, 0)

		for _, c := range text {
			line = append(line, c)
		}

		lines = append(lines, line)
	}

	// find ^ in lines and get x y
	startingPoints := make(map[string]struct{})

	for y, line := range lines {
		for x, c := range line {
			if c == '0' {
				fmt.Printf("Starting point found at %d %d\n", x, y)
				startingPoints[fmt.Sprintf("%d %d", x, y)] = struct{}{}
			}
		}
	}

	result := 0

	// for each starting point find the path
	for startingPoint := range startingPoints {
		score := 0

		// check x+ x- y+ y- for numToFind starting from startingPoint which is num 0, increment score if found 9 after 1-8
		visited := make(map[int][]string)
		visited[0] = []string{startingPoint}
		numToFind := 1

		for _, coords := range visited[numToFind-1] {
			// find the starting point
			x, y := 0, 0
			fmt.Sscanf(startingPoint, "%d %d", &x, &y)

			// check x+ x- y+ y- for numToFind starting from startingPoint which is num 0, increment score if found 9 after 1-8

			for _, dx := range []int{-1, 0, 1} {
				for _, dy := range []int{-1, 0, 1} {
					if dx == 0 && dy == 0 {
						continue
					}

					newX, newY := x+dx, y+dy

				}
			}

			if numToFind == 9 {
				break
			}
		}

		result += score
	}

	elapsed := time.Since(start)

	fmt.Println("Result:", result)
	fmt.Println("Took:", elapsed)
}
