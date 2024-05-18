package main 

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "asteroids/state"
    "asteroids/game"
    "asteroids/ui"
)

func main() {
	rl.InitWindow(1280, 960, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
    rl.SetConfigFlags(rl.FlagMsaa4xHint) // Anti-aliasing
    
    g := game.NewGame()
    g.State = state.Start
    
    ui := ui.NewUI(g)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
        g.Update()
        ui.Update(g) // UI relies on game variables
        
		rl.EndDrawing()
	}
}
