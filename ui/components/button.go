package components

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
    Rectangle rl.Rectangle
    Text string
    TextSize int32
    clicked bool
}

func NewButton(x, y, width, height float32, text string, textSize int32, ) *Button {
    return &Button{
        Rectangle: rl.NewRectangle(x - width/2, y - height/2, width, height),
        Text: text,
        TextSize: textSize,
    }
}

func (b *Button) Update() bool {
    b.Draw()
    return b.Frame()
}

func (b *Button) Frame() bool {
    if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.Rectangle) {
        if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
            b.clicked = true
        }
    }

    if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.Rectangle) {
        if rl.IsMouseButtonReleased(rl.MouseLeftButton) && b.clicked {
            b.clicked = false
            return true
        }
    }
    return false
}

func (b *Button) Draw() {
    rl.DrawRectangleLinesEx(b.Rectangle, 3.0, rl.White)
    rl.DrawText(b.Text, 
        int32(b.Rectangle.X + b.Rectangle.Width/2 - float32(rl.MeasureText(b.Text, b.TextSize)/2)), 
        int32(b.Rectangle.Y + b.Rectangle.Height/2 - float32(b.TextSize/2)), 
        b.TextSize, 
        rl.White,
    )
}
