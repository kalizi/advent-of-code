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

func indexLookup(position int, line string) int {
    lookup := [9] string {
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    }

    for l := 0; l < 9; l++ {
        if strings.HasPrefix(line[position:], lookup[l]) {
            return l + 1
        }
    }

    return -1
}

func main() {
    str := readInput()

    sum := 0
	lines := strings.Split(str, "\n")

	for _, line := range lines {
        lineLen := len(line)
        if lineLen > 0 {
            var num, decimal int

            i := 0
            lookup := indexLookup(i, line)
            for (line[i] <= '0' || line[i] > '9') && (lookup == -1) {
                i++
                lookup = indexLookup(i, line)
            }
            if lookup != -1 {
                decimal = lookup
            } else {
                decimal, _ = strconv.Atoi(string(line[i]))
            }

            i = lineLen - 1
            lookup = indexLookup(i, line)
            for (line[i] <= '0' || line[i] > '9') && (lookup == -1) {
                i--
                lookup = indexLookup(i, line)
            }
            if lookup != -1 {
                num = lookup
            } else {
                num, _ = strconv.Atoi(string(line[i]))
            }

            sum += decimal * 10 + num
        }
	}

	fmt.Println(sum)
}


