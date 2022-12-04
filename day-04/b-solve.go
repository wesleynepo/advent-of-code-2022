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

    var count int

    for buffer.Scan() {
        a, b := parse(buffer.Text())

        if (a.isOverlaped(b) || b.isOverlaped(a)) {
            count++
        }
    }

    fmt.Println(count)
}

func parse(line string) (*Section, *Section) {
    splitted := strings.Split(line, ",")

    return convertToSection(splitted[0]), convertToSection(splitted[1])
}

func convertToSection(section string) (*Section) {
    splitted := strings.Split(section, "-")
    lower, _ := strconv.Atoi(splitted[0])
    upper, _:= strconv.Atoi(splitted[1])
    return &Section{lower, upper}
}

type Section struct {
    lower int
    upper int
}

func (s *Section) isOverlaped(other *Section) bool {
    return s.lower >= other.lower && s.lower <= other.upper || s.upper >= other.lower && s.upper <= other.upper
}
