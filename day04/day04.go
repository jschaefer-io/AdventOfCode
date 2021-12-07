package day04

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

func Solve(data string, result *orchestration.Result) error {
    groups := strings.Split(data, "\n\n")

    // Build Number Set
    set := make([]int, 0)
    for _, sN := range strings.Split(groups[0], ",") {
        n, err := strconv.Atoi(sN)
        if err != nil {
            return err
        }
        set = append(set, n)
    }

    // Build Bingo-Boards
    boards := make([]BingoBoard, 0)
    for i := 1; i < len(groups); i++ {
        b := NewBingoBoard()
        err := b.Fill(groups[i])
        if err != nil {
            return err
        }
        boards = append(boards, b)
    }

    g := NewGame(set, boards)
    for !g.Done() {
        g.Tick()
    }
    winners := g.WinnerList()

    // A
    firstWinner := winners[0]
    result.AddResult(strconv.Itoa(firstWinner.Call * firstWinner.Board.Score()))

    // B
    lastWinner := winners[len(winners)-1]
    result.AddResult(strconv.Itoa(lastWinner.Call * lastWinner.Board.Score()))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day04", orchestration.NewSolver(Solve))
}
