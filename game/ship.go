package game 

import rl "github.com/gen2brain/raylib-go/raylib"

type Ship struct {
    Position rl.Vector2
    Dir rl.Vector2
    Velocity rl.Vector2
    MaxVelocity float32
    Rotation float32
    ShowFire bool
    ShowFireTime float32
    Points [8]rl.Vector2 
}

// NewShip creates a new Ship object with default values
func NewShip() *Ship {
    return &Ship{
        Position: rl.NewVector2(float32(rl.GetScreenWidth() / 2), float32(rl.GetScreenHeight() / 2)),
        Dir: rl.NewVector2(0, 1),
        Velocity: rl.NewVector2(0, 0),
        MaxVelocity: 5,
        Rotation: 0,
        ShowFire: false,
        ShowFireTime: 0,
        Points: [8]rl.Vector2{ rl.NewVector2(0.3, 0.4), rl.NewVector2(0.4, 0.5), 
        rl.NewVector2(0.0, -0.5), rl.NewVector2(-0.4, 0.5), rl.NewVector2(-0.3, 0.4), 
        rl.NewVector2(0.3, 0.4), rl.NewVector2(0.0, 0.8), rl.NewVector2(-0.3, 0.4) },
    }
}

// Update updates the ship state
func (s *Ship) Update() {
}






