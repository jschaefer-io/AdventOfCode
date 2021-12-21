package day21

import "fmt"

type Pawn struct {
    Score    int
    Position int
}

type Game struct {
    Turn      bool
    A         Pawn
    B         Pawn
    End       bool
    Dice      Dice
    Condition int
}

func (g Game) Play() []Game {
    games := make([]Game, 0)
    for _, roll := range g.Dice.Roll() {
        var nextGame Game
        var newPos int
        var newScore int
        if !g.Turn {
            newPos = (g.A.Position + roll) % 10
            newScore = g.A.Score + newPos + 1
            nextGame = Game{
                A: Pawn{newScore, newPos},
                B: Pawn{g.B.Score, g.B.Position},
            }
        } else {
            newPos = (g.B.Position + roll) % 10
            newScore = g.B.Score + newPos + 1
            nextGame = Game{
                A: Pawn{g.A.Score, g.A.Position},
                B: Pawn{newScore, newPos},
            }
        }
        nextGame.End = nextGame.A.Score >= g.Condition || nextGame.B.Score >= g.Condition
        nextGame.Turn = !g.Turn
        nextGame.Dice = g.Dice
        nextGame.Condition = g.Condition
        games = append(games, nextGame)
    }
    return games
}

func (g Game) String() string {
    return fmt.Sprintf("%d,%d|%d,%d|%v", g.A.Position+1, g.A.Score, g.B.Position+1, g.B.Score, g.Turn)
}
