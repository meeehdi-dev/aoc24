package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Direction int

const (
	top    = 0
	right  = 2
	bottom = 4
	left   = 6
)

func key(a int, b int) string {
	s := []string{strconv.Itoa(a), strconv.Itoa(b)}
	return strings.Join(s, "_")
}

func main() {
	start := time.Now()

	file, err := os.ReadFile("d06/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	result := 1
	obstacles := 0
	length := 0
	m := make(map[string]string)
	y := 0
	currentX := -1
	currentY := -1
	dir := top
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		x := 0
		length = len(line)

		for range length {
			k := key(x, y)
			m[k] = string(line[x])

			if m[k] == "^" {
				currentX = x
				currentY = y
				m[k] = "u"
			}

			x++
		}

		y++
	}

	o := make(map[string]bool)
	for true {
		if currentX < 0 || currentX >= length || currentY < 0 || currentY >= length {
			break
		}

		if dir == top {
			k := key(currentX, currentY-1)
			if m[k] == "#" {
				dir = right
			} else {
				if m[k] == "." {
					m[k] = "u"
					result++
				} else if strings.Index(m[k], "u") == -1 {
					m[k] = m[k] + "u"
				}
				if !o[key(currentX, currentY-1)] {
					nextX := currentX + 1
					for true {
						next := m[key(nextX, currentY)]
						if nextX >= length || next == "#" {
							break
						}
						if strings.Index(next, "r") != -1 {
							o[key(currentX, currentY-1)] = true
							obstacles++
							break
						}
						nextX++
					}
				}
				currentY--
			}
		}
		if dir == right {
			k := key(currentX+1, currentY)
			if m[k] == "#" {
				dir = bottom
			} else {
				if m[k] == "." {
					m[k] = "r"
					result++
				} else if strings.Index(m[k], "r") == -1 {
					m[k] = m[k] + "r"
				}
				if !o[key(currentX+1, currentY)] {
					nextY := currentY + 1
					for true {
						next := m[key(currentX, nextY)]
						if nextY >= length || next == "#" {
							break
						}
						if strings.Index(next, "d") != -1 {
							o[key(currentX+1, currentY)] = true
							obstacles++
							break
						}
						nextY++
					}
				}
				currentX++
			}
		}
		if dir == bottom {
			k := key(currentX, currentY+1)
			if m[k] == "#" {
				dir = left
			} else {
				if m[k] == "." {
					m[k] = "d"
					result++
				} else if strings.Index(m[k], "d") == -1 {
					m[k] = m[k] + "d"
				}
				if !o[key(currentX, currentY+1)] {
					nextX := currentX - 1
					for true {
						next := m[key(nextX, currentY)]
						if nextX < 0 || next == "#" {
							break
						}
						if strings.Index(next, "l") != -1 {
							o[key(currentX, currentY+1)] = true
							obstacles++
							break
						}
						nextX--
					}
				}
				currentY++
			}
		}
		if dir == left {
			k := key(currentX-1, currentY)
			if m[k] == "#" {
				dir = top
			} else {
				if m[k] == "." {
					m[k] = "l"
					result++
				} else if strings.Index(m[k], "l") == -1 {
					m[k] = m[k] + "l"
				}
				if !o[key(currentX-1, currentY)] {
					nextY := currentY - 1
					for true {
						next := m[key(currentX, nextY)]
						if nextY < 0 || next == "#" {
							break
						}
						if strings.Index(next, "u") != -1 {
							o[key(currentX-1, currentY)] = true
							obstacles++
							break
						}
						nextY--
					}
				}
				currentX--
			}
		}
	}

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Obstacles: %d\n", obstacles)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
