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
    rope := make([]Point, 10)

    locations := map[Point]struct{}{}
    direction := map[string]Point{"U": {0, -1}, "R": {1, 0}, "D": {0, 1}, "L": {-1, 0}}

    for buffer.Scan() {
        position, moves := parse(buffer.Text())

        for i := 0; i < moves; i++ {
            rope[0].Add(direction[position])

            for i := 1; i < len(rope); i++ {
                if d := rope[i-1].Sub(rope[i]); Abs(d.x) > 1 || Abs(d.y) > 1 {
                    rope[i].Add(Point{sgn(d.x), sgn(d.y)})
                }
            }

            locations[rope[len(rope)-1]] = struct{}{}

        }
    }
    fmt.Println(rope)
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

func (p *Point) Add(move Point) {
    p.x += move.x
    p.y += move.y
}

func (p *Point) Sub(move Point) Point {
    return Point{p.x-move.x, p.y-move.y}
}

func Abs(v int) int {
    if (v > 0) {
        return v
    }

    return -v
}

func sgn(x int) int {
    if x < 0 {
        return -1
    } else if x > 0 {
        return 1
    }

    return 0
}
