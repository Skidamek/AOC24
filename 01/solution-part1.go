package main

import (
  "bufio"
	"fmt"
	"os"
  "strings"
  "strconv"
)


func main() {
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

    if (err != nil) {
      fmt.Println(err)
      continue
    }

    rightInt, err := strconv.Atoi(right)

    if (err != nil) {
      fmt.Println(err)
      continue
    }

    leftList = append(leftList, leftInt)
    rightList = append(rightList, rightInt)
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
  }

  fmt.Println(leftList)
  fmt.Println(rightList)


  // grab min value

  var totalDistance int = 0
  
  for len(leftList) > 0 && len(rightList) > 0 {
    leftMin, newLeftList := minValue(leftList)
    rightMin, newRightList := minValue(rightList)

    leftList = newLeftList
    rightList = newRightList

    distance := getDistance(leftMin, rightMin)

    totalDistance += distance

    fmt.Println("Left Right Distance:", leftMin, rightMin, distance)
  }

  fmt.Println("totalDistance:", totalDistance)
}

func minValue(list []int) (int, []int) {
    min := list[0]
    minIndex := 0

    for i, num := range list {
        if num < min {
            min = num
            minIndex = i
        }
    }

    newlist := removeElement(list, minIndex)
    return min, newlist
}


func removeElement(slice []int, index int) []int {
    if index < 0 || index >= len(slice) {
        return slice
    }
    
    return append(slice[:index], slice[index+1:]...)
}

func getDistance(numOne int, numTwo int) int {
  distance := numOne - numTwo

  if distance < 0 {
    return -distance
  }
  return distance
}
