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
    head := Point{x:0, y: 0}
    tail := Point{x:0, y: 0}

    locations := make(map[string]int)

    for buffer.Scan() {
        position, moves := parse(buffer.Text())

        locations[fmt.Sprint(tail.x,tail.y)]++
        for i := 0; i < moves; i++ {
            switch position {
            case "U":
                head.y++
            case "D":
                head.y--
            case "L":
                head.x--
            case "R":
                head.x++
            }

            deltaX := head.x - tail.x
            deltaY := head.y - tail.y

            if (Abs(deltaX) > 1) {
                if (deltaX > 0) {
                    tail.x++
                } else {
                    tail.x--
                }
                if deltaY != 0 {
                    tail.y = head.y
                }
            }

            if (Abs(deltaY) > 1) {
                if (deltaY > 0) {
                    tail.y++
                } else {
                    tail.y--
                }
                if deltaY != 0 {
                    tail.x = head.x
                }
            }

            locations[fmt.Sprint(tail.x,tail.y)]++
        }
    }

    fmt.Println(len(locations))
}

func parse(input string) (string, int) {
    splitted := strings.Split(input, " ")
    value, _ := strconv.Atoi(splitted[1])
    return splitted[0], value
}

type Point struct {
    x int
    y int
}

func Abs(v int) int {
    if (v > 0) {
        return v
    }

    return -v
}
