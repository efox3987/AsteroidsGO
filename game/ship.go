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
    StaticShip bool // If true, the ship will not move
    StaticFire bool // If true, the ship will always show fire
}

const SHIP_SCALE = 45 // Graphical scale to transform ship points to world points
const SHIP_ACCELERATION = 600 // Acceleration of the ship
const SHIP_FIRE_INTERVAL = 0.02 // Time between fire checks
const SHIP_DRAG = 200

// NewShip creates a new Ship object with default values
func NewShip() *Ship {
    return &Ship{
        Position: rl.NewVector2(float32(rl.GetScreenWidth() / 2), float32(rl.GetScreenHeight() / 2)),
        Dir: rl.NewVector2(0, -1),
        Velocity: rl.NewVector2(0, 0),
        MaxVelocity: 300,
        Rotation: 0,
        ShowFire: false,
        ShowFireTime: 0,
        Points: [8]rl.Vector2{ rl.NewVector2(0.3, 0.4), rl.NewVector2(0.4, 0.5), 
        rl.NewVector2(0.0, -0.5), rl.NewVector2(-0.4, 0.5), rl.NewVector2(-0.3, 0.4), 
        rl.NewVector2(0.3, 0.4), rl.NewVector2(0.0, 0.8), rl.NewVector2(-0.3, 0.4) },
        StaticShip: false,
        StaticFire: false,
    }
}

// Calls the Frame and Draw functions for the ship
func (s *Ship) Update() {
    s.Frame()
    s.Draw()
}

// Updates the ship variables
func (s *Ship) Frame() {
    // Determine if fire should show
    s.ShouldFireShow()

    // Process the inputs
    s.ProcessInputs()

    //Update the ship's position
    s.Position = rl.Vector2Add(s.Position, rl.Vector2Scale(s.Velocity, rl.GetFrameTime()))

    //Wrap the ship around the screen
    if s.Position.X > float32(rl.GetScreenWidth()) {
        s.Position.X = 0
    } else if s.Position.X < 0 {
        s.Position.X = float32(rl.GetScreenWidth())
    }

    if s.Position.Y > float32(rl.GetScreenHeight()) {
        s.Position.Y = 0
    } else if s.Position.Y < 0 {
        s.Position.Y = float32(rl.GetScreenHeight())
    }
}

// Determines if the ship should show fire
func (s *Ship) ShouldFireShow() {
    if rl.IsKeyDown(rl.KeyW) || s.StaticFire {
        s.ShowFireTime += rl.GetFrameTime()
        // If the fire time is much larger than the fire interval, reset the fire time
        if s.ShowFireTime > SHIP_FIRE_INTERVAL * 8 {
            s.ShowFire = false
            s.ShowFireTime = 0
        } else if s.ShowFireTime > SHIP_FIRE_INTERVAL {
            s.ShowFire = true
        } else {
            s.ShowFire = false
        } 
    }
}

// Processes the inputs for the ship, calculates the ships velocity and rotation
func (s *Ship) ProcessInputs() {

    // If the ship is static, don't process inputs
    if s.StaticShip {
        return
    }

    // Ship acceleration
    if rl.IsKeyDown(rl.KeyW) {
        accel := rl.Vector2Scale(rl.Vector2Normalize(s.Dir), SHIP_ACCELERATION * rl.GetFrameTime())
        s.Velocity = rl.Vector2Add(s.Velocity, accel)

        // Clamp the velocity
        if rl.Vector2Length(s.Velocity) > s.MaxVelocity {
            s.Velocity = rl.Vector2Scale(rl.Vector2Normalize(s.Velocity), s.MaxVelocity)
        }
    } else {
        // Slow the ship down if no acceleration
        drag := rl.Vector2Scale(rl.Vector2Normalize(s.Velocity), SHIP_DRAG * rl.GetFrameTime()) 
        s.Velocity = rl.Vector2Subtract(s.Velocity, drag)

        s.ShowFire = false
    }

    // Ship rotation
    if rl.IsKeyDown(rl.KeyA) {
        rot := -5 * rl.GetFrameTime()
        s.Rotation += rot
        s.Dir = rl.Vector2Rotate(s.Dir, rot)
    }

    if rl.IsKeyDown(rl.KeyD) {
        rot := 5 * rl.GetFrameTime()
        s.Rotation += rot 
        s.Dir = rl.Vector2Rotate(s.Dir, rot)
    }
}


// Draws the ship to the screen
func (s *Ship) Draw() {
    pointsToDraw := len(s.Points)
    if !s.ShowFire { pointsToDraw -= 3 }

    for i := 0; i < pointsToDraw; i++ {
        rl.DrawLineEx(
            s.TransformPoint(i),
            s.TransformPoint((i + 1) % len(s.Points)),
            2,
            rl.White,
        )
    }
}

// Transforms a point from the ship's local space to world space
func (s *Ship) TransformPoint(i int) rl.Vector2 {
    result := rl.Vector2Rotate(s.Points[i], s.Rotation) 
    result = rl.Vector2Scale(result, SHIP_SCALE)
    result = rl.Vector2Add(result, s.Position)
    return result
}



