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

	nums := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		num, _ := strconv.Atoi(text)

		nums = append(nums, num)
	}

	fmt.Println(nums)

	result := 0

	for _, secret := range nums {
		newsecret := secret
		for i := 0; i < 2000; i++ {
			next := newsecret * 64
			newsecret = newsecret ^ next
			newsecret = newsecret % 16777216
			next = newsecret / 32
			newsecret = newsecret ^ next
			newsecret = newsecret % 16777216
			next = newsecret * 2048
			newsecret = newsecret ^ next
			newsecret = newsecret % 16777216
		}

		result += newsecret
	}

	elapsed := time.Since(start)

	fmt.Println("Result:", result)
	fmt.Println("Took:", elapsed)
}
