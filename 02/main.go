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

    var green, blue, red int = 13, 14, 12

    sum := 0
    lines := strings.Split(input, "\n")
    // parse input
    for _, line := range lines {
        if len(line) > 0 {
            // parse game id
            game := strings.Split(line, ":")
            // extract the number from the string
            game_id, _ := strconv.Atoi(game[0][strings.Index(game[0], " ")+1:])

            r, g, b := 0, 0, 0
            invalid := false

            // split each line
            for _, v := range strings.Split(game[1], ";") {
                if (!invalid) {
                    // split each game
                    for _, w := range strings.Split(v, ",") {
                        // split each color
                        if strings.Contains(w, "green") {
                            // extract green count
                            g = strings.Index(w, " green")
                            g, _ = strconv.Atoi(w[1:g])
                        } else if strings.Contains(w, "blue") {
                            // extract blue count
                            b = strings.Index(w, " blue")
                            b, _ = strconv.Atoi(w[1:b])
                        } else if strings.Contains(w, "red") {
                            // extract red count
                            r = strings.Index(w, " red")
                            r, _ = strconv.Atoi(w[1:r])
                        }
                    }
                    // check if the game is valid
                    invalid = r > red || g > green || b > blue
                }
            }

            // if the game is valid, add the game id to the sum
            if !invalid {
                sum += game_id
            }
        }
    }

    println(sum)
}
