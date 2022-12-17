package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    cycles := make([]int, 0)
    x := 1

    for buffer.Scan() {
        instruction := strings.Split(buffer.Text(), " ")
        cycles = append(cycles, x)

        if instruction[0] == "addx" {
            cycles = append(cycles, x)
            value, _ := strconv.Atoi(instruction[1])
            x += value
        }
    }

    cycles = append(cycles, x)
    for i, v := range cycles {
        current := i % 40
        if (current == 39) {
            fmt.Println(isVisible(current, v))
        } else {
            fmt.Print(isVisible(current, v))
        }
    }
}


func isVisible(crt int, cycle int) string {
    if (crt >= cycle -1 && crt <= cycle +1) {
        return "#"
    }

    return "."
}

