package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)


func main() {
    input, _ := os.Open("./input")
    defer input.Close()

    buffer := bufio.NewScanner(input)

    valves := map[string]Valve{}

    for buffer.Scan() {
        spaces := strings.Split(strings.ReplaceAll(buffer.Text(), ",", ""), " ")

        var rate int
        fmt.Sscanf(spaces[4], "rate=%d;", &rate)
        valves[spaces[1]] = Valve{rate:rate, lead: spaces[9:], name: spaces[1]}
    }

    open := map[string]bool{"AA": true}

    for name, valve := range valves {
        if valve.rate == 0 {
            open[name] = true
        }
    }


    solution := bfs(valves, "AA", 30, 0, open, map[string]int{})
    fmt.Println(solution)
}

func hash(currentRoom string, minutesLeft int, open map[string]bool, currentPressure int) string {
	rms := []string{}
	for k := range open {
		rms = append(rms, k)
	}
	sort.Strings(rms)
	return fmt.Sprint(currentRoom, minutesLeft, rms, currentPressure)
}

func bfs(graph map[string]Valve, currentRoom string, minutesLeft, currentPressure int, open map[string]bool, memo map[string]int) int {
    if minutesLeft == 0 {
        return 0
    }

    key := hash(currentRoom, minutesLeft, open, currentPressure)

    if v, ok := memo[key]; ok {
        return v
    }

    bestFlow := 0

    if !open[currentRoom] {
        open[currentRoom] = true

        newPressure := currentPressure + graph[currentRoom].rate

        maybeBest := currentPressure + bfs(graph, currentRoom, minutesLeft-1, newPressure, open, memo)

        if maybeBest > bestFlow {
            bestFlow = maybeBest
        }

        open[currentRoom] = false
    }

    for _, neighbor := range graph[currentRoom].lead {
        maybeBest := currentPressure + bfs(graph, neighbor, minutesLeft-1, currentPressure, open, memo)

        if maybeBest > bestFlow {
            bestFlow = maybeBest
        }
    }

    memo[key] = bestFlow

    return bestFlow
}

func pressure(valves map[string]Valve, opened []string) int {
    total := 0
    for _, open := range opened {
        total += valves[open].rate
    }

    return total
}

func notin(valves []string, valve string) bool {
    for _, v := range valves {
        if v == valve {
            return false
        }
    }

    return true
}

type Seen struct {
    time int
    where string
}

type State struct {
    time int
    score int
    where string
    opened []string
}


type Valve struct {
    name string
    rate int
    lead []string
}
