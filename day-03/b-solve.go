package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    var totalPriority int
    var rucksacks []string

    for buffer.Scan() {
        rucksacks = append(rucksacks, buffer.Text())
    }

    for i := 2; i < len(rucksacks); i+= 3 {
        elfOne := hash(rucksacks[i-2])
        elfTwo := hash(rucksacks[i-1])

        for _, rune := range rucksacks[i] {
            if (elfOne[rune] && elfTwo[rune]) {
                if (rune > 90) {
                    totalPriority += int(rune) - 96
                } else {
                    totalPriority += int(rune) - 38
                }
                break
            }
        }
    }

    fmt.Println(totalPriority)
}

func hash(rucksack string) map[rune]bool {
    dic := make(map[rune]bool)

    for _,letter := range rucksack {
        dic[letter] = true
    }

    return dic
}
