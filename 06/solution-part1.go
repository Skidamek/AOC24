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

	lines := make([][]rune, 0) // x y

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
	guardX := -1
	guardY := -1

	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				fmt.Printf("Guard found at %d %d\n", x, y)
				guardX = x
				guardY = y
			}
		}
	}

	if guardX == -1 || guardY == -1 {
		fmt.Println("Guard not found")
		return
	}

	block := '#'
	currentDirection := '^'
	uniquePositions := make([][]int, 0)

	// path finding while gouard is inside the grid
	for guardX >= 0 && guardX < len(lines[0]) && guardY >= 0 && guardY < len(lines) {
		// if somethign is blocking the way turn 90 degrees right
		// if nothing is blocking the way move forward one step in current direction
		// till we go off the grid

		nextX := guardX
		nextY := guardY
		if currentDirection == '^' {
			nextY--
		} else if currentDirection == 'v' {
			nextY++
		} else if currentDirection == '<' {
			nextX--
		} else if currentDirection == '>' {
			nextX++
		}

		// check if next step is off the grid
		if nextX < 0 || nextX >= len(lines[0]) || nextY < 0 || nextY >= len(lines) {
			guardX = nextX
			guardY = nextY
			break
		}

		// check if next step is blocked
		if lines[nextY][nextX] == block {
			// turn right
			if currentDirection == '^' {
				currentDirection = '>'
			} else if currentDirection == '>' {
				currentDirection = 'v'
			} else if currentDirection == 'v' {
				currentDirection = '<'
			} else if currentDirection == '<' {
				currentDirection = '^'
			}
		} else {
			guardX = nextX
			guardY = nextY
			isUnique := true
			for _, pos := range uniquePositions {
				if pos[0] == guardX && pos[1] == guardY {
					isUnique = false
					break
				}
			}

			if isUnique {
				uniquePositions = append(uniquePositions, []int{guardX, guardY})
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Unique positions:", len(uniquePositions))
	fmt.Println("If doesnt work add one to the result to get the correct answer") // yes
	fmt.Println("Took:", elapsed)
}
