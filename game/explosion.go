package game

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const EXPLOSION_SPEED = 50
const EXPLOSION_SIZE = 10
const EXPLOSION_TIME = 1

type ExplosionBox struct {
    Position rl.Vector2
    Velocity rl.Vector2
    Size float32
    ExplosionTime float32
}

func NewExplosionBox(pos rl.Vector2, dir rl.Vector2, size float32) *ExplosionBox {
    return &ExplosionBox{
        Position: pos,
        Velocity: rl.Vector2Scale(rl.Vector2Normalize(dir), EXPLOSION_SPEED),
        Size: size,
        ExplosionTime: 0,
    }
}

func (e *ExplosionBox) Update() {
    e.Frame()
    e.Draw()
}

func (e *ExplosionBox) Frame() {
    e.ExplosionTime += rl.GetFrameTime()
    e.Position = rl.Vector2Add(e.Position, rl.Vector2Scale(e.Velocity, rl.GetFrameTime()))
}

func (e *ExplosionBox) Draw() {
    rl.DrawRectangleV(e.Position, rl.NewVector2(e.Size, e.Size), rl.White)
}

func (e *ExplosionBox) IsExpired() bool {
    return e.ExplosionTime > EXPLOSION_TIME
}
