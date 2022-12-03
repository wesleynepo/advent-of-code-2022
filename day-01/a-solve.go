package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    scan := bufio.NewScanner(input)

    var highestCountCalories int
    handleCount(scan, &highestCountCalories)

    fmt.Println(highestCountCalories)
}

func handleCount(buffer *bufio.Scanner, highestCountCalories *int) {
    var count int
    testAndAssing := func () {
        if (count > *highestCountCalories) {
            *highestCountCalories = count
        }
        count = 0
    }

    // last input doesn't end with space so count is full at end
    defer testAndAssing()

    for buffer.Scan() {
        value, err := strconv.Atoi(buffer.Text())
        if (err != nil) {
            testAndAssing()
        } else {
            count += value
        }
    }
}

