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
    var count int

    for scan.Scan() {
        value, err := strconv.Atoi(scan.Text())
        if (err != nil) {
            if (count > highestCountCalories) {
                highestCountCalories = count
            }
            count = 0
        } else {
            count += value
        }
    }

    fmt.Println(highestCountCalories)
}
