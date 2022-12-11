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

    start := newFolder("", nil)

    currentFolder := start

    for buffer.Scan() {
        line := buffer.Text()
        commands := strings.Split(line, " ")

        if commands[1] == "cd" {
            if commands[2] == ".." {
                currentFolder = currentFolder.upperFolder
            } else {
                currentFolder = currentFolder.createDir(commands[2])
            }
        } else {
            num, err := strconv.Atoi(commands[0])

            if err != nil {
                continue
            }

            currentFolder.filesSize += num
        }


    }

    var total int
    outputFolder(start, &total)
    fmt.Println(total)
}

func outputFolder(input *Folder, sum *int) {
    total := input.totalSize()

    if (total < 100000) {
        *sum += total
    }

    for _, folder := range input.subfolders {
        outputFolder(folder, sum)
    }
}

type Folder struct {
    name string
    upperFolder *Folder
    subfolders []*Folder
    filesSize int
}

func (f *Folder) String() string {
    return f.name
}

func (f *Folder) createDir(name string) *Folder {
    new := newFolder(name, f)
    f.subfolders = append(f.subfolders, new)

    return new
}

func (f *Folder) totalSize() int {
    totalSize := f.filesSize

    if f.subfolders != nil {
        for _, folder := range f.subfolders {
            totalSize += folder.totalSize()
        }
    }

    return totalSize
}



func newFolder(folderName string, upper *Folder) (*Folder) {
    return &Folder{name: folderName, upperFolder: upper}
}


