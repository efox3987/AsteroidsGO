package game

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "math"
)

type Size int

const (
    Small Size = iota
    Large Size = iota
)

type StartPosition int

const (
    Top StartPosition = iota
    Right StartPosition = iota
    Bottom StartPosition = iota
    Left StartPosition = iota
)

const ASTEROID_SCALE = 5
const ASTEROID_POINTS = 8

type Asteroid struct {
    Position rl.Vector2
    Velocity rl.Vector2
    Radius float32
    Points []rl.Vector2
    AsteroidType Size
    Rotation float32
    RotationSpeed float32
}

func NewAsteroid(size Size) *Asteroid {
    // Set the starting position of the asteroid
    var pos rl.Vector2
    switch StartPosition(rl.GetRandomValue(0, 3)) {
    case Top:
        pos = rl.NewVector2(float32(rl.GetRandomValue(0, int32(rl.GetScreenWidth()))), 0)
    case Right:
        pos = rl.NewVector2(float32(rl.GetScreenWidth()), float32(rl.GetRandomValue(0, int32(rl.GetScreenHeight()))))
    case Bottom:
        pos = rl.NewVector2(float32(rl.GetRandomValue(0, int32(rl.GetScreenWidth()))), float32(rl.GetScreenHeight()))
    case Left:
        pos = rl.NewVector2(0, float32(rl.GetRandomValue(0, int32(rl.GetScreenHeight()))))
    }

    // Set the velocity of the asteroid
    //Velocity direction will be towards the center of the screen with a random angle of 30 degrees in either direction
    angle := rl.GetRandomValue(-30, 30)
    vel := rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2))
    vel = rl.Vector2Subtract(vel, pos)
    vel = rl.Vector2Normalize(vel)
    vel = rl.Vector2Rotate(vel, float32(angle))
    // Set the velocity to a random speed between 50 and 100
    vel = rl.Vector2Scale(vel, float32(rl.GetRandomValue(50, 200)))

    // Set the radius of the asteroid based on the size
    var rad float32
    if size == Large {
        rad = float32(rl.GetRandomValue(10, 15))
    } else {
        rad = float32(rl.GetRandomValue(3, 6))
    }

    // Set the points of the asteroid
    points := make([]rl.Vector2, ASTEROID_POINTS)
    // Set 8 points around the asteroid of varying lengths
    for i := 0; i < ASTEROID_POINTS; i++ {
        angle := float64(i) * rl.Pi / 4
        // Set the length of the point to be a random value between 0.5 and 1.0 times the radius
        length := float64(rl.GetRandomValue(5, 10)) / 10 * float64(rad)
        points[i] = rl.NewVector2(float32(length * math.Cos(angle)), float32(length * math.Sin(angle)))
    }

    // Set the rotation of the asteroid
    rotation := float32(rl.GetRandomValue(0, 360))

    // Set the rotation speed of the asteroid
    rotationSpeed := float32(rl.GetRandomValue(-30, 30)) / 10

    return &Asteroid {
        Position: pos,
        Velocity: vel,
        Radius: rad,
        Points: points,
        AsteroidType: size,
        Rotation: rotation,
        RotationSpeed: rotationSpeed,
    }
}

func (a *Asteroid) Update() {
    a.Frame()
    a.Draw()
}

func (a *Asteroid) Frame() {
    a.Position = rl.Vector2Add(a.Position, rl.Vector2Scale(a.Velocity, rl.GetFrameTime()))

    // Wrap the asteroid around the screen
    if a.Position.X > float32(rl.GetScreenWidth()) {
        a.Position.X = 0
    } else if a.Position.X < 0 {
        a.Position.X = float32(rl.GetScreenWidth())
    }

    if a.Position.Y > float32(rl.GetScreenHeight()) {
        a.Position.Y = 0
    } else if a.Position.Y < 0 {
        a.Position.Y = float32(rl.GetScreenHeight())
    }

    a.Rotation += a.RotationSpeed * rl.GetFrameTime()
}

func (a *Asteroid) Draw() {
    for i := 0; i < ASTEROID_POINTS; i++ {
        rl.DrawLineEx(a.TransformPoint(i), 
            a.TransformPoint((i + 1) % ASTEROID_POINTS), 
            2, 
            rl.White,
        )
    } 
}

func (a *Asteroid) TransformPoint(i int) rl.Vector2 {
    result := rl.Vector2Rotate(a.Points[i], a.Rotation) 
    result = rl.Vector2Scale(result, ASTEROID_SCALE)
    result = rl.Vector2Add(result, a.Position)
    return result
}
















