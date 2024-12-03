package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.ReadFile("d03/input.txt")
	if err != nil {
		panic("file not found")
	}
	lines := strings.Split(string(file), "\n")

	result := 0
	enabled := true
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		index := 0
		mul := false
		p1 := false
		num1 := ""
		c := false
		num2 := ""
		chars := strings.Split(line, "")
		for range chars {
			if index >= len(line) {
				break
			}

			if enabled && index+7 < len(line) {
				tok := line[index : index+7]
				if tok == "don't()" {
					enabled = false
					index += 7
					continue
				}
			}
			if !enabled && index+4 < len(line) {
				tok := line[index : index+4]
				if tok == "do()" {
					enabled = true
					index += 4
					continue
				}
			}

			if enabled {
				if !mul && index+3 < len(line) {
					op := line[index : index+3]
					if op == "mul" {
						mul = true
						index += 3
						continue
					}
				} else {
					if !p1 {
						char := line[index]
						if char == '(' {
							p1 = true
							index++
							continue
						}
					} else {
						if !c {
							char := line[index]
							if char == ',' {
								if len(num1) > 0 {
									c = true
									index++
									continue
								}
							} else {
								s := string(char)
								_, err := strconv.Atoi(s)
								if err == nil {
									num1 += s
									index++
									continue
								}
							}
						} else {
							char := line[index]
							if char == ')' {
								if len(num2) > 0 {
									n1, _ := strconv.Atoi(num1)
									n2, _ := strconv.Atoi(num2)
									result += n1 * n2
								}
							} else {
								s := string(char)
								_, err := strconv.Atoi(s)
								if err == nil {
									num2 += s
									index++
									continue
								}
							}

						}
					}
				}
			}

			// reset
			mul = false
			p1 = false
			num1 = ""
			c = false
			num2 = ""
			index++
		}
	}

	fmt.Printf("Result: %d\n", result)

	elapsed := time.Since(start)
	fmt.Printf("[%.2fms]\n", float64(elapsed)/1000000)
}
