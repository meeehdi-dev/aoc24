package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func key(a int, b int) string {
	s := []string{strconv.Itoa(a), strconv.Itoa(b)}
	return strings.Join(s, "_")
}

type Direction int

const (
	top          = 0
	top_right    = 1
	right        = 2
	bottom_right = 3
	bottom       = 4
	bottom_left  = 5
	left         = 6
	top_left     = 7
)

func check(m map[string]string, x int, y int, direction Direction, search string) int {
	found := 0

	if m[key(y, x)] == "" {
		// fmt.Println("NOPE!", search, x, y)
		return found
	}

	if m[key(y, x)] != string(search[0]) {
		// fmt.Println("WRONG!", search, x, y)
		return found
	}

	next := search[1:]
	if len(next) == 0 {
		// fmt.Println("OK!", search, m[key(y, x)], key(y, x), x, y)
		return found + 1
	}

	if direction == top {
		found += check(m, x, y-1, direction, next)
	} else if direction == top_right {
		found += check(m, x+1, y-1, direction, next)
	} else if direction == right {
		found += check(m, x+1, y, direction, next)
	} else if direction == bottom_right {
		found += check(m, x+1, y+1, direction, next)
	} else if direction == bottom {
		found += check(m, x, y+1, direction, next)
	} else if direction == bottom_left {
		found += check(m, x-1, y+1, direction, next)
	} else if direction == left {
		found += check(m, x-1, y, direction, next)
	} else if direction == top_left {
		found += check(m, x-1, y-1, direction, next)
	}

	return found
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d04/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	result := 0
	length := 0
	m := make(map[string]string)
	y := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		x := 0
		length = len(line)

		for range length {
			m[key(x, y)] = string(line[x])
			x++
		}

		y++
	}

	xmas := 0
	for y := range length {
		for x := range length {
			result += check(m, x, y, top, "XMAS")
			result += check(m, x, y, top_right, "XMAS")
			result += check(m, x, y, right, "XMAS")
			result += check(m, x, y, bottom_right, "XMAS")
			result += check(m, x, y, bottom, "XMAS")
			result += check(m, x, y, bottom_left, "XMAS")
			result += check(m, x, y, left, "XMAS")
			result += check(m, x, y, top_left, "XMAS")

			if m[key(y, x)] == "A" {
				// top
				if m[key(y-1, x-1)] == "M" && m[key(y-1, x+1)] == "M" && m[key(y+1, x-1)] == "S" && m[key(y+1, x+1)] == "S" {
					xmas++
				}
				// right
				if m[key(y-1, x+1)] == "M" && m[key(y+1, x+1)] == "M" && m[key(y-1, x-1)] == "S" && m[key(y+1, x-1)] == "S" {
					xmas++
				}
				// bottom
				if m[key(y-1, x-1)] == "S" && m[key(y-1, x+1)] == "S" && m[key(y+1, x-1)] == "M" && m[key(y+1, x+1)] == "M" {
					xmas++
				}
				// left
				if m[key(y-1, x+1)] == "S" && m[key(y+1, x+1)] == "S" && m[key(y-1, x-1)] == "M" && m[key(y+1, x-1)] == "M" {
					xmas++
				}
			}
		}
	}

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("X-MAS: %d\n", xmas)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
