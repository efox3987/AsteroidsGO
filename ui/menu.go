package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    . "asteroids/ui/components"
    . "asteroids/state"
    "asteroids/game"
)

type MenuScreen struct {
    title *Title
    playButton *Button
    quitButton *Button
    ship *game.Ship
}     

func NewMenuScreen() *MenuScreen {
    return &MenuScreen{
        title: NewTitle(0),
        playButton: NewButton(float32(rl.GetScreenWidth()/3), float32(rl.GetScreenHeight()/3*2), 200, 100, "PLAY", 50),
        quitButton: NewButton(float32(rl.GetScreenWidth()/3*2), float32(rl.GetScreenHeight()/3*2), 200, 100, "QUIT", 50),
        ship: game.NewShip(),
    }
}

func (m *MenuScreen) Update(st *State) {
    m.Frame(st)
}

func (m *MenuScreen) Frame(st *State) {
    m.title.Update()
    if m.playButton.Update() {
        *st = Play
    }
    
    if m.quitButton.Update() {
        rl.CloseWindow()
    }

    m.ship.Update()
}


