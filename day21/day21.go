package day21

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "strconv"
    "strings"
)

func findWinner(game Game, winners map[string]Game, counts map[string][2]int) (int, int) {
    if game.End {
        winners[game.String()] = game
        if game.A.Score >= game.Condition {
            return 1, 0
        } else {
            return 0, 1
        }
    }
    aWin := 0
    bWin := 0
    for _, g := range game.Play() {
        if res, ok := counts[g.String()]; ok {
            aWin += res[0]
            bWin += res[1]
        } else {
            a, b := findWinner(g, winners, counts)
            aWin += a
            bWin += b
        }
    }
    counts[game.String()] = [2]int{aWin, bWin}
    return aWin, bWin
}

func PlayGame(a, b int, limit int, dice Dice) (map[string]Game, int, int) {
    winners := make(map[string]Game)
    winnerCount := make(map[string][2]int)
    game := Game{
        A:         Pawn{Position: a},
        B:         Pawn{Position: b},
        Dice:      dice,
        Condition: limit,
    }
    aWin, bWin := findWinner(game, winners, winnerCount)
    return winners, aWin, bWin
}

func Solve(data string, result *orchestration.Result) error {
    startingPositions := make([]int, 0)
    for _, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        parts := strings.Split(line, ": ")
        pos, err := strconv.Atoi(parts[1])
        if err != nil {
            return err
        }
        startingPositions = append(startingPositions, pos-1)
    }

    // a
    winners, _, _ := PlayGame(startingPositions[0], startingPositions[1], 1000, &DeterministicDice{})
    for _, winner := range winners {
        a := winner.Dice.Count() * int(math.Min(float64(winner.A.Score), float64(winner.B.Score)))
        result.AddResult(strconv.Itoa(a))
    }

    // b
    _, aPawn, bPawn := PlayGame(startingPositions[0], startingPositions[1], 21, &DiracDice{})
    b := int(math.Max(float64(aPawn), float64(bPawn)))
    result.AddResult(strconv.Itoa(b))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day21", orchestration.NewSolver(Solve))
}
