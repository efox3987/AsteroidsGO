package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "asteroids/game"
    . "asteroids/state"
    . "asteroids/ui/components"
)

/*
    Loading Screen for the game
    Will be displayed when the game is in its start state
    Depends on the time the game has been in the start state
    Will display the title of the game and a loading message
    Modifies the state to Play when the loading time has been reached
*/

const LOAD_TIME = 3.2 //Time in seconds to display the loading screen
const LOADING = "LOADING"
const LOADSZ = 65 //Font size for loading message
const TEXT_FADE_TIME = LOAD_TIME - 1

type Loading struct {
    ship *game.Ship
    loadingPeriods int //Number of periods in the loading message
    textColor rl.Color //Color of the text
    title *Title
}

//Creates a new loading screen
func NewLoading() *Loading {
    l := &Loading{
        ship: game.NewShip(),
        title: NewTitle(TEXT_FADE_TIME),
    }
    l.ship.Position = rl.NewVector2(0, float32(rl.GetScreenHeight() / 2))
    l.ship.Rotation = rl.Pi / 2
    l.ship.Velocity = rl.NewVector2(float32(rl.GetScreenWidth())/LOAD_TIME, 0)
    l.ship.MaxVelocity = 1e9
    l.ship.StaticShip = true
    l.ship.StaticFire = true
    return l
}
  
//Calls the Frame and Draw functions for the loading screen
func (l *Loading) Update(st *State, time float32) {
    l.Frame(st, time)
    l.Draw()
}

//Updates the laoding screen variables
func (l *Loading) Frame(st *State, time float32) {
    //When we have hit the loading time, change the state to Menu
    if time > LOAD_TIME {
        *st = Menu
        return
    }

    l.title.Update()

    l.loadingPeriods = int(time) % RoundUp(LOAD_TIME)

    l.ship.Update()

    //Update the color of the loading message
    l.textColor = rl.NewColor(255, 255, 255, uint8(rl.Clamp(0, 255, 255 * time / TEXT_FADE_TIME)))
}

//Draws the loading screen
func (l *Loading) Draw() {
    loadWidth := rl.MeasureText(LOADING, 20)
    rl.DrawText(LOADING + PeriodString(l.loadingPeriods), int32(rl.GetScreenWidth() / 2- int(loadWidth)/2), 
        int32(rl.GetScreenHeight() / 3 + LOADSZ * 2 - LOADSZ/2), 20, l.textColor)
}

//Returns a string of periods of length periods
func PeriodString(periods int) string {
    s := ""
    for i := 0; i < periods; i++ {
        s += "."
    }
    return s
}

//Helper function to round up a float constant to an int
func RoundUp(f float32) int {
    return int(f + 1)
}
