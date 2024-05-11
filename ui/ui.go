package ui

import (
    //rl "github.com/gen2brain/raylib-go/raylib"
    "asteroids/state"
)

type UI struct {
    background Background
    loading Loading
}

func NewUI() *UI {
    return &UI{
        background: NewBackground(),
    }
}

func (ui *UI) Update(state *state.State, time float32) {
    ui.Draw(state, time)
}

func (ui *UI) Draw(state *state.State, time float32) {
    ui.background.Draw()
    ui.loading.Update(state, time)
}
