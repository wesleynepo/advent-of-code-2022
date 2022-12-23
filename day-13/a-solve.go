package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    pairs := make([]any, 0)

    for buffer.Scan() {
        value := buffer.Text()

        if value == "" {
            continue
        }
        var a any
        json.Unmarshal([]byte(value), &a)

        pairs = append(pairs, a)
    }
    fmt.Println(pairs)

    count := 0
    for i := 0; i < len(pairs); i+= 2 {
        if cmp(pairs[i], pairs[i+1]) <= 0 {
            count += i /2 + 1
        }
    }

    fmt.Println(count)
}

func cmp(a, b any) int {
    as, aok := a.([]any)
    bs, bok := b.([]any)

    switch {
    case !aok && !bok:
        return int(a.(float64) - b.(float64))
    case !aok:
        as = []any{a}
    case !bok:
        bs = []any{b}
    }

    for i := 0; i < len(as) && i < len(bs); i++ {
        if c := cmp(as[i], bs[i]); c != 0 {
            return c
        }
    }

    return len(as) - len(bs)
}
