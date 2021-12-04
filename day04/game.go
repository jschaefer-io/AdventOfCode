package day04

import "errors"

type Winner struct {
    Call  int
    Board BingoBoard
}

type Game struct {
    call    int
    numbers []int
    winners [][]BingoBoard
    boards  []BingoBoard
    ignore  map[int]struct{}
}

func NewGame(numbers []int, boards []BingoBoard) Game {
    return Game{
        call:    0,
        numbers: numbers,
        winners: make([][]BingoBoard, 0),
        boards:  boards,
        ignore:  make(map[int]struct{}),
    }
}

func (g *Game) checkWinners() {
    winners := make([]BingoBoard, 0)
    for i := range g.boards {
        if _, ok := g.ignore[i]; ok {
            continue
        }
        if g.boards[i].Won() {
            winners = append(winners, g.boards[i])
            g.ignore[i] = struct{}{}
        }
    }
    g.winners = append(g.winners, winners)
}

func (g *Game) WinnerList() []Winner {
    list := make([]Winner, 0)
    for call, boards := range g.winners {
        for _, board := range boards {
            list = append(list, Winner{
                Call:  g.numbers[call],
                Board: board,
            })
        }
    }
    return list
}

func (g *Game) LastCall() (int, error) {
    if g.call == 0 {
        return 0, errors.New("no numbers called yet")
    }
    return g.numbers[g.call-1], nil
}

func (g *Game) Tick() {
    n := g.numbers[g.call]
    for i := range g.boards {
        if _, ok := g.ignore[i]; ok {
            continue
        }
        g.boards[i].CallNumber(n)
    }
    g.checkWinners()
    g.call++
}

func (g *Game) Done() bool {
    return g.call >= len(g.numbers)
}
