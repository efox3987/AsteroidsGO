package ui

import (
    //rl "github.com/gen2brain/raylib-go/raylib"
    "asteroids/state"
)

type UI struct {
    background Background
    loading *Loading
}

func NewUI() *UI {
    return &UI{
        background: NewBackground(),
        loading: NewLoading(),
    }
}

func (ui *UI) Update(st *state.State, time float32) {
    if *st == state.Start {
        ui.loading.Update(st, time)
    }
    ui.Draw(st, time)
}

func (ui *UI) Draw(st *state.State, time float32) {
    ui.background.Draw()
    ui.loading.Update(st, time)
}
