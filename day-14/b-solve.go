package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strings"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)
    points := make(map[image.Point]string, 0)

    lowest := 0
    for buffer.Scan() {
        value := make([]image.Point,0)

        for _, point := range strings.Split(buffer.Text(), "->") {
            var column int
            var row int

            fmt.Sscanf(point, "%d,%d", &column, &row)

            if (row > lowest) {
                lowest = row
            }

            value = append(value, image.Point{column, row})
        }

        for i := 0; i < len(value) -1; i++ {
            diff := value[i+1].Sub(value[i])

            if diff.X == 0 {
                connect(points,value[i], value[i+1], image.Point{0,1}, diff.Y)
            } else {
                connect(points,value[i], value[i+1], image.Point{1,0}, diff.X)
            }
        }
    }

    dirs := []image.Point{{0,1},{-1,1},{1,1}}

    baseline := image.Pt(0, lowest+2)
    connect(points, baseline, baseline.Add(image.Pt(1000, 0)), image.Pt(1,0), 1)

    count := 0
    finished := false

    for !finished {
        sand := image.Point{500,0}
        for true {
            before := sand
            for _, dir := range dirs {
                curr := sand.Add(dir)
                _, has := points[curr]

                if (!has) {
                    sand = curr
                    break
                }
            }

            if (before == sand) {
                if (sand.Y == 0) {
                    finished = true
                }
                points[sand] = "O"
                count++
                break
            }
        }
    }

    fmt.Println(count)
}



func connect(points map[image.Point]string, position, destination, move image.Point, moves int) {
        points[position] = "#"
        for position != destination {
            if (moves > 0) {
                position = position.Add(move)
            } else {
                position = position.Sub(move)
            }
            points[position] = "#"
        }
}
