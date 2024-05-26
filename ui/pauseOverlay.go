package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    . "asteroids/ui/components"
    . "asteroids/state"
)

var padding float32 = 100

type PauseOverlay struct {
    playButton *Button
    menuButton *Button
    title *Title
    rect rl.Rectangle
}

func NewPauseOverlay() *PauseOverlay {
    return &PauseOverlay{
        playButton: NewButton(float32(rl.GetScreenWidth()/3), float32(rl.GetScreenHeight()/3*2), 300, 100, "RESUME", 50),
        menuButton: NewButton(float32(rl.GetScreenWidth()/3*2), float32(rl.GetScreenHeight()/3*2), 300, 100, "MENU", 50),
        title: NewTitle(0),
        rect: rl.NewRectangle(padding, padding, float32(rl.GetScreenWidth()) - padding*2, float32(rl.GetScreenHeight()) - padding*2),
    }
}

func (p *PauseOverlay) Update(st *State) {
    p.Frame(st)
    p.Draw()
}

func (p *PauseOverlay) Frame(st *State) {
    if p.playButton.Update() {
        *st = Play
    }
    if p.menuButton.Update() {
        *st = Menu
    }
}

func (p *PauseOverlay) Draw() {
    rl.DrawRectangleRec(p.rect, rl.Fade(rl.Gray, 0.5))
    rl.DrawRectangleLinesEx(p.rect, 10.0, rl.White)
    p.title.Update()
    p.playButton.Draw()
    p.menuButton.Draw()
}

