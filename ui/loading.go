package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "asteroids/game"
    . "asteroids/state"
)

const LOAD_TIME = 3.0
const TITLE = "ASTEROIDS"
const LOADING = "Loading..."
const TITLESZ = 100
const LOADSZ = 65

type Loading struct {
    ship *game.Ship
    titleAlpha float32
}

func NewLoading() *Loading {
    l := &Loading{
        ship: game.NewShip(),
        titleAlpha: 0,
    }
    l.ship.Position = rl.NewVector2(0, float32(rl.GetScreenHeight() / 2))
    l.ship.Rotation = 90
    l.ship.Velocity = rl.NewVector2(200, 0)
    return l
}
   
func (l *Loading) Update(st *State, time float32) {
    l.Frame(st, time)
    l.Draw()
}

func (l *Loading) Frame(st *State, time float32) {
    if time > 3.0 {
        *st = Play 
        return
    }
    l.ship.Update()
    l.titleAlpha = 255 * (time / 3.0)
    l.titleAlpha = rl.Clamp(l.titleAlpha, 0, 255)
    rl.NewColor(255, 255, 255, byte(l.titleAlpha))
}

func (l *Loading) Draw() {
    titleWidth := rl.MeasureText(TITLE, 100)
    loadWidth := rl.MeasureText(LOADING, 20)
    rl.DrawText(TITLE, int32(rl.GetScreenWidth() / 2 - int(titleWidth)/2), 
        int32(rl.GetScreenHeight() / 3 - TITLESZ/2), 100, rl.NewColor(255, 255, 255, byte(l.titleAlpha)))
    rl.DrawText(LOADING, int32(rl.GetScreenWidth() / 2- int(loadWidth)/2), 
        int32(rl.GetScreenHeight() / 3 + int(float32(LOADSZ) * 1.5) - LOADSZ/2), 20, rl.NewColor(255, 255, 255, byte(l.titleAlpha)))
}

