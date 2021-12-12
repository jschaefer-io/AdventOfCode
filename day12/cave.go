package day12

import "strings"

type Cave struct {
    isBig       bool
    name        string
    connections []string
}

func NewCave(name string) Cave {
    return Cave{
        name:        name,
        connections: make([]string, 0),
        isBig:       strings.ToLower(name) != name,
    }
}

func (c *Cave) AddConnection(str string) {
    c.connections = append(c.connections, str)
}
