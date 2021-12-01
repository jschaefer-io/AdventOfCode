package main

import (
    "fmt"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "os"
    "strings"
)

func getAocWorkload() orchestration.WorkLoad {
    dir, err := os.ReadDir("./_inputs")
    if err != nil {
        panic(err)
    }
    w := make(orchestration.WorkLoad)
    for _, entry := range dir {
        if !entry.IsDir() {
            name := strings.ReplaceAll(entry.Name(), ".txt", "")
            content, err := os.ReadFile("./_inputs/" + entry.Name())
            if err != nil {
                panic(err)
            }
            w[name] = string(content)
        }
    }
    return w
}

func main() {
    fmt.Println()
    orchestration.MainDispatcher.Start(getAocWorkload())
}
