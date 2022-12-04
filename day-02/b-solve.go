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
        hand, result := parse(scan.Text())

        score += hand.score(result)
    }

    fmt.Println(score)
}

func parse(text string) (Hand, string) {
    return mapToHand(text[0:1]), text[2:3]
}

func mapToHand(letter string) Hand {
    switch {
    case letter == "A":
        return Hand{value: Rock, win:Scissor, lose: Paper}
    case letter == "B":
        return Hand{value: Paper, win:Rock, lose: Scissor}
    case letter == "C":
        return Hand{value: Scissor, win:Paper, lose: Rock}
    default:
        panic("Error")
    }
}

type Hand struct {
    value int
    lose int
    win int
}

func (h Hand) score(result string) int {
    switch result {
    case "Z":
        return h.lose + WIN
    case "Y":
        return h.value + TIE
    default:
        return h.win + LOSS
    }
}
