package day21

import (
    "fmt"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Pawn struct {
    Score    int
    Position int
}

type Game struct {
    Current Pawn
    Next    Pawn
    End     bool
    Dice    Dice
}

func (g Game) Play(limit int) map[Game]int {
    games := make(map[Game]int)
    for roll, count := range g.Dice.Roll() {
        newPos := (g.Current.Position + roll) % 10
        newScore := g.Current.Score + newPos + 1
        test := Game{
            Pawn{g.Next.Score, g.Next.Position},
            Pawn{newScore, newPos},
            newScore >= limit,
            g.Dice,
        }
        games[test] = count
    }
    return games
}

func (g Game) String() string {
    return fmt.Sprintf("%d-%d|%d-%d", g.Current.Score, g.Current.Position, g.Next.Score, g.Next.Position)
}

type Transformation struct {
    Target string
    Count  int
}

func PlayGame(a, b int, limit int, dice Dice) {
    state := make(map[string]int)
    transforms := make(map[string][]Transformation)
    games := make(map[string]Game)

    // Handle Start-State
    start := Game{Pawn{0, a}, Pawn{0, b}, false, dice}
    startId := start.String()
    games[startId] = start
    state[startId] = 1

    for len(state) > 0 {
        nextState := make(map[string]int)
        for id, count := range state {
            targets, ok := transforms[id]
            if ok {
                for _, target := range targets {
                    nextState[target.Target] += target.Count * count
                }
            } else {
                game := games[id]
                transformList := make([]Transformation, 0)
                for res, resCount := range game.Play(limit) {
                    resId := res.String()
                    if !res.End {
                        transformList = append(transformList, Transformation{resId, resCount})
                        nextState[resId] += resCount
                    }
                    games[resId] = res
                }
                transforms[id] = transformList
            }
        }
        state = nextState
        for k, _ := range games {
            if games[k].End {
                fmt.Println("WINNER!!!")
                game := games[k]
                fmt.Println(game.Dice.Count() * game.Current.Score)
                return
            }
        }
    }
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

    PlayGame(startingPositions[0], startingPositions[1], 1000, &DeterministicDice{})

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day21", orchestration.NewSolver(Solve))
}
