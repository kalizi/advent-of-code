package main

import (
    "strings"
    "strconv"
    "os"
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

    sum := 0
    lines := strings.Split(input, "\n")
    // parse input
    for _, line := range lines {
        if len(line) > 0 {
            // parse game id
            game := strings.Split(line, ":")

            maxR, maxG, maxB := 0, 0, 0
            r, g, b := 0, 0, 0

            // split each line
            for _, v := range strings.Split(game[1], ";") {
                // split each game
                for _, w := range strings.Split(v, ",") {
                    // split each color
                    if strings.Contains(w, "green") {
                        // extract green count
                        g = strings.Index(w, " green")
                        g, _ = strconv.Atoi(w[1:g])
                        if g > maxG {
                            maxG = g
                        }
                    } else if strings.Contains(w, "blue") {
                        // extract blue count
                        b = strings.Index(w, " blue")
                        b, _ = strconv.Atoi(w[1:b])
                        if b > maxB {
                            maxB = b
                        }
                    } else if strings.Contains(w, "red") {
                        // extract red count
                        r = strings.Index(w, " red")
                        r, _ = strconv.Atoi(w[1:r])
                        if r > maxR {
                            maxR = r
                        }
                    }
                }
            }

            sum += maxR * maxG * maxB
        }
    }

    println(sum)
}
