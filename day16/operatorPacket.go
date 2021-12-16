package day16

import "strconv"

type OperatorPacket struct {
    Packet
    Packets []Evaluator
}

func (oP OperatorPacket) Eval() int {
    switch oP.Type {
    case 0:
        sum := 0
        for _, p := range oP.Packets {
            sum += p.Eval()
        }
        return sum
    case 1:
        prod := 1
        for _, p := range oP.Packets {
            prod *= p.Eval()
        }
        return prod
    case 2:
        var min int
        for i, p := range oP.Packets {
            v := p.Eval()
            if i == 0 || v < min {
                min = v
            }
        }
        return min
    case 3:
        var max int
        for i, p := range oP.Packets {
            v := p.Eval()
            if i == 0 || v > max {
                max = v
            }
        }
        return max
    case 5:
        if oP.Packets[0].Eval() > oP.Packets[1].Eval() {
            return 1
        }
        return 0
    case 6:
        if oP.Packets[0].Eval() < oP.Packets[1].Eval() {
            return 1
        }
        return 0
    case 7:
        if oP.Packets[0].Eval() == oP.Packets[1].Eval() {
            return 1
        }
        return 0
    }
    panic("unhandled eval type")
}

func (oP OperatorPacket) Sum() int {
    sum := oP.Packet.Sum()
    for _, e := range oP.Packets {
        sum += e.Sum()
    }
    return sum
}

func NewOperatorPacket(p Packet, str string) (OperatorPacket, int) {
    op := OperatorPacket{
        Packet: p,
    }
    var length int
    if str[6] == '0' {
        v, _ := strconv.ParseInt(str[7:22], 2, 64)
        target := 22 + int(v)
        op.Packets, length = FindPackets(str[22:target], -1)
        length += 15
    } else {
        v, _ := strconv.ParseInt(str[7:18], 2, 64)
        op.Packets, length = FindPackets(str[18:], int(v))
        length += 11
    }
    return op, length + 7
}
