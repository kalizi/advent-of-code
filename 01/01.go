/*
 * Hello, I'm trying to learn Go with Advent of Code 2023.
 */

package main

import (
    "fmt"
    "strings"
    "strconv"
    "os"
)

func readInput() string {
    str, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("Error reading file")
        panic(err)
    }
    return string(str)
}

func main() {
    str := readInput()

    sum := 0
	lines := strings.Split(str, "\n")

	for _, line := range lines {
        if len(line) > 0 {
            var num, decimal int

            i := 0
            for line[i] <= '0' || line[i] > '9' {
                i++
            }
            decimal, _ = strconv.Atoi(string(line[i]))

            i = len(line) - 1
            for line[i] < '0' || line[i] > '9' {
                i--
            }
            num, _ = strconv.Atoi(string(line[i]))

            sum += decimal * 10 + num
        }
	}

	fmt.Println(sum)
}
