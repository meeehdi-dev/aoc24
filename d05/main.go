package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func has(nums []int, num int) bool {
	for i := range nums {
		if nums[i] == num {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d05/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	result := 0
	almost := 0
	section := 1
	order := make(map[int][]int)
	for _, line := range lines {
		if len(line) == 0 {
			if section == 1 {
				section = 2
			}
			continue
		}

		if section == 1 {
			numbers := strings.Split(line, "|")

			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			order[num1] = append(order[num1], num2)
		}
		if section == 2 {
			numbers := strings.Split(line, ",")
			error := false
			err := false
			for i := range len(numbers) {
				if error == true {
					continue
				}
				for j := i + 1; j < len(numbers); j++ {
					num1, _ := strconv.Atoi(numbers[i])
					num2, _ := strconv.Atoi(numbers[j])

					if !has(order[num1], num2) && len(order[num2]) != 0 {
						if has(order[num2], num1) {
							numbers[i], numbers[j] = numbers[j], numbers[i]
							j--
							err = true
							continue
						}
						error = true
						continue
					}

				}
			}
			if error == false {
				idx := int(len(numbers) / 2)
				mid, _ := strconv.Atoi(numbers[idx])
				if err == true {
					almost += mid
				} else {
					result += mid
				}
			}
		}
	}

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Almost: %d\n", almost)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
