package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkSafe(levels []string) (bool, int) {
	tendency := 0
	current := 0
	for idx, levelStr := range levels {
		if len(levelStr) == 0 {
			continue
		}
		level, _ := strconv.Atoi(levelStr)

		if idx == 0 {
			current = level
			continue
		}

		if level == current {
			return false, idx
		}

		if math.Abs(float64(current)-float64(level)) > 3 {
			return false, idx
		}

		if idx == 1 {
			if level > current {
				tendency = 1
			} else {
				tendency = -1
			}
		} else {
			if level > current && tendency == -1 {
				return false, idx
			} else if level < current && tendency == 1 {
				return false, idx
			}
		}

		current = level

		if idx == len(levels)-1 {
			return true, -1
		}
	}

	fmt.Println("NOPE!")
	return false, -1
}

func concatSlice[T any](first []T, second []T) []T {
	n := len(first)
	return append(first[:n:n], second...)
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d02/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	safe := 0
	almostSafe := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		levels := strings.Split(line, " ")
		s, err := checkSafe(levels)
		if s {
			safe++
			continue
		}
		levels0 := levels[1:] // remove first element to check tendency
		s, _ = checkSafe(levels0)
		if s {
			almostSafe++
			continue
		}
		if err >= 0 { // try to remove errored index + the one before
			levels1 := concatSlice(levels[:err], levels[err+1:])
			s, _ = checkSafe(levels1)
			if s {
				almostSafe++
				continue
			}
			if err > 0 { // should never be equal to 0
				levels2 := concatSlice(levels[:err-1], levels[err:])
				s, _ = checkSafe(levels2)
				if s {
					almostSafe++
					continue
				}
			}
		}
	}

	fmt.Printf("Safe: %d\n", safe)
	fmt.Printf("Almost safe: %d\n", safe+almostSafe)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
