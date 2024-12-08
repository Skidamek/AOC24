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

	// find non . chars in lines and get x y
	antenas := make(map[rune][][]int, 0)

	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				if _, ok := antenas[c]; !ok {
					antenas[c] = make([][]int, 0)
				}

				antenas[c] = append(antenas[c], []int{x, y})
			}
		}
	}

	antinodes := make(map[string]struct{}, 0)

	// check each antena against other antena with of the same type
	for _, pos1 := range antenas {
		for i, pos2 := range pos1 {
			for j := i + 1; j < len(pos1); j++ {
				x1 := pos2[0]
				y1 := pos2[1]
				x2 := pos1[j][0]
				y2 := pos1[j][1]

				deltaX := x2 - x1
				deltaY := y2 - y1

				opositeX1 := x1 - deltaX
				opositeY1 := y1 - deltaY

				opositeX2 := x2 + deltaX
				opositeY2 := y2 + deltaY

				if opositeX1 < 0 || opositeX1 >= len(lines[0]) || opositeY1 < 0 || opositeY1 >= len(lines) {
					opositeX1 = -1
					opositeY1 = -1
				}

				if opositeX2 < 0 || opositeX2 >= len(lines[0]) || opositeY2 < 0 || opositeY2 >= len(lines) {
					opositeX2 = -1
					opositeY2 = -1
				}

				if opositeX1 != -1 {
					antinode := fmt.Sprintf("%d,%d", opositeX1, opositeY1)
					antinodes[antinode] = struct{}{}
				}

				if opositeX2 != -1 {
					antinode := fmt.Sprintf("%d,%d", opositeX2, opositeY2)
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}

	elapsed := time.Since(start)

	printGrid(lines, antinodes)

	fmt.Println("Antinodes:", len(antinodes))
	fmt.Println("Took:", elapsed)
}

func printGrid(lines [][]rune, antinodes map[string]struct{}) {
	// print grid normall but mark loops postions with *
	for y, line := range lines {
		for x, c := range line {
			isAntinode := false
			for pos := range antinodes {
				parts := strings.Split(pos, ",")
				aX, _ := strconv.Atoi(parts[0])
				aY, _ := strconv.Atoi(parts[1])
				if aX == x && aY == y {
					isAntinode = true
					break
				}
			}

			if c == '.' && isAntinode {
				fmt.Print("#")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}
