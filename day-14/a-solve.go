package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strings"
)


func main() {
    input, _ := os.Open("./test")
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
                step := image.Point{0, 1}
                if diff.Y < 0 {
                    step = image.Point{0, -1}
                }
                current := value[i]

                points[current] = "#"
                for current != value[i+1] {
                    current = current.Add(step)
                    points[current] = "#"
                }
            }

            if diff.Y == 0 {
                step := image.Point{1, 0}
                if diff.X < 0 {
                    step = image.Point{-1, 0}
                }
                current := value[i]

                points[current] = "#"
                for current != value[i+1] {
                    current = current.Add(step)
                    points[current] = "#"
                }
            }
        }
    }

    dirs := []image.Point{{0,1},{-1,1},{1,1}}

    count := 0
    void := true
    for void {
        sand := image.Point{500,1}
        for true {
            if (sand.Y > lowest) {
                void = false
                break
            }

            before := sand
            for _, dir := range dirs {
                curr := sand.Add(dir)
                _, has := points[curr]

                if (!has) {
                    delete(points,sand)
                    sand = curr
                    break
                }
            }
            points[sand] = "O"

            if (before == sand) {
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
