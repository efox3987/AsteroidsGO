package game 

// Contains the state for a game and other functions to manipulate the game state

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    . "asteroids/state"
)

const STARTING_ASTEROIDS = 3
const CREATE_ASTEROID_TIME = 1
const STARTING_ASTEROID_CAP = 5
const ASTEROID_CAP_TIME = 7

type Game struct {
    State State
    Ship *Ship
    Bullets []*Bullet
    Asteroids []*Asteroid
    Score int 
    Time float32
    AsteroidTimer float32
    AsteroidCap int
    AsteroidCapTimer float32
}

// NewGame creates a new Game object with default values
func NewGame() *Game {
    asteroids := []*Asteroid{}
    for i := 0; i < STARTING_ASTEROIDS; i++ {
        asteroids = append(asteroids, NewAsteroid(Large))
    }

    return &Game{
        State: Start,
        Ship: NewShip(),
        Bullets: []*Bullet{},
        Asteroids: asteroids,
        Score: 0,
        Time: 0,
        AsteroidTimer: 0,
        AsteroidCap: STARTING_ASTEROID_CAP,
        AsteroidCapTimer: 0,
    }
}

// Update updates the game state
func (g *Game) Update() {
    g.Time += rl.GetFrameTime()

    switch g.State {
    case Start:
        // No game related updates here
        break
    case Menu:
        // No game related updates here
        break
    case Play:
        g.ProcessInputs()

        g.Ship.Update()

        g.ProcessAsteroids()

        for _, b := range g.Bullets {
            b.Update()
        }

        for _, a := range g.Asteroids {
            a.Update()
        }

        //g.checkCollisions()
        //g.checkGameOver()
    case GameOver:
        if rl.IsKeyPressed(rl.KeySpace) {
            g.State = Start
            g.Score = 0
            g.Time = 0
            g.Ship = NewShip()
            //g.Bullets = []*Bullet{}
            //g.Asteroids = []*Asteroid{}
        }
    }
}

func (g *Game) ProcessInputs() {
    if rl.IsKeyPressed(rl.KeySpace) {
        g.Shoot()
    }
}

func (g *Game) Shoot() {
    // Create the new bullet at the front point of the ship 
    // Front point is at the 2 index of the Points array
    b := NewBullet(g.Ship.TransformPoint(2), g.Ship.Dir)
    g.Bullets = append(g.Bullets, b)
}

func (g *Game) ProcessAsteroids() {
    g.AsteroidTimer += rl.GetFrameTime()
    if g.AsteroidTimer > CREATE_ASTEROID_TIME {
        if len(g.Asteroids) < g.AsteroidCap {
            g.AsteroidTimer = 0
            g.Asteroids = append(g.Asteroids, NewAsteroid(Large))
        }
    }

    g.AsteroidCapTimer += rl.GetFrameTime()
    if g.AsteroidCapTimer > ASTEROID_CAP_TIME {
        g.AsteroidCapTimer = 0
        g.AsteroidCap++
    }
}






 
