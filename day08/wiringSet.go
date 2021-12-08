package day08

type WiringSet struct {
    Signals [10]Wiring
    Output  [4]Wiring
}

func (set WiringSet) ResolveWiring() map[string]int {
    letters := make(map[int]Wiring)
    leftover := make([]Wiring, 0)
    for _, s := range set.Signals {
        switch s.Count {
        case 2:
            letters[1] = s
        case 3:
            letters[7] = s
        case 4:
            letters[4] = s
        case 7:
            letters[8] = s
        default:
            leftover = append(leftover, s)
        }
    }

    // Find 3
    i3 := find3(letters[1].Segments(), leftover)
    letters[3] = leftover[i3]
    leftover = append(leftover[:i3], leftover[i3+1:]...)

    // Find 6
    i6 := find6(letters[1].Segments(), leftover)
    letters[6] = leftover[i6]
    leftover = append(leftover[:i6], leftover[i6+1:]...)

    // Find 9
    i9 := find9(letters[3], letters[4], leftover)
    letters[9] = leftover[i9]
    leftover = append(leftover[:i9], leftover[i9+1:]...)

    // Find 0
    i0 := find0(leftover)
    letters[0] = leftover[i0]
    leftover = append(leftover[:i0], leftover[i0+1:]...)

    i2 := find2(letters[6], letters[8], leftover)
    letters[2] = leftover[i2]
    leftover = append(leftover[:i2], leftover[i2+1:]...)

    letters[5] = leftover[0]

    reducer := make(map[string]int)
    for k, v := range letters {
        reducer[SortChars(v.String())] = k
    }
    return reducer
}

func find3(contains []rune, list []Wiring) int {
    for s, w := range list {
        if w.Count != 5 {
            continue
        }
        _, fok := w.Letters[contains[0]]
        _, sok := w.Letters[contains[1]]
        if fok && sok {
            return s
        }
    }
    return -1
}

func find6(contains []rune, list []Wiring) int {
    for s, w := range list {
        if w.Count != 6 {
            continue
        }
        _, fok := w.Letters[contains[0]]
        _, sok := w.Letters[contains[1]]
        if !fok && sok || fok && !sok {
            return s
        }
    }
    return -1
}

func find9(letters3, letters4 Wiring, list []Wiring) int {
    letters := append(letters3.Segments(), letters4.Segments()...)
    for s, w := range list {
        if w.Count != 6 {
            continue
        }
        match := true
        for _, l := range letters {
            if _, ok := w.Letters[l]; !ok {
                match = false
                break
            }
        }
        if match {
            return s
        }
    }
    return -1
}

func find0(list []Wiring) int {
    for s, w := range list {
        if w.Count == 6 {
            return s
        }
    }
    return -1
}

func find2(letters6, letters8 Wiring, list []Wiring) int {
    var search rune
    for k := range letters8.Letters {
        if _, ok := letters6.Letters[k]; !ok {
            search = k
            break
        }
    }
    for s, w := range list {
        if _, ok := w.Letters[search]; ok {
            return s
        }
    }
    return -1
}
