package game 

// Contains the state for a game and other functions to manipulate the game state

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "asteroids/state"
)

type Game struct {
    State state.State
    Ship *Ship
    //Bullets []*Bullet
    //Asteroids []*Asteroid
    Score int 
    Time float32
}

// NewGame creates a new Game object with default values
func NewGame() *Game {
    return &Game{
        State: state.Start,
        Ship: NewShip(),
        //Bullets: []*Bullet{},
        //Asteroids: []*Asteroid{},
        Score: 0,
        Time: 0,
    }
}

// Update updates the game state
func (g *Game) Update() {
    g.Time += rl.GetFrameTime()

    switch g.State {
    case state.Start:
        if rl.IsKeyPressed(rl.KeySpace) {
            g.State = state.Play
        }
    case state.Play:
        g.Ship.Update()
        //for _, b := range g.Bullets {
        //    b.Update()
        //}
        //for _, a := range g.Asteroids {
        //    a.Update()
        //}
        //g.checkCollisions()
        //g.checkGameOver()
    case state.GameOver:
        if rl.IsKeyPressed(rl.KeySpace) {
            g.State = state.Start
            g.Score = 0
            g.Time = 0
            g.Ship = NewShip()
            //g.Bullets = []*Bullet{}
            //g.Asteroids = []*Asteroid{}
        }
    }
}
