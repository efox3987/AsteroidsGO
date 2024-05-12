package game

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

type Line struct {
    Start rl.Vector2
    End rl.Vector2
}

func NewLine(start rl.Vector2, end rl.Vector2) *Line {
    return &Line{
        Start: start,
        End: end,
    }
}

