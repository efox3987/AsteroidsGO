package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    . "asteroids/game"
    . "asteroids/state"
    . "asteroids/ui/components"
    "strconv"
)


type PlayScreen struct {
    pauseButton *Button
    ships []*Ship
}

func NewPlayScreen(lives int) *PlayScreen {
    ships := make([]*Ship, 0)
    for i := 0; i < lives; i++ {
        s := NewShip()
        s.Position = rl.NewVector2(float32(50 + i*50), 70)
        s.StaticShip = true
        s.NeverFire = true
        ships = append(ships, s)
    }

    return &PlayScreen{
        pauseButton: NewButton(float32(rl.GetScreenWidth()-100), 60, 100, 50, "PAUSE", 20),
        ships: ships,
    }
}

func (p *PlayScreen) Update(g *Game) {
    p.Frame(g)
    p.Draw(g)
}

func (p *PlayScreen) Frame(g *Game) {
    if p.pauseButton.Update() {
        g.State = Pause
    }
}

func (p *PlayScreen) Draw(g *Game) {
    p.pauseButton.Draw()

    // Draw the ship lives
    for i := 0; i < g.Lives; i++ {
        p.ships[i].Update()
    }

    // Draw the score
    rl.DrawText("SCORE: " + strconv.Itoa(g.Score), 10, 10, 20, rl.White)
}


