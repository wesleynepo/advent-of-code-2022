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

    for buffer.Scan() {
        rucksack := buffer.Text()
        left := rucksack[0:len(rucksack)/2]
        right := rucksack[len(rucksack)/2:]
        dic := hash(left)

        for _, rune := range right {
            if (dic[rune]) {
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
