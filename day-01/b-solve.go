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

    var topThreeCalories = make([]int, 3)
    var count int

    for scan.Scan() {
        value, err := strconv.Atoi(scan.Text())
        if (err != nil) {
            for place, calorie := range topThreeCalories {
                if (calorie < count) {
                    topThreeCalories[place] = count
                    break
                }
            }
            count = 0
        } else {
            count += value
        }
    }

    fmt.Println(topThreeCalories[0] + topThreeCalories[1] + topThreeCalories[2])
}
