package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"time"
)


func main() {
    start := time.Now()
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    acoef := make(map[int]int)
    bcoef := make(map[int]int)
    bound := 4000000

    radius := make(map[image.Point]int)

    for buffer.Scan() {
        format := "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d"
        var sX int
        var sY int
        var bX int
        var bY int

        fmt.Sscanf(buffer.Text(), format, &sX, &sY, &bX, &bY)
        r := abs(sX - bX) + abs(sY - bY)

        acoef[sY-sX+r+1]++
        acoef[sY-sX-r-1]++
        bcoef[sX+sY+r+1]++
        bcoef[sX+sY-r-1]++
        radius[image.Pt(sX, sY)] = r
    }

    for a := range acoef {
        for b := range bcoef {
            p := image.Point{(b-a)/2, (a+b)/2}
            flawless := true
            if (p.X > 0 && p.X < bound && p.Y > 0 && p.Y < bound) {
                for sensor, radius := range radius {
                    if distance(p, sensor) <= radius {
                        flawless = false
                        break
                    }
                }

                if flawless {
                    fmt.Println(4000000*p.X+p.Y)
                }
            }
        }
    }

    fmt.Println(time.Since(start))
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
