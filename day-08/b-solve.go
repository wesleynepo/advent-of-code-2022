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

    grid := make([][]int, 0)

    for buffer.Scan() {
        line := buffer.Text()
        row := strings.Split(line, "")
        integers := convertRow(row)
        grid = append(grid, integers)
    }

    var sum int

    calculate(grid, transpose(grid), &sum)
    fmt.Println(sum)
}

func calculate(matrix [][]int, transposed [][]int, sum *int) {
    rowLength, columnLength := lengths(matrix)

    for i := 1; i < columnLength - 1; i++ {
        for j := 1; j < rowLength - 1; j++ {
            actual := matrix[i][j]
            leftSlice := matrix[i][0:j]
            rightSlice := matrix[i][j+1:]
            upSlice := transposed[j][0:i]
            downSlice := transposed[j][i+1:]


            count := 1
            count*= distanceBackward(leftSlice, actual)
            count*= distanceForward(rightSlice, actual)
            count*= distanceBackward(upSlice, actual)
            count*= distanceForward(downSlice, actual)

            if (count > *sum) {
                *sum = count
            }

        }
    }
}

func distanceForward(slice []int, value int) int {
    var distance int

    for i := 0; i < len(slice); i++{
        distance++
        if slice[i] >= value {
            break
        }
    }
    return distance
}

func distanceBackward(slice []int, value int) int {
    var distance int

    for i := len(slice)-1; i >= 0; i--{
        distance++
        if slice[i] >= value {
            break
        }
    }
    return distance
}

func lengths(matrix [][]int) (int, int) {
    return len(matrix[0]), len(matrix)
}

func convertRow(row []string) []int {
    rowConverted := make([]int, 0)
    for _, value := range row {
        integer, _ := strconv.Atoi(value)
        rowConverted = append(rowConverted, integer)
    }

    return rowConverted
}

func transpose(matrix [][]int) [][]int {
    rowLength, columnLength := lengths(matrix)
    transposed := make([][]int, rowLength)
    for i := range transposed {
        transposed[i] = make([]int, columnLength)
    }

    for i := 0; i < rowLength; i++ {
        for j := 0; j < columnLength; j++ {
            transposed[i][j] = matrix[j][i]
        }
    }

    return transposed
}
