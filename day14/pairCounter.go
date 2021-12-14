package day14

import (
    "strings"
)

func NewPairCounter(str string) PairCounter {
    c := make(PairCounter)
    lStr := len(str)
    var sb strings.Builder

    for i := 1; i < lStr; i++ {
        sb.WriteRune(rune(str[i-1]))
        sb.WriteRune(rune(str[i]))
        c[sb.String()]++
        sb.Reset()
    }
    return c
}

func (p PairCounter) ApplyPolymer(mappings map[string]string) PairCounter {
    newPairs := make(PairCounter)
    for pair, count := range p {
        mapping, ok := mappings[pair]
        if !ok {
            newPairs[pair] = count
            continue
        }
        newPairs[pair[:1]+mapping] += count
        newPairs[mapping+pair[1:]] += count
    }
    return newPairs
}

func (p PairCounter) Score(first, last rune) int {
    counter := make(map[rune]int)
    var min rune
    var max rune
    for key, count := range p {
        for _, r := range key {
            counter[r] += count
            if max == 0 || counter[r] > counter[max] {
                max = r
            }
            if min == 0 || counter[r] < counter[min] {
                min = r
            }
        }
    }
    counter[first]++
    counter[last]++
    return counter[max]/2 - counter[min]/2
}
