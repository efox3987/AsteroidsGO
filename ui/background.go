package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
    stars []rl.Vector2
}

const starsInRow = 20
const starsInColumn = 15

func NewBackground() Background {
    stars := make([]rl.Vector2, (rl.GetScreenWidth() / starsInRow) * (rl.GetScreenHeight() / starsInColumn))
    for i := range stars {
        stars[i] = GetStarVector(i % starsInRow, i / starsInRow)
    }
    return Background{stars: stars}
}

func GetStarVector(row int, column int) rl.Vector2 {
    //Get random value from the start of the row column to the end of the row column
    x := rl.GetRandomValue(int32(rl.GetScreenWidth() / starsInRow * row), int32(rl.GetScreenWidth() / starsInRow * (row + 1)))
    y := rl.GetRandomValue(int32(rl.GetScreenHeight() / starsInColumn * column), int32(rl.GetScreenHeight() / starsInColumn * (column + 1)))

    return rl.NewVector2(float32(x), float32(y))
}

func (b *Background) Draw() {
    for _, star := range b.stars {
        rl.DrawCircle(int32(star.X), int32(star.Y), 1, rl.White)
    }
}
