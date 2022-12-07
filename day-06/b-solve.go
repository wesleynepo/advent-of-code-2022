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

    var packets int

    for buffer.Scan() {
        line := buffer.Text()

        var found bool

        for i := 0; i < len(line); i++ {
            keys := make(map[string]bool)
            found = false

            for _, rune := range line[i:i+14] {
                letter := string(rune)

                if(keys[letter]) {
                    found = true
                    break
                }

                keys[letter] = true
            }

            if(!found) {
                packets += (i + 14)
                break
            }
        }
    }


    fmt.Println(packets)
}
