package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
    number int
    items []int
    operation func(int) int
    divisible int
    falseTo int
    trueTo int
    interactions int
}

func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    monkeys := make([]*Monkey, 0)

    for buffer.Scan() {
        value := buffer.Text()

        if strings.Contains(value, "Monkey") {
            var id int
            fmt.Sscanf(value, "Monkey %d:", &id)
            //Second line
            buffer.Scan()
            items := convertItems(buffer.Text())
            //Third line
            buffer.Scan()
            operation := getOperation(buffer.Text())
            buffer.Scan()
            var divisor int
            fmt.Sscanf(buffer.Text(), "  Test: divisible by %d", &divisor)
            buffer.Scan()
            var trueMonkey int
            fmt.Sscanf(buffer.Text(), "    If true: throw to monkey %d", &trueMonkey)
            buffer.Scan()
            var falseMonkey int
            fmt.Sscanf(buffer.Text(), "    If false: throw to monkey %d", &falseMonkey)

            monkey := Monkey{number: id, items: items, operation: operation, divisible: divisor, falseTo: falseMonkey, trueTo: trueMonkey}

            monkeys = append(monkeys, &monkey)
        }
    }

    bored := func (v int) int { return v/3 }
    inspect(monkeys, 20, bored)

    sort.SliceStable(monkeys, func(i, j int) bool {
        return monkeys[i].interactions > monkeys[j].interactions
    })

    fmt.Println(monkeys[0].interactions * monkeys[1].interactions)
}

func inspect(monkeys []*Monkey, count int, bored func(int) int) {

    for i := 0; i < count; i++ {
        for _, monkey := range monkeys {
            for _, item := range monkey.items {
                worry := monkey.operation(item)
                worry = bored(worry)
                if (worry % monkey.divisible == 0) {
                    monkeys[monkey.trueTo].items = append(monkeys[monkey.trueTo].items, worry)
                } else {
                    monkeys[monkey.falseTo].items = append(monkeys[monkey.falseTo].items, worry)
                }

                monkey.interactions++
            }
            monkey.items = nil
        }
    }
}

func getOperation(input string) func(int) int {
    sanitazed := strings.ReplaceAll(input, "* old", "^ 2")
    splitted := strings.Split(sanitazed, " ")
    value, _ := strconv.Atoi(splitted[len(splitted) -1])

    operations := map[string]func(int) int{
        "+": func(v int) int { return  v + value },
        "-": func(v int) int { return  v - value },
        "*": func(v int) int { return  v * value },
        "^": func(v int) int { return  v * v },
    }

    return operations[splitted[len(splitted)-2]]
}

func convertItems(input string) []int {
    items := strings.Split(strings.Split(input, ":")[1], ",")
    return convertInt(items)
}

func convertInt(items []string) []int {
    values := make([]int, 0)
    for _, v := range items {
        num, _ := strconv.Atoi(strings.TrimSpace(v))
        values = append(values, num)
    }

    return values
}
