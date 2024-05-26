
package state

type State int

const (
    Start       State = iota
    Menu   
    Play 
    Pause 
    GameOver 
)
