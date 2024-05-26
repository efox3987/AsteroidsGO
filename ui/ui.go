package ui

import (
    //rl "github.com/gen2brain/raylib-go/raylib"
    . "asteroids/game"
    . "asteroids/state"
)

type UI struct {
    background Background
    loading *Loading
    menu *MenuScreen
    play *PlayScreen
    pause *PauseOverlay
    gameOver *GameOverScreen
}

func NewUI(g *Game) *UI {
    return &UI{
        background: NewBackground(),
        loading: NewLoading(),
        menu: NewMenuScreen(),
        play: NewPlayScreen(g.Lives),
        pause: NewPauseOverlay(),
        gameOver: NewGameOverScreen(),
    }
}

func (ui *UI) Update(g *Game) {
    switch g.State {
    case Start:
        ui.loading.Update(&g.State, g.Time)
    case Menu:
        ui.menu.Update(&g.State)
    case Play:
        ui.play.Update(g)
    case Pause:
        ui.pause.Update(&g.State)
    case GameOver:
        ui.gameOver.Update(&g.State, g.Score)
    }

    // Draw the static elements of the UI
    ui.StaticDraw()
}

// Function to draw UI elements that do not change
func (ui *UI) StaticDraw() {
    ui.background.Draw()
}
