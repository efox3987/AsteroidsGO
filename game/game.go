package game 

// Contains the state for a game and other functions to manipulate the game state

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    . "asteroids/state"
    "math"
)

const STARTING_ASTEROIDS = 3
const CREATE_ASTEROID_TIME = 1
const STARTING_ASTEROID_CAP = 5
const ASTEROID_CAP_TIME = 7
const SHIP_EXPLODE_TIME = 3
const SHOOT_COOLDOWN = 0.2

type Game struct {
    State State
    Ship *Ship
    Bullets []*Bullet
    Asteroids []*Asteroid
    Score int 
    Lives int
    Time float32
    AsteroidTimer float32
    AsteroidCap int
    AsteroidCapTimer float32
    ShipExploding bool
    ShipExplodeTime float32
    Explosion []*ExplosionBox
    ShootTimer float32
}

// NewGame creates a new Game object with default values
func NewGame() *Game {
    return &Game{
        State: Start,
        Ship: NewShip(),
        Bullets: []*Bullet{},
        Asteroids: InitAsteroids(),
        Score: 0,
        Lives: 3,
        Time: 0,
        AsteroidTimer: 0,
        AsteroidCap: STARTING_ASTEROID_CAP,
        AsteroidCapTimer: 0,
        ShipExploding: false,
        ShipExplodeTime: 0,
        Explosion: []*ExplosionBox{},
        ShootTimer: 0,
    }
}

// Function to restart the game when the player dies
func (g *Game) RestartGame() {
    if g.Lives <= 0 {
        g.State = GameOver
    }
    g.Ship = NewShip() 
    g.Bullets = []*Bullet{}
    g.Asteroids = InitAsteroids()
    g.Lives--
    g.AsteroidTimer = 0
    g.AsteroidCap = STARTING_ASTEROID_CAP
    g.AsteroidCapTimer = 0
    g.ShipExploding = false
    g.ShipExplodeTime = 0
    g.Explosion = []*ExplosionBox{}
    g.ShootTimer = 0
}

// Initialize the asteroid slice
func InitAsteroids() []*Asteroid {    
asteroids := []*Asteroid{}
    for i := 0; i < STARTING_ASTEROIDS; i++ {
        asteroids = append(asteroids, NewAsteroid(Large))
    }
    return asteroids
}

// Process game objects and draw game specific elements
func (g *Game) Update() {
    g.Time += rl.GetFrameTime()
    if g.State != Play {
        return
    }   
    

    g.ProcessInputs()
    
    g.ProcessShip()

    g.ProcessAsteroids()
    
    g.ProcessBullets()

    g.ProcessExplosions()

    g.checkCollisions()
}

// Process game inputs, ship movement is handled by the ship object
func (g *Game) ProcessInputs() {
    g.ShootTimer += rl.GetFrameTime()
    if rl.IsKeyPressed(rl.KeySpace) {
        if g.ShootTimer > SHOOT_COOLDOWN {
            g.ShootTimer = 0
            g.Shoot()
        }
    }
}

// Process the ship variables and draw the ship
func (g *Game) ProcessShip() {
     if g.ShipExploding {
        g.ShipExplodeTime += rl.GetFrameTime()
        if g.ShipExplodeTime > EXPLOSION_TIME {
            g.ShipExploding = false
            g.ShipExplodeTime = 0
            g.RestartGame()
        }
    } else {
        g.Ship.Update()
    }
}

func (g *Game) ProcessBullets() {
    for _, b := range g.Bullets {
        b.Update()
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
        if g.GetLargeAsteroids() < g.AsteroidCap {
            g.AsteroidTimer = 0
            g.Asteroids = append(g.Asteroids, NewAsteroid(Large))
        }
    }

    g.AsteroidCapTimer += rl.GetFrameTime()
    if g.AsteroidCapTimer > ASTEROID_CAP_TIME {
        g.AsteroidCapTimer = 0
        g.AsteroidCap++
    }

    for _, a := range g.Asteroids {
        a.Update()
    }
}

// Returns the number of large asteroids in the game
func (g *Game) GetLargeAsteroids() int {
    count := 0
    for _, a := range g.Asteroids {
        if a.AsteroidType == Large {
            count++
        }
    }
    return count
}

func (g *Game) checkCollisions() {
    g.checkBulletAsteroidCollisions()
    g.checkShipAsteroidCollisions()
}

func (g *Game) checkBulletAsteroidCollisions() {
    for _, b := range g.Bullets {
        for _, a := range g.Asteroids {
            for _, line := range a.GetLines() {
                if rl.CheckCollisionPointLine(b.Position, line.Start, line.End, 20) {
                    g.RemoveBullet(b)
                    g.RemoveAsteroid(a)
                    g.NewExplosion(a.Position, 7)
                    g.Score += 10
                    break
                }
            }
        }
    }
}

func (g *Game) checkShipAsteroidCollisions() {
    for _, a := range g.Asteroids {
        for _, aLine := range a.GetLines() {
            for _, sLine := range g.Ship.GetLines() {
                collisionPoint := rl.NewVector2(0, 0)
                if rl.CheckCollisionLines(aLine.Start, aLine.End, sLine.Start, sLine.End, &collisionPoint) {
                    if !g.ShipExploding {
                        g.NewExplosion(g.Ship.Position, 15)
                        g.ShipExploding = true
                        g.ShipExplodeTime = 0
                    }
                    return
                }
            }
        }
    }
}

func (g *Game) RemoveBullet(b *Bullet) {
    for i, bullet := range g.Bullets {
        if bullet == b {
            g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
            return
        }
    }
}

func (g *Game) RemoveAsteroid(a *Asteroid) {
    for i, asteroid := range g.Asteroids {
        if asteroid == a {
            g.Asteroids = append(g.Asteroids[:i], g.Asteroids[i+1:]...)
            if a.AsteroidType == Large {
                // Insert two new Small asteroids
                g.Asteroids = append(g.Asteroids, NewAsteroidPos(Small, a.Position))
                g.Asteroids = append(g.Asteroids, NewAsteroidPos(Small, a.Position))
            }
            return
        }
    }
}

func (g *Game) NewExplosion(pos rl.Vector2, size float32) {
    // Create a new explosion at the ship's position
    for i := 0; i < 8; i += 1 {
        angle := float32(i) * rl.Pi / 4
        dir := rl.NewVector2(float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle))))
        g.Explosion = append(g.Explosion, NewExplosionBox(pos, dir, size))
    }
}

func (g *Game) ProcessExplosions() {
    for i, e := range g.Explosion {
        if i >= len(g.Explosion) {
            break
        }
        if e.IsExpired() {
            g.Explosion = append(g.Explosion[:i], g.Explosion[i+1:]...)
        }
    }

    for _, e := range g.Explosion {
        e.Update()
    }
} 
