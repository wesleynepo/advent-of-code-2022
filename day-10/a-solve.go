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

    fmt.Println(cycles)
}


