package day16

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

func FindPackets(str string, limit int) ([]Evaluator, int) {
    list := make([]Evaluator, 0)
    offset := 0
    count := 0
    for limit < 0 || count < limit {
        e, c := ParsePacket(str[offset:])
        offset += c
        if e == nil {
            break
        } else {
            count++
            list = append(list, e)
        }
    }
    return list, offset
}

func ParsePacket(str string) (Evaluator, int) {
    defer func() { recover() }() // @todo properly handle edge cases
    version, _ := strconv.ParseInt(str[:3], 2, 64)
    pType, _ := strconv.ParseInt(str[3:6], 2, 64)
    packet := Packet{
        Version: int(version),
        Type:    int(pType),
    }
    switch pType {
    case 4:
        return NewLiteralPacket(packet, str)
    default:
        return NewOperatorPacket(packet, str)
    }
}

func Solve(data string, result *orchestration.Result) error {
    data = strings.Trim(data, " \n")
    var bitStr strings.Builder
    for _, r := range data {
        n, err := strconv.ParseInt(string(r), 16, 64)
        if err != nil {
            return err
        }
        str := strconv.FormatInt(n, 2)
        strLen := len(str)
        for i := 0; i < 4-strLen; i++ {
            bitStr.WriteString("0")
        }
        bitStr.WriteString(str)
    }

    packet, _ := ParsePacket(bitStr.String())
    result.AddResult(strconv.Itoa(packet.Sum()))
    result.AddResult(strconv.Itoa(packet.Eval()))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day16", orchestration.NewSolver(Solve))
}
