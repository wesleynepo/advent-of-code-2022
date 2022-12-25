package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
)


func main() {
    input, _ := os.Open("./test")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    positions := make(map[image.Point]string)

    for buffer.Scan() {
        format := "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d"
        var sX int
        var sY int
        var bX int
        var bY int

        fmt.Sscanf(buffer.Text(), format, &sX, &sY, &bX, &bY)
        beacon := image.Pt(bX, bY)
        sensor := image.Pt(sX, sY)

        positions[beacon] = "B"
        positions[sensor] = "S"
    }


    for j := -2; j < 23; j++ {
        for i := 0; i < 26; i++ {
            value, valid := positions[image.Pt(i,j)]

            if (valid) {
                fmt.Print(value)
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }


    count := 0

    for i:= -2; i < 25; i++ {
        p := image.Pt(i, 10)
        for key := range positions {
            diff := p.Sub(key)
            squared := math.Pow(diff.X,2) + math.Pow(diff.Y,2)
            distance := math.Sqrt2()
        }
    }

}
