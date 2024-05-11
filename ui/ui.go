package ui

import (
    //rl "github.com/gen2brain/raylib-go/raylib"
    //"asteroids/state"
)

type UI struct {
    background Background
}

func NewUI() *UI {
    return &UI{
        background: NewBackground(),
    }
}

func (ui *UI) Update() {
    ui.Draw()
}

func (ui *UI) Draw() {
    ui.background.Draw()
}
