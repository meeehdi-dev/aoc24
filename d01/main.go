package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.ReadFile("d01/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	var left []int
	var right []int
	counts := make(map[int]int)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		numbers := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(numbers[0])
		rightNum, _ := strconv.Atoi(numbers[1])

		_, ok := counts[rightNum]
		if !ok {
			counts[rightNum] = 0
		}
		counts[rightNum]++

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	similarity := 0
	for index := range left {
		leftNum := left[index]
		v, ok := counts[leftNum]
		if ok {
			similarity += leftNum * v
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for index := range left {
		leftNum := left[index]
		rightNum := right[index]
		distance := int(math.Abs(float64(rightNum - leftNum)))
		totalDistance += distance
	}

	fmt.Printf("Distance: %d\n", totalDistance)
	fmt.Printf("Similarity: %d\n", similarity)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
