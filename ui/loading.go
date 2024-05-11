package ui

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "asteroids/game"
    . "asteroids/state"
)

/*
    Loading Screen for the game
    Will be displayed when the game is in its start state
    Depends on the time the game has been in the start state
    Will display the title of the game and a loading message
    Modifies the state to Play when the loading time has been reached
*/

const LOAD_TIME = 3.2 //Time in seconds to display the loading screen
const TITLE = "ASTEROIDS"
const LOADING = "LOADING"
const TITLESZ = 100 //Font size for title
const LOADSZ = 65 //Font size for loading message

type Loading struct {
    ship *game.Ship
    loadingPeriods int //Number of periods in the loading message
    textColor rl.Color //Color of the text
}

//Creates a new loading screen
func NewLoading() *Loading {
    l := &Loading{
        ship: game.NewShip(),
    }
    l.ship.Position = rl.NewVector2(0, float32(rl.GetScreenHeight() / 2))
    l.ship.Rotation = 90
    l.ship.Velocity = rl.NewVector2(200, 0)
    return l
}
  
//Calls the Frame and Draw functions for the loading screen
func (l *Loading) Update(st *State, time float32) {
    l.Frame(st, time)
    l.Draw()
}

//Updates the laoding screen variables
func (l *Loading) Frame(st *State, time float32) {
    //When we have hit the loading time, change the state to Play
    if time > LOAD_TIME {
        *st = Play 
        return
    }

    l.loadingPeriods = int(time) % RoundUp(LOAD_TIME)

    l.ship.Update()

    //Calculate the alpha value for the title
    titleAlpha := 255 * (time / (LOAD_TIME - 0.5))
    titleAlpha = rl.Clamp(titleAlpha, 0, 255)
    l.textColor = rl.NewColor(255, 255, 255, byte(titleAlpha))
}

//Draws the loading screen
func (l *Loading) Draw() {
    titleWidth := rl.MeasureText(TITLE, 100)
    loadWidth := rl.MeasureText(LOADING, 20)
    rl.DrawText(TITLE, int32(rl.GetScreenWidth() / 2 - int(titleWidth)/2), 
        int32(rl.GetScreenHeight() / 3 - TITLESZ/2), 100, l.textColor)
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
