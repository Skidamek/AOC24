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
	rootGuardX := -1
	rootGuardY := -1

	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				fmt.Printf("Guard found at %d %d\n", x, y)
				rootGuardX = x
				rootGuardY = y
			}
		}
	}

	if rootGuardX == -1 || rootGuardY == -1 {
		fmt.Println("Guard not found")
		return
	}

	block := '#'
	currentDirection := '^'
	guardX := rootGuardX
	guardY := rootGuardY
	loopsMap := make(map[string]struct{}) // maps in go have O(1) lookup time while slices (lists) have O(n)
	loops := make([][]int, 0)
	blocks := make(map[string]struct{})

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
			positionKey := fmt.Sprintf("%d,%d,%c", nextX, nextY, currentDirection)
			blocks[positionKey] = struct{}{}

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

			key := fmt.Sprintf("%d,%d", nextX, nextY)

			if _, exists := loopsMap[key]; !exists {
				if isLooped(guardX, guardY, currentDirection, lines, nextX, nextY) {
					loops = append(loops, []int{nextX, nextY})
				}
			}

			loopsMap[key] = struct{}{}

			guardX = nextX
			guardY = nextY
		}
	}

	elapsed := time.Since(start)

	printGrid(lines, loops)
	fmt.Println("Loops:", len(loops))
	fmt.Println("Took:", elapsed)
}

func printGrid(lines [][]rune, loops [][]int) {
	// print grid normall but mark loops postions with *
	for y, line := range lines {
		for x, c := range line {
			isLoop := false
			for _, pos := range loops {
				if pos[0] == x && pos[1] == y {
					isLoop = true
					break
				}
			}

			if isLoop {
				fmt.Print("O")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}

func isLooped(guardX, guardY int, currentDirection rune, lines [][]rune, blockX, blockY int) bool {

	block := '#'
	blocks := make(map[string]struct{})

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
		if lines[nextY][nextX] == block || (nextX == blockX && nextY == blockY) {

			positionKey := fmt.Sprintf("%d,%d,%c", nextX, nextY, currentDirection)

			if _, exists := blocks[positionKey]; exists {
				return true // we found it!
			}

			blocks[positionKey] = struct{}{}

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
		}
	}

	return false
}
