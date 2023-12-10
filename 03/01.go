package main

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "regexp"
) 

func readInput() string {
    str, err := os.ReadFile("input.txt")
    if err != nil {
        println("Error reading file")
        panic(err)
    }
    return string(str)
}

func main() {
    input := readInput()
// Test input from reddit:
// Thanks to https://www.reddit.com/r/adventofcode/comments/189q9wv/2023_day_3_another_sample_grid_to_use/
//     input := `12.......*..
// +.........34
// .......-12..
// ..78........
// ..*....60...
// 78..........
// .......23...
// ....90*12...
// ............
// 2.2......12.
// .*.........*
// 1.1.......56
// `

// Default input:
//     input := `467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..`

    sum := 0

    // remove all \n
    lines := strings.Split(input, "\n")
    // check if the last line is empty
    if len(lines[len(lines)-1]) == 0 {
        lines = lines[:len(lines)-1]
    }

    for i, line := range lines {
        numberStart := -1
        numberEnd := -1

        for j, c := range line {
            if c >= '0' && c <= '9' {
                if numberStart == -1 {
                    numberStart = j
                }
                numberEnd = j
            } 
            if numberStart != -1 && (numberEnd != j || j == len(line) - 1) {
                currentNumber, _ := strconv.Atoi(line[numberStart:numberEnd+1])

                re := regexp.MustCompile("^[.0-9]+$")

                valid := false

                // up/down start/end
                start := numberStart
                end := numberEnd

                // check left and right
                if numberStart > 0 {
                    valid = line[numberStart - 1] != '.'
                    start--
                } 

                if !valid && numberEnd < len(line) - 1 {
                    valid = line[numberEnd + 1] != '.'
                    end++
                }

                // up
                if !valid && i > 0 {
                    // println("up", currentNumber, i, start, end)
                    valid = !re.MatchString(lines[i-1][start:end+1])
                }

                // down
                if !valid && i < len(lines) - 1 {
                    valid = !re.MatchString(lines[i+1][start:end+1])
                }

                if valid {
                    sum += currentNumber
                }

                numberStart = -1
                numberEnd = -1
            }
        }
    }
    fmt.Println(sum)
}
