package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    pkgs := make([]any, 0)

    for buffer.Scan() {
        value := buffer.Text()

        if value == "" {
            continue
        }
        var a any
        json.Unmarshal([]byte(value), &a)

        pkgs= append(pkgs, a)
    }

    pkgs = append(pkgs, []any{[]any{2.}})
    pkgs = append(pkgs, []any{[]any{6.}})

    sort.Slice(pkgs, func(i, j int) bool {
        return cmp(pkgs[i],pkgs[j]) < 0
    })

    key := 1
    for i, pkg := range pkgs {
        printed := fmt.Sprint(pkg)
        if printed == "[[2]]" || printed == "[[6]]" {
            key *= i + 1
        }
    }

    fmt.Println(key)
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
