package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	start := time.Now()

	file, _ := os.Open("input.txt")
	defer file.Close()

	lines := make([][]rune, 0) // y x
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		lines = append(lines, line)
	}

	var startNode, endNode Point
	for y, row := range lines {
		for x, cell := range row {
			if cell == 'S' {
				startNode = Point{X: x, Y: y}
			}
			if cell == 'E' {
				endNode = Point{X: x, Y: y}
			}
		}
	}

	graph := buildGraph(lines)

	startDirection := Point{X: 1, Y: 0} // East
	distances := Dijkstra(graph, startNode, startDirection)

	if dist, exists := distances[endNode]; exists {
		fmt.Printf("Shortest path from start to end: %d\n", dist)
	} else {
		fmt.Println("No path found from start to end.")
	}

	elapsed := time.Since(start)
	fmt.Printf("Took: %s\n", elapsed)
}

type Point struct {
	X, Y int
}

type Edge struct {
	To        Point
	Weight    int
	Direction Point
}

func buildGraph(grid [][]rune) map[Point][]Edge {
	graph := make(map[Point][]Edge)

	directions := []Point{
		{1, 0},  // East
		{0, 1},  // South
		{-1, 0}, // West
		{0, -1}, // North
	}

	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				continue
			}

			currentPos := Point{X: x, Y: y}

			for _, direction := range directions {
				newPos := Point{X: x + direction.X, Y: y + direction.Y}

				if newPos.X < 0 || newPos.X >= len(grid[0]) || newPos.Y < 0 || newPos.Y >= len(grid) {
					continue
				}

				if grid[newPos.Y][newPos.X] == '#' {
					continue
				}

				graph[currentPos] = append(graph[currentPos], Edge{
					To:        newPos,
					Weight:    1,
					Direction: direction,
				})
			}
		}
	}

	return graph
}

func Dijkstra(graph map[Point][]Edge, start Point, startDirection Point) map[Point]int {
	distances := make(map[Point]int)
	visited := make(map[Point]bool)
	prevDirections := make(map[Point]Point)

	for node := range graph {
		distances[node] = math.MaxInt
	}

	distances[start] = 0
	prevDirections[start] = startDirection

	for {
		// Find the nearest unvisited point
		var currentPoint Point
		currentDistance := math.MaxInt
		found := false
		for point, distance := range distances {
			if !visited[point] && distance < currentDistance {
				currentPoint = point
				currentDistance = distance
				found = true
			}
		}

		if !found {
			break
		}

		visited[currentPoint] = true

		// Update distances to neighbors
		for _, edge := range graph[currentPoint] {
			weight := edge.Weight

			// Determine the weight based on direction change
			prevDir := prevDirections[currentPoint]
			if edge.Direction.X != prevDir.X && edge.Direction.Y != prevDir.Y { // currentDir != prevDir, only 90 degree turn are possible, 180 would mean we go backwards
				weight += 1000
			}

			newDist := distances[currentPoint] + weight
			if newDist < distances[edge.To] {
				distances[edge.To] = newDist
				prevDirections[edge.To] = edge.Direction
			}
		}
	}

	return distances
}
