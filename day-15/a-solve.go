package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    y := 2000000
    positions := make(map[image.Point]int)
    beaconsAtY := make(map[image.Point]int)

    for buffer.Scan() {
        format := "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d"
        var sX int
        var sY int
        var bX int
        var bY int

        fmt.Sscanf(buffer.Text(), format, &sX, &sY, &bX, &bY)
        beacon := image.Pt(bX, bY)
        sensor := image.Pt(sX, sY)

        positions[sensor] = distance(sensor, beacon)
        if (bY == y) {
            beaconsAtY[beacon]++
        }
    }

    overlap := make(map[image.Point]int)

    for sensor, radius := range positions {
        height := abs(sensor.Y - y)
        if height > radius {
            continue
        }

        counts := (radius - height)
        center := image.Pt(sensor.X, y)
        overlap[center]++

        for i:= 1; i <= counts; i++ {
            left := image.Pt(center.X-i, y)
            right := image.Pt(center.X+i, y)

            overlap[left]++
            overlap[right]++
        }
    }
    fmt.Println(len(overlap) - len(beaconsAtY))
}

func abs(v int) int {
    if v > 0 {
        return v
    }
    return -v
}

func distance(from, to image.Point) int {
    x := abs(from.X - to.X)
    y := abs(from.Y - to.Y)
    return x + y
}
