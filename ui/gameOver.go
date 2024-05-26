package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    . "asteroids/ui/components"
    . "asteroids/state"
    "fmt"
)

type GameOverScreen struct {
    playButton *Button
    menuButton *Button
    gtextWidth int32
}

func NewGameOverScreen() *GameOverScreen {
    return &GameOverScreen {
        playButton: NewButton(float32(rl.GetScreenWidth()/3), float32(rl.GetScreenHeight()/3*2), 200, 100, "PLAY", 50),
        menuButton: NewButton(float32(rl.GetScreenWidth()/3*2), float32(rl.GetScreenHeight()/3*2), 200, 100, "MENU", 50),
        gtextWidth: rl.MeasureText("GAME OVER", 100),
    }
}

func (g *GameOverScreen) Update(st *State, score int) {
    g.Frame(st)
    g.Draw(score)
}

func (g *GameOverScreen) Frame(st *State) {
    if g.playButton.Update() {
        *st = Play
    }
    if g.menuButton.Update() {
        *st = Menu
    }
}

func (g *GameOverScreen) Draw(score int) {
    stext := fmt.Sprintf("SCORE: %d", score)
    stextWidth := rl.MeasureText(stext, 50)
    rl.DrawText("GAME OVER", int32(rl.GetScreenWidth())/2-g.gtextWidth/2, int32(rl.GetScreenHeight()/2-100), 100, rl.White)
    rl.DrawText(stext, int32(rl.GetScreenWidth())/2-stextWidth/2, int32(rl.GetScreenHeight()/2), 50, rl.White)
    g.playButton.Draw()
    g.menuButton.Draw()
}

