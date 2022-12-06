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

    container := make(map[int][]string);

    //Parse structure
    for buffer.Scan() {
        line := buffer.Text()

        if (line == "") {
            break
        }

        for i, r := range line {
            if (r >= 'A' && r <= 'Z') {
                column := getColumn(i)
                if _, ok := container[column]; !ok {
                    container[column] = make([]string, 0)
                }
                container[column] = append(container[column], string(r))
            }
        }
    }

    for buffer.Scan() {
        action := strings.Split(buffer.Text(), " ")
        amount, _ := strconv.Atoi(action[1])
        from, _:= strconv.Atoi(action[3])
        to, _ := strconv.Atoi(action[5])

        removed := container[from][0:amount]
        container[from] = append([]string{}, container[from][amount:]...)
        container[to] = append(removed, container[to]...)
    }

    for i := 1; i <= len(container); i++ {
        fmt.Print(container[i][0:1])
    }

}

func getColumn(position int) (int) {
    if (position == 1) {
        return 1
    }

    return (position / 4) + 1
}
