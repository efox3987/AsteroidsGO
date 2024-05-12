package components

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const TITLE = "ASTEROIDS"
const TITLE_SIZE = 100

type Title struct {
    FadeInTime float32
    timeElapsed float32
    textColor rl.Color
}

func NewTitle(fadeTime float32) *Title {
    return &Title{
        FadeInTime: fadeTime,
    }
}

func (t *Title) Update() {
    t.Frame()
    t.Draw()
}

func (t *Title) Frame() {
    t.timeElapsed += rl.GetFrameTime()
    if t.timeElapsed > t.FadeInTime {
        t.textColor = rl.White
    } else {
        t.textColor = rl.NewColor(255, 255, 255, uint8(rl.Clamp(0, 255, 255 * t.timeElapsed / t.FadeInTime)))
    }
}

func (t *Title) Draw() {
    titleWidth := rl.MeasureText(TITLE, TITLE_SIZE)
    rl.DrawText(TITLE, int32(rl.GetScreenWidth() / 2 - int(titleWidth)/2), 
        int32(rl.GetScreenHeight() / 3 - TITLE_SIZE/2), 100, t.textColor)
}

