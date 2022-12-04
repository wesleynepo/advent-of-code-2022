package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
    WIN = 6
    TIE = 3
    LOSS = 0
)

const (
    Rock = 1
    Paper = 2
    Scissor = 3
)

func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    scan := bufio.NewScanner(input)

    var score int

    for scan.Scan() {
        left, right := parse(scan.Text())

        score += right.value + right.compare(left)
    }

    fmt.Println(score)
}

func parse(text string) (Hand, Hand) {
    return mapToHand(text[0:1]), mapToHand(text[2:3])
}

func mapToHand(letter string) Hand {
    switch {
    case letter == "A" || letter == "X":
        return Hand{value: Rock, winAgainst: Scissor}
    case letter == "B" || letter == "Y":
        return Hand{value: Paper, winAgainst: Rock}
    case letter == "C" || letter == "Z":
        return Hand{value: Scissor, winAgainst: Paper}
    default:
        panic("Error")
    }
}

type Hand struct {
    value int
    winAgainst int
}

func (h Hand) compare(other Hand) int {
    switch other.value {
    case h.value:
        return TIE
    case h.winAgainst:
        return WIN
    default:
        return LOSS
    }
}
