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

    var topCalories = make([]int, 3)

    handleCount(scan, topCalories)

    fmt.Println(topCalories[0] + topCalories[1] + topCalories[2])
}

func handleCount(buffer *bufio.Scanner, topCalories []int) {
    var count int
    testAndAssing := func () {
        for place, calorie := range topCalories {
            if (calorie < count) {
                topCalories[place] = count
                break
            }
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
