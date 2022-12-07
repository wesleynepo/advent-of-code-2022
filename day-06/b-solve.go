package main

import (
	"bufio"
	"fmt"
	"os"
)

const MARKER_SIZE = 14

func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    var packets int

    for buffer.Scan() {
        packets+= getMarker(buffer.Text())
    }


    fmt.Println(packets)
}

func getMarker(line string) int {
    for i := 0; i < len(line); i++ {
        keys := make(map[rune]bool)

        for _, rune := range line[i:i+MARKER_SIZE] {
            keys[rune] = true
        }

        if(len(keys) == MARKER_SIZE) {
            return i + MARKER_SIZE
        }
    }

    return 0
}

