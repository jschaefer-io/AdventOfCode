package day04

import (
    "errors"
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

type BingoBoard struct {
    board   [5][5]int
    checks  [5][5]bool
    numbers map[int][2]int
}

func NewBingoBoard() BingoBoard {
    return BingoBoard{
        numbers: make(map[int][2]int),
    }
}

func (b *BingoBoard) Fill(group string) error {
    exp := regexp.MustCompile("\\s+")
    lines := strings.Split(group, "\n")
    if len(lines) < 5 {
        return errors.New("less than 5 lines found")
    }
    for y := 0; y < 5; y++ {
        split := exp.Split(strings.Trim(lines[y], " "), -1)
        if len(split) < 5 {
            return errors.New("line does not contain at least 5 numbers")
        }
        for x := 0; x < 5; x++ {
            n, err := strconv.Atoi(split[x])
            if err != nil {
                return err
            }
            b.board[y][x] = n
            b.numbers[n] = [2]int{y, x}
        }
    }
    return nil
}

func (b *BingoBoard) CallNumber(num int) {
    p, ok := b.numbers[num]
    if !ok {
        return
    }
    b.checks[p[0]][p[1]] = true
}

func (b *BingoBoard) Won() bool {
    for u := 0; u < 5; u++ {
        rowWon := true
        colWon := true
        for i := 0; i < 5; i++ {
            if !b.checks[u][i] {
                rowWon = false
            }
            if !b.checks[i][u] {
                colWon = false
            }
            if !rowWon && !colWon {
                break
            }
        }
        if rowWon || colWon {
            return true
        }
    }
    return false
}

func (b *BingoBoard) Score() int {
    score := 0
    for y := 0; y < 5; y++ {
        for x := 0; x < 5; x++ {
            if !b.checks[y][x] {
                score += b.board[y][x]
            }
        }
    }
    return score
}

func (b BingoBoard) String() string {
    var sb strings.Builder
    sb.WriteString("\n")
    for y := 0; y < 5; y++ {
        for x := 0; x < 5; x++ {
            delimit := " "
            if b.checks[y][x] {
                delimit = "-"
            }
            _, _ = fmt.Fprintf(&sb, "%[1]s%02d%[1]s", delimit, b.board[y][x])
        }
        sb.WriteString("\n")
    }
    return sb.String()
}
