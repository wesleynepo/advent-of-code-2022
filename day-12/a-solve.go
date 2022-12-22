package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)


func main() {
    input, _ := os.Open("./test")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    var start, end image.Point
    height := map[image.Point]rune{}

    count := 0
    for buffer.Scan() {
        for i, rune := range buffer.Text() {
            height[image.Point{count, i}] = rune

            if rune == 'S' {
                start = image.Point{count, i}
            } else if rune == 'E' {
                end = image.Point{count, i}
            }
        }
        count++
    }

    height[start], height[end] = 'a', 'z'

    dist := map[image.Point]int{end: 0}
    a, b := dist[image.Point{0,1}]
    fmt.Println(a, b)
    queue := []image.Point{end}
    var shortest *image.Point

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        if height[cur] == 'a' && shortest == nil {
            shortest = &cur
        }

        for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
            next := cur.Add(d)
            _, seen := dist[next]
            _, valid := height[next]

            if !seen && valid && height[cur] <= height[next]+1 {
                dist[next] = dist[cur] + 1
                queue = append(queue, next)
            }
        }
    }
    fmt.Println(dist[start])
    fmt.Println(dist[*shortest])
}
