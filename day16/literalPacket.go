package day16

import (
    "strconv"
    "strings"
)

type LiteralPacket struct {
    Packet
    Value int
}

func (lP LiteralPacket) Eval() int {
    return lP.Value
}

func NewLiteralPacket(p Packet, str string) (LiteralPacket, int) {
    index := 6
    groups := 0
    var rawValue strings.Builder
    for {
        groups++
        rawValue.WriteString(str[index+1 : index+5])
        if rune(str[index]) == '0' {
            break
        }
        index += 5
    }
    value, _ := strconv.ParseInt(rawValue.String(), 2, 64)
    packetSize := groups*5 + 6
    return LiteralPacket{
        Packet: p,
        Value:  int(value),
    }, packetSize
}
