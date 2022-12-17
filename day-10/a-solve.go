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

    steps := []int{20, 60, 100, 140, 180, 220}

    sum := 0

    for _, v := range steps {
        sum+= (cycles[v-1] * v)
    }

    fmt.Println(sum)
}


