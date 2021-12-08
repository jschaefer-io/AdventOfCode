package day08

import (
    "errors"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "sort"
    "strconv"
    "strings"
)

func SortChars(s string) string {
    l := strings.Split(s, "")
    sort.Strings(l)
    var sb strings.Builder
    for _, s := range l {
        sb.WriteString(s)
    }
    return sb.String()
}

func Solve(data string, result *orchestration.Result) error {

    wiringSets := make([]WiringSet, 0)
    for _, line := range strings.Split(data, "\n") {
        set := WiringSet{}
        if len(line) == 0 {
            continue
        }
        parts := strings.Split(line, " | ")
        signalPart := strings.Split(parts[0], " ")
        outputPart := strings.Split(parts[1], " ")
        if len(signalPart) != 10 {
            return errors.New("less than 10 signal patterns detected")
        }
        if len(outputPart) != 4 {
            return errors.New("less than 4 output patterns detected")
        }
        for i, signal := range signalPart {
            set.Signals[i] = NewWiring(signal)
        }
        for i, output := range outputPart {
            set.Output[i] = NewWiring(output)
        }
        wiringSets = append(wiringSets, set)
    }

    // A
    count := 0
    for _, set := range wiringSets {
        for _, wiring := range set.Output {
            if wiring.Count == 2 || wiring.Count == 3 || wiring.Count == 4 || wiring.Count == 7 {
                count++
            }
        }
    }
    result.AddResult(strconv.Itoa(count))

    // B
    s := 0
    for _, ws := range wiringSets {
        m := ws.ResolveWiring()
        out := 0
        for i, w := range ws.Output {
            out += int(math.Pow(10, 3.0-float64(i))) * m[SortChars(w.String())]
        }
        s += out
    }
    result.AddResult(strconv.Itoa(s))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day08", orchestration.NewSolver(Solve))
}
