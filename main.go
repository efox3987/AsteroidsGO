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
    
    g := game.NewGame()
    g.State = state.Start
    
    ui := ui.NewUI()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
        ui.Update()
        
		rl.EndDrawing()
	}
}
