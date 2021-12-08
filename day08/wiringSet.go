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

    // Find 3 by looking for a 5 length wiring using both 1 segments
    oneLetters := letters[1].Segments()
    lengthFilter(3, 5, letters, &leftover, func(w Wiring) bool {
        _, fok := w.Letters[oneLetters[0]]
        _, sok := w.Letters[oneLetters[1]]
        return fok && sok
    })

    // Find 6 by looking for a 6 length wiring having exactly one off the 1 segments
    lengthFilter(6, 6, letters, &leftover, func(w Wiring) bool {
        _, fok := w.Letters[oneLetters[0]]
        _, sok := w.Letters[oneLetters[1]]
        return !fok && sok || fok && !sok
    })

    // Find 9 by looking for a 6 length wiring which contain all segments from 3 and 4
    threeFourLetters := append(letters[3].Segments(), letters[4].Segments()...)
    lengthFilter(9, 6, letters, &leftover, func(w Wiring) bool {
        match := true
        for _, l := range threeFourLetters {
            if _, ok := w.Letters[l]; !ok {
                match = false
                break
            }
        }
        return match
    })

    // 0 is the only wiring left with the length of 6
    lengthFilter(0, 6, letters, &leftover, func(w Wiring) bool {
        return true
    })

    // Find 2 by looking for a wiring left, which contain the segment by removing all 6 segments form 8
    diffLetter := diffLettersRight(letters[8], letters[6])[0]
    lengthFilter(2, -1, letters, &leftover, func(w Wiring) bool {
        _, ok := w.Letters[diffLetter]
        return ok
    })

    // The only letter left at this point will be 5
    letters[5] = leftover[0]

    // Prepare wiring resolver
    reducer := make(map[string]int)
    for k, v := range letters {
        reducer[SortChars(v.String())] = k
    }
    return reducer
}

func lengthFilter(target int, length int, letters map[int]Wiring, leftover *[]Wiring, find func(w Wiring) bool) {
    for s, w := range *leftover {
        if length > 0 && w.Count != length {
            continue
        }
        if find(w) {
            l := *leftover
            newLeftover := append(l[:s], l[s+1:]...)
            *leftover = newLeftover
            letters[target] = w
            return
        }
    }
    panic("unable to apply filter function")
}

func diffLettersRight(a, b Wiring) []rune {
    diff := make([]rune, 0)
    for lA := range a.Letters {
        if _, ok := b.Letters[lA]; !ok {
            diff = append(diff, lA)
        }
    }
    return diff
}
