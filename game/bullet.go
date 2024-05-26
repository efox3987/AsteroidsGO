package game

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const BULLET_SPEED = 800

type Bullet struct {
    Position rl.Vector2
    Velocity rl.Vector2
}

func NewBullet(pos rl.Vector2, dir rl.Vector2) *Bullet {
    return &Bullet{
        Position: pos,
        Velocity: rl.Vector2Scale(rl.Vector2Normalize(dir), BULLET_SPEED),
    }
}

func (b *Bullet) Update() {
    b.Frame()
    b.Draw()
}

func (b *Bullet) Pause() {
    b.Draw()
}

func (b *Bullet) Frame() {
    b.Position = rl.Vector2Add(b.Position, rl.Vector2Scale(b.Velocity, rl.GetFrameTime()))
}

func (b *Bullet) Draw() {
    rl.DrawCircleV(b.Position, 3, rl.White)
}
